package feed

import (
	"context"
	"errors"

	"github.com/chronicleprotocol/oracle-suite/pkg/datapoint"
	"github.com/chronicleprotocol/oracle-suite/pkg/log/null"
	"github.com/chronicleprotocol/oracle-suite/pkg/transport/messages"
	"github.com/chronicleprotocol/oracle-suite/pkg/util/timeutil"

	"github.com/chronicleprotocol/oracle-suite/pkg/log"
	"github.com/chronicleprotocol/oracle-suite/pkg/transport"
)

const LoggerTag = "FEED"

// Feed is a service which periodically fetches data points and then sends them to
// the network using transport layer.
type Feed struct {
	ctx    context.Context
	waitCh chan error
	log    log.Logger

	dataProvider datapoint.Provider
	dataModels   []string
	signers      []datapoint.Signer
	transport    transport.Transport
	interval     *timeutil.Ticker
}

// Config is the configuration for the Feed.
type Config struct {
	// DataModels is a list of data models handled by the Feed.
	DataModels []string

	// DataProvider is a data provider which is used to fetch data points.
	DataProvider datapoint.Provider

	// Signers is a list of signers used to sign data points.
	Signers []datapoint.Signer

	// Transport is an implementation of transport used to send prices to
	// the network.
	Transport transport.Transport

	// Interval describes how often data points should be sent to the network.
	Interval *timeutil.Ticker

	// Logger is a current logger interface used by the Feed.
	// If nil, null logger will be used.
	Logger log.Logger
}

// New creates a new instance of the Feed.
func New(cfg Config) (*Feed, error) {
	if cfg.DataModels == nil {
		return nil, errors.New("data provider must not be nil")
	}
	if cfg.Transport == nil {
		return nil, errors.New("transport must not be nil")
	}
	if cfg.Logger == nil {
		cfg.Logger = null.New()
	}
	g := &Feed{
		waitCh:       make(chan error),
		log:          cfg.Logger.WithField("tag", LoggerTag),
		dataProvider: cfg.DataProvider,
		dataModels:   cfg.DataModels,
		signers:      cfg.Signers,
		transport:    cfg.Transport,
		interval:     cfg.Interval,
	}
	return g, nil
}

// Start implements the supervisor.Service interface.
func (f *Feed) Start(ctx context.Context) error {
	if f.ctx != nil {
		return errors.New("service can be started only once")
	}
	if ctx == nil {
		return errors.New("context must not be nil")
	}
	f.log.Infof("Starting")
	f.ctx = ctx
	f.interval.Start(f.ctx)
	go f.broadcasterRoutine()
	go f.contextCancelHandler()
	return nil
}

// Wait implements the supervisor.Service interface.
func (f *Feed) Wait() <-chan error {
	return f.waitCh
}

// broadcast sends data point to the network.
func (f *Feed) broadcast(model string, point datapoint.Point) {
	found := false
	for _, signer := range f.signers {
		if !signer.Supports(f.ctx, point) {
			continue
		}
		found = true
		sig, err := signer.Sign(f.ctx, model, point)
		if err != nil {
			f.log.
				WithError(err).
				WithField("dataPoint", point).
				Error("Unable to sign data point")
		}
		msg := &messages.DataPoint{
			Model:     model,
			Value:     point,
			Signature: *sig,
		}
		if err := f.transport.Broadcast(messages.DataPointV1MessageName, msg); err != nil {
			f.log.
				WithError(err).
				WithField("dataPoint", point).
				Error("Unable to broadcast data point")
		}
		f.log.
			WithField("dataPoint", point).
			Info("Data point broadcast")
	}
	if !found {
		f.log.
			WithField("dataPoint", point).
			Warn("Unable to find handler for data point")
	}
}

func (f *Feed) broadcasterRoutine() {
	for {
		select {
		case <-f.ctx.Done():
			return
		case <-f.interval.TickCh():
			// Fetch all data points from the provider to update them
			// at once.
			_, err := f.dataProvider.DataPoints(f.ctx, f.dataModels...)
			if err != nil {
				f.log.
					WithError(err).
					Error("Unable to update data points")
				continue
			}

			// Send data points to the network.
			for _, model := range f.dataModels {
				point, err := f.dataProvider.DataPoint(f.ctx, model)
				if err != nil {
					f.log.
						WithError(err).
						Error("Unable to get data points")
					continue
				}
				f.broadcast(model, point)
			}
		}
	}
}

func (f *Feed) contextCancelHandler() {
	defer func() { close(f.waitCh) }()
	defer f.log.Info("Stopped")
	<-f.ctx.Done()
}
