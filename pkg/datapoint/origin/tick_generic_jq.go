package origin

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/itchyny/gojq"

	"github.com/chronicleprotocol/oracle-suite/pkg/datapoint"
	"github.com/chronicleprotocol/oracle-suite/pkg/datapoint/value"

	"github.com/chronicleprotocol/oracle-suite/pkg/log"
	"github.com/chronicleprotocol/oracle-suite/pkg/log/null"
	"github.com/chronicleprotocol/oracle-suite/pkg/util/bn"
)

const TickGenericJQLoggerTag = "TICK_GENERIC_JQ_ORIGIN"

type TickGenericJQOptions struct {
	// URL is an TickGenericHTTP endpoint that returns JSON data. It may contain
	// the following variables:
	//   - ${lcbase} - lower case base asset
	//   - ${ucbase} - upper case base asset
	//   - ${lcquote} - lower case quote asset
	//   - ${ucquote} - upper case quote asset
	//   - ${lcbases} - lower case base assets joined by commas
	//   - ${ucbases} - upper case base assets joined by commas
	//   - ${lcquotes} - lower case quote assets joined by commas
	//   - ${ucquotes} - upper case quote assets joined by commas
	URL string

	// Query is a JQ query that is used to parse JSON data. It must
	// return a single value that will be used as a price or an object with the
	// following fields:
	//   - price - a price
	//   - time - a timestamp (optional)
	//   - volume - a 24h volume (optional)
	//
	// The JQ query may contain the following variables:
	//   - $lcbase - lower case base asset
	//   - $ucbase - upper case base asset
	//   - $lcquote - lower case quote asset
	//   - $ucquote - upper case quote asset
	//
	// Price and volume must be a number or a string that can be parsed as a number.
	// Time must be a number or a string that can be parsed as a number or a string
	// that can be parsed as a time.
	Query string

	// Headers is a set of TickGenericHTTP headers that are sent with each request.
	Headers http.Header

	// Client is an TickGenericHTTP client that is used to fetch data from the
	// TickGenericHTTP endpoint. If nil, http.DefaultClient is used.
	Client *http.Client

	// Logger is an TickGenericHTTP logger that is used to log errors. If nil,
	// null logger is used.
	Logger log.Logger
}

// TickGenericJQ is a generic origin implementation that uses JQ to parse JSON data
// from an TickGenericHTTP endpoint.
type TickGenericJQ struct {
	http *TickGenericHTTP

	rawQuery string
	query    *gojq.Code
	logger   log.Logger
}

// NewTickGenericJQ creates a new TickGenericJQ instance.
//
// The client argument is an TickGenericHTTP client that is used to fetch data from the
// TickGenericHTTP endpoint.
//
// The url argument is an TickGenericHTTP endpoint that returns JSON data. It may contain
// the following variables:
//   - ${lcbase} - lower case base asset
//   - ${ucbase} - upper case base asset
//   - ${lcquote} - lower case quote asset
//   - ${ucquote} - upper case quote asset
//   - ${lcbases} - lower case base assets joined by commas
//   - ${ucbases} - upper case base assets joined by commas
//   - ${lcquotes} - lower case quote assets joined by commas
//   - ${ucquotes} - upper case quote assets joined by commas
//
// The jq argument is a JQ query that is used to parse JSON data. It must
// return a single value that will be used as a price or an object with the
// following fields:
//   - price - a price
//   - time - a timestamp (optional)
//   - volume - a 24h volume (optional)
//
// The JQ query may contain the following variables:
//   - $lcbase - lower case base asset
//   - $ucbase - upper case base asset
//   - $lcquote - lower case quote asset
//   - $ucquote - upper case quote asset
//
// Price and volume must be a string that can be parsed as a number or a number.
//
// Time must be a string that can be parsed as time or a number that represents
// a UNIX timestamp.
//
// If JQ query returns multiple values, the dataPoint will be invalid.
func NewTickGenericJQ(opts TickGenericJQOptions) (*TickGenericJQ, error) {
	if opts.URL == "" {
		return nil, fmt.Errorf("url cannot be empty")
	}
	if opts.Query == "" {
		return nil, fmt.Errorf("query must be specified")
	}
	if opts.Client == nil {
		opts.Client = http.DefaultClient
	}
	if opts.Logger == nil {
		opts.Logger = null.New()
	}
	parsed, err := gojq.Parse(opts.Query)
	if err != nil {
		return nil, err
	}
	compiled, err := gojq.Compile(parsed, gojq.WithVariables([]string{
		"$lcbase",
		"$ucbase",
		"$lcquote",
		"$ucquote",
	}))
	if err != nil {
		return nil, err
	}
	jq := &TickGenericJQ{}
	gh, err := NewTickGenericHTTP(TickGenericHTTPOptions{
		URL:      opts.URL,
		Headers:  opts.Headers,
		Callback: jq.handle,
		Client:   opts.Client,
		Logger:   opts.Logger,
	})
	if err != nil {
		return nil, err
	}
	jq.http = gh
	jq.rawQuery = opts.Query
	jq.query = compiled
	jq.logger = opts.Logger.WithField("tag", TickGenericJQLoggerTag)
	return jq, nil
}

