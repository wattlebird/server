package essearch

import (
	"context"

	"github.com/bangumi/server/internal/pkg/generic/slice"
	"github.com/bangumi/server/internal/subject"
	"github.com/trim21/errgo"
	"go.uber.org/zap"

	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/null"
	"github.com/bangumi/server/web/accessor"
	"github.com/bangumi/server/web/res"
)

func (h Handler) extractSubject(ctx context.Context, r *model.SearchResult, user *accessor.Accessor) ([]res.SubjectV0, error) {
	subjectIdList := slice.Map(r.Subjects, func(sub model.Subject) uint32 {
		return sub.ID
	})
	subjectMap, err := h.subject.GetByIDs(ctx, subjectIdList, subject.Filter{NSFW: null.Bool{Value: false, Set: !user.AllowNSFW()}})
	if err != nil {
		h.log.Error("Get subject by IDs failed", zap.Any("IDs", subjectIdList), zap.Error(err))
		return nil, errgo.Wrap(err, "Get subject by IDs failed")
	}

	data := []res.SubjectV0{}
	for _, sub := range r.Subjects {
		if subject, ok := subjectMap[sub.ID]; ok {
			data = append(data, res.ConvertModelSubject(subject))
		}
	}
	return data, nil
}

func (h Handler) extractCelebrity(ctx context.Context, r *model.SearchResult) ([]res.CelebrityRes, error) {
	personIdList := slice.MapFilter(r.Celebrities, func(celeb model.Celebrity) (uint32, bool) {
		return celeb.ID, celeb.Category == 0
	})
	characterIdList := slice.MapFilter(r.Celebrities, func(celeb model.Celebrity) (uint32, bool) {
		return celeb.ID, celeb.Category == 1
	})

	personMap, err := h.person.GetByIDs(ctx, personIdList)
	if err != nil {
		h.log.Error("Get person by IDs failed", zap.Any("IDs", personIdList), zap.Error(err))
		return nil, errgo.Wrap(err, "Get person by IDs failed")
	}
	characterMap, err := h.character.GetByIDs(ctx, characterIdList)
	if err != nil {
		h.log.Error("Get character by IDs failed", zap.Any("IDs", characterIdList), zap.Error(err))
		return nil, errgo.Wrap(err, "Get character by IDs failed")
	}
	data := []res.CelebrityRes{}

	for _, celeb := range r.Celebrities {
		if celeb.Category == 0 {
			if person, ok := personMap[celeb.ID]; ok {
				data = append(data, res.ConvertModelPersonToCelebrity(person))
			}
		} else if celeb.Category == 1 {
			if character, ok := characterMap[celeb.ID]; ok {
				data = append(data, res.ConvertModelCharacterToCelebrity(character))
			}
		}
	}
	return data, nil
}
