package essearch

import (
	"context"

	"github.com/bangumi/server/internal/model"
)

const CELEBRITY_INDEX = "celebrity_v5"
const SUBJECT_INDEX = "subject_v2"

type Repo interface {
	AutoComplete(ctx context.Context, prefix string, category string) ([]string, error)
	Celebrity(ctx context.Context, query string, typ string) (model.SearchResult, error)
	Subject(ctx context.Context, query string, typ string, tags []string, dateRange map[string]string) (model.SearchResult, error)
	Scroll(ctx context.Context, scrollId string) (model.SearchResult, error)
}
