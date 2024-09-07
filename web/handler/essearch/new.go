package essearch

import (
	"go.uber.org/zap"

	"github.com/bangumi/server/internal/character"
	"github.com/bangumi/server/internal/essearch"
	"github.com/bangumi/server/internal/person"
	"github.com/bangumi/server/internal/pkg/cache"
	"github.com/bangumi/server/internal/subject"
	"github.com/bangumi/server/web/handler/common"
)

type Handler struct {
	common    common.Common
	cache     cache.RedisCache
	log       *zap.Logger
	person    person.Repo
	character character.Repo
	subject   subject.Repo
	search    essearch.Repo
}

func New(
	search essearch.Repo,
	person person.Repo,
	character character.Repo,
	subject subject.Repo,
	common common.Common,
	cache cache.RedisCache,
	log *zap.Logger,
) Handler {
	return Handler{
		search:    search,
		person:    person,
		character: character,
		subject:   subject,
		common:    common,
		cache:     cache,
		log:       log.Named("web.handler.essearch"),
	}
}
