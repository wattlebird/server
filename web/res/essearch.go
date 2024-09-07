package res

import (
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/compat"
	"github.com/bangumi/server/internal/pkg/null"
	"github.com/bangumi/server/pkg/wiki"
)

type CelebrityRes struct {
	Category  uint8        `json:"category"`
	Images    PersonImages `json:"images"`
	Summary   string       `json:"summary"`
	Name      string       `json:"name"`
	Infobox   v0wiki       `json:"infobox"`
	Stat      Stat         `json:"stat"`
	ID        uint32       `json:"id"`
	Locked    bool         `json:"locked"`
	Type      uint8        `json:"type"`
	BloodType *uint8       `json:"blood_type"`
	BirthYear *uint16      `json:"birth_year"`
	BirthDay  *uint8       `json:"birth_day"`
	BirthMon  *uint8       `json:"birth_mon"`
	Gender    *string      `json:"gender"`
}

type SearchResult struct {
	Total    int64       `json:"total"`
	Took     int64       `json:"took"`
	Timedout bool        `json:"timedout"`
	ScrollID string      `json:"scrollid"`
	Data     interface{} `json:"data"`
}

func ConvertModelPersonToCelebrity(s model.Person) CelebrityRes {
	img := PersonImage(s.Image)

	return CelebrityRes{
		Category: 0,
		ID:       s.ID,
		Type:     s.Type,
		Name:     s.Name,
		Images:   img,
		Summary:  s.Summary,
		Infobox:  compat.V0Wiki(wiki.ParseOmitError(s.Infobox).NonZero()),
		Stat: Stat{
			Comments: s.CommentCount,
			Collects: s.CollectCount,
		},
		Locked:    s.Locked,
		Gender:    null.NilString(GenderMap[s.FieldGender]),
		BloodType: null.NilUint8(s.FieldBloodType),
		BirthYear: null.NilUint16(s.FieldBirthYear),
		BirthMon:  null.NilUint8(s.FieldBirthMon),
		BirthDay:  null.NilUint8(s.FieldBirthDay),
	}
}

func ConvertModelCharacterToCelebrity(s model.Character) CelebrityRes {
	img := PersonImage(s.Image)

	return CelebrityRes{
		Category: 1,
		ID:       s.ID,
		Type:     s.Type,
		Name:     s.Name,
		Images:   img,
		Summary:  s.Summary,
		Infobox:  compat.V0Wiki(wiki.ParseOmitError(s.Infobox).NonZero()),
		Stat: Stat{
			Comments: s.CommentCount,
			Collects: s.CollectCount,
		},
		Locked:    s.Locked,
		Gender:    null.NilString(GenderMap[s.FieldGender]),
		BloodType: null.NilUint8(s.FieldBloodType),
		BirthYear: null.NilUint16(s.FieldBirthYear),
		BirthMon:  null.NilUint8(s.FieldBirthMon),
		BirthDay:  null.NilUint8(s.FieldBirthDay),
	}
}

func ConvertModelSubject(s model.Subject) SubjectV0 {
	return SubjectV0{
		ID:      s.ID,
		Name:    s.Name,
		NameCN:  s.NameCN,
		Date:    null.NilString(s.Date),
		Summary: s.Summary,
		Image:   SubjectImage(s.Image),
		Eps:     s.Eps,
		Volumes: s.Volumes,
		Infobox: compat.V0Wiki(wiki.ParseOmitError(s.Infobox).NonZero()),
		Rating: Rating{
			Rank:  s.Rating.Rank,
			Score: s.Rating.Score,
			Count: Count(s.Rating.Count),
			Total: s.Rating.Total,
		},
		Locked: s.Locked(),
		NSFW:   s.NSFW,
		TypeID: s.TypeID,
	}
}
