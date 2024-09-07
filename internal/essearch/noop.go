package essearch

import (
	"context"
	"errors"

	"github.com/bangumi/server/internal/model"
)

var _ Repo = NoopRepo{}

type NoopRepo struct {
}

func (n NoopRepo) AutoComplete(ctx context.Context, prefix string, category string) ([]string, error) {
	return nil, errors.New("essearch not enabled")
}

func (n NoopRepo) Celebrity(ctx context.Context, query string, typ string) (model.SearchResult, error) {
	searchResult := model.SearchResult{}
	return searchResult, errors.New("essearch not enabled")
}

func (n NoopRepo) Subject(ctx context.Context, query string, typ string, tags []string, dateRange map[string]string) (model.SearchResult, error) {
	searchResult := model.SearchResult{}
	return searchResult, errors.New("essearch not enabled")
}

func (n NoopRepo) Scroll(ctx context.Context, scrollId string) (model.SearchResult, error) {
	searchResult := model.SearchResult{}
	return searchResult, errors.New("essearch not enabled")
}
