package driver

import (
	"context"
	"time"

	"github.com/bangumi/server/config"
	"github.com/bangumi/server/internal/pkg/logger"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/trim21/errgo"
)

func NewElasticSearch(c config.AppConfig) (*elasticsearch.TypedClient, error) {
	if c.Search.ElasticSearch.URL == "" {
		logger.Warn("ElasticSearch endpoint is not specified. ElasticSearch API will be unavailable.")
		return nil, nil
	}

	cfg := elasticsearch.Config{
		Addresses: []string{c.Search.ElasticSearch.URL},
		APIKey:    c.Search.ElasticSearch.Key,
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		errgo.Wrap(err, "elasticsearch: Failed to create ElasticSearch client.")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ok, err := es.Ping().Do(ctx)
	if !ok {
		errgo.Wrap(err, "elasticsearch: Error pinging the Elasticsearch server")
		return nil, err
	}

	return es, nil
}