// FetchDataPoints implements the Origin interface.
func (g *TickGenericJQ) FetchDataPoints(ctx context.Context, query []any) (map[any]datapoint.Point, error) {
	return g.http.FetchDataPoints(ctx, query)
}

//nolint:funlen
func (g *TickGenericJQ) handle(ctx context.Context, pairs []value.Pair, body io.Reader) map[any]datapoint.Point {
	points := make(map[any]datapoint.Point)

	// Parse JSON data.
	var decoded any
	if err := json.NewDecoder(body).Decode(&decoded); err != nil {
		for _, pair := range pairs {
			points[pair] = datapoint.Point{Error: err}
		}
		return points
	}

	// Run JQ query for each pair and parse the result.
	for _, pair := range pairs {
		g.logger.
			WithFields(log.Fields{
				"url":   g.http.url,
				"query": g.rawQuery,
				"pairs": pairs,
			}).
			Debug("JQ request")

		point := datapoint.Point{Time: time.Now()}
		tick := value.Tick{Pair: pair}
		iter := g.query.RunWithContext(
			ctx,
			decoded,
			strings.ToLower(pair.Base),  // $lcbase
			strings.ToUpper(pair.Base),  // $ucbase
			strings.ToLower(pair.Quote), // $lcquote
			strings.ToUpper(pair.Quote), // $ucquote
		)
		v, ok := iter.Next()
		if !ok {
			point.Value = tick
			point.Error = fmt.Errorf("no result from JQ query")
			points[pair] = point
			continue
		}
		if err, ok := v.(error); ok {
			point.Value = tick
			point.Error = err
			points[pair] = point
			continue
		}
		if _, ok := iter.Next(); ok {
			point.Value = tick
			point.Error = fmt.Errorf("multiple results from JQ query")
			points[pair] = point
			continue
		}
		switch v := v.(type) {
		case map[string]any:
			for k, v := range v {
				switch k {
				case "price":
					tick.Price = bn.Float(v)
				case "volume":
					tick.Volume24h = bn.Float(v)
				case "time":
					if tm, ok := anyToTime(v); ok {
						point.Time = tm
					}
				default:
					point.Error = fmt.Errorf("unknown key in JQ result: %s", k)
				}
			}
		case int, int32, int64, uint, uint32, uint64, float32, float64:
			tick.Price = bn.Float(v)
		}
		point.Value = tick
		points[pair] = point
	}
	return points
}

// anyToTime converts an arbitrary value to a time.Time.
func anyToTime(v any) (time.Time, bool) {
	switch v := v.(type) {
	case string:
		for _, layout := range []string{
			time.RFC3339,
			time.RFC3339Nano,
			time.RFC1123,
			time.RFC1123Z,
			time.RFC822,
			time.RFC822Z,
			time.RFC850,
			time.ANSIC,
			time.UnixDate,
			time.RubyDate,
		} {
			t, err := time.Parse(layout, v)
			if err == nil {
				return t, true
			}
		}
	case int, int32, int64:
		return time.Unix(v.(int64), 0), true
	case uint, uint32, uint64:
		return time.Unix(int64(v.(uint64)), 0), true
	case float32, float64:
		return time.Unix(int64(v.(float64)), 0), true
	}
	return time.Time{}, false
}
