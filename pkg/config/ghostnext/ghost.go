package ghostnext

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/hcl/v2"

	dataproviderConfig "github.com/chronicleprotocol/oracle-suite/pkg/config/dataprovider"
	ethereumConfig "github.com/chronicleprotocol/oracle-suite/pkg/config/ethereum"
	feedConfig "github.com/chronicleprotocol/oracle-suite/pkg/config/feednext"
	loggerConfig "github.com/chronicleprotocol/oracle-suite/pkg/config/logger"
	transportConfig "github.com/chronicleprotocol/oracle-suite/pkg/config/transport"
	"github.com/chronicleprotocol/oracle-suite/pkg/feed"
	"github.com/chronicleprotocol/oracle-suite/pkg/log"

	pkgSupervisor "github.com/chronicleprotocol/oracle-suite/pkg/supervisor"
	"github.com/chronicleprotocol/oracle-suite/pkg/sysmon"
	pkgTransport "github.com/chronicleprotocol/oracle-suite/pkg/transport"
	"github.com/chronicleprotocol/oracle-suite/pkg/transport/messages"
)

// Config is the configuration for Lair.
type Config struct {
	Ghost     feedConfig.Config         `hcl:"ghostnext,block"`
	Gofer     dataproviderConfig.Config `hcl:"gofernext,block"`
	Ethereum  ethereumConfig.Config     `hcl:"ethereum,block"`
	Transport transportConfig.Config    `hcl:"transport,block"`
	Logger    *loggerConfig.Config      `hcl:"logger,block,optional"`

	// HCL fields:
	Remain  hcl.Body        `hcl:",remain"` // To ignore unknown blocks.
	Content hcl.BodyContent `hcl:",content"`
}

// Services returns the services that are configured from the Config struct.
type Services struct {
	Feed      *feed.Feed
	Transport pkgTransport.Transport
	Logger    log.Logger

	supervisor *pkgSupervisor.Supervisor
}

// Start implements the supervisor.Service interface.
func (s *Services) Start(ctx context.Context) error {
	if s.supervisor != nil {
		return fmt.Errorf("services already started")
	}
	s.supervisor = pkgSupervisor.New(s.Logger)
	s.supervisor.Watch(s.Transport, s.Feed, sysmon.New(time.Minute, s.Logger))
	if l, ok := s.Logger.(pkgSupervisor.Service); ok {
		s.supervisor.Watch(l)
	}
	return s.supervisor.Start(ctx)
}

// Wait implements the supervisor.Service interface.
func (s *Services) Wait() <-chan error {
	return s.supervisor.Wait()
}

// Services returns the services configured for Lair.
func (c *Config) Services(baseLogger log.Logger) (*Services, error) {
	logger, err := c.Logger.Logger(loggerConfig.Dependencies{
		AppName:    "ghost",
		BaseLogger: baseLogger,
	})
	if err != nil {
		return nil, err
	}
	keys, err := c.Ethereum.KeyRegistry(ethereumConfig.Dependencies{Logger: logger})
	if err != nil {
		return nil, err
	}
	clients, err := c.Ethereum.ClientRegistry(ethereumConfig.Dependencies{Logger: logger})
	if err != nil {
		return nil, err
	}
	transport, err := c.Transport.Transport(transportConfig.Dependencies{
		Keys:    keys,
		Clients: clients,
		Messages: map[string]pkgTransport.Message{
			messages.DataPointV1MessageName: (*messages.DataPoint)(nil),
		},
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}
	dataProvider, err := c.Gofer.ConfigureDataProvider(dataproviderConfig.Dependencies{
		Clients: clients,
		Logger:  logger,
	})
	if err != nil {
		return nil, err
	}
	feedService, err := c.Ghost.ConfigureFeed(feedConfig.Dependencies{
		KeysRegistry: keys,
		DataProvider: dataProvider,
		Transport:    transport,
		Logger:       logger,
	})
	if err != nil {
		return nil, err
	}
	return &Services{
		Feed:      feedService,
		Transport: transport,
		Logger:    logger,
	}, nil
}
