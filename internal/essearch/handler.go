package essearch

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/bangumi/server/dal/query"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/generic/slice"
	"github.com/bangumi/server/internal/pkg/null"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/childscoremode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
)

type esRepo struct {
	q   *query.Query
	es  *elasticsearch.TypedClient
	log *zap.Logger
}

func NewElasticSearchRepo(q *query.Query, es *elasticsearch.TypedClient, log *zap.Logger) (Repo, error) {
	if es == nil || q == nil {
		return NoopRepo{}, nil
	}
	return esRepo{q: q, es: es, log: log.Named("essearch.Repo")}, nil
}

func (r esRepo) AutoComplete(ctx context.Context, prefix string, category string) ([]string, error) {
	var index string
	if category == "celebrity" || category == "person" || category == "character" {
		index = CELEBRITY_INDEX
	} else {
		index = SUBJECT_INDEX
	}

	suggester := types.Suggester{
		Suggesters: map[string]types.FieldSuggester{
			"autocomplete": {
				Prefix: null.NilString(prefix),
				Completion: &types.CompletionSuggester{
					Size:           null.NilInt(10),
					SkipDuplicates: null.NilBool(true),
					Field:          "suggest",
					Contexts: map[string][]types.CompletionContext{
						"metatype": {types.CompletionContext{Context: category}},
					},
				},
			},
		},
	}

	res, err := r.es.Search().Index(index).Suggest(&suggester).Source_(false).Do(ctx)

	if err != nil {
		return nil, err
	}

	var suggestions []string = make([]string, 0)
	if res.Suggest != nil && res.Suggest["autocomplete"] != nil {
		optionsVar := reflect.ValueOf(res.Suggest["autocomplete"][0]).Elem().FieldByName("Options")
		size := optionsVar.Len()
		for idx := 0; idx < size && idx < 10; idx++ {
			opt := optionsVar.Index(idx)
			suggestions = append(suggestions, opt.FieldByName("Text").String())
		}
	}

	return suggestions, nil
}

func (r esRepo) Celebrity(ctx context.Context, query string, typ string) (model.SearchResult, error) {
	q := types.Query{
		Bool: &types.BoolQuery{
			Should: []types.Query{
				{
					Prefix: map[string]types.PrefixQuery{
						"name.keyword": {
							Value: query,
							Boost: null.NilFloat32(10),
						},
					},
				},
				{
					Prefix: map[string]types.PrefixQuery{
						"alias.keyword": {
							Value: query,
							Boost: null.NilFloat32(10),
						},
					},
				},
				{
					MultiMatch: &types.MultiMatchQuery{
						Query:  query,
						Type:   &textquerytype.Phrase,
						Fields: []string{"name", "alias", "name.jp", "alias.jp"},
						Boost:  null.NilFloat32(5),
					},
				},
				{
					MultiMatch: &types.MultiMatchQuery{
						Query:  query,
						Type:   &textquerytype.Phrase,
						Fields: []string{"name.pinyin", "alias.pinyin", "name.romaji", "alias.romaji"},
						Boost:  null.NilFloat32(2),
					},
				},
				{
					MultiMatch: &types.MultiMatchQuery{
						Query:         query,
						Fuzziness:     "AUTO",
						MaxExpansions: null.NilInt(10),
						PrefixLength:  null.NilInt(3),
						Fields:        []string{"name", "alias", "name.pinyin", "alias.pinyin", "name.jp", "alias.jp", "name.romaji", "alias.romaji"},
					},
				},
			},
		},
	}

	if typ != "celebrity" {
		q.Bool.Filter = []types.Query{
			{
				Term: map[string]types.TermQuery{
					"type": {
						Value: typ,
					},
				},
			},
		}
	}

	sort := types.Sort{types.SortOptions{
		Score_: &types.ScoreSort{
			Order: &sortorder.Desc,
		},
		SortOptions: map[string]types.FieldSort{
			"score": {
				Order: &sortorder.Desc,
			},
		},
	}}

	searchResult := model.SearchResult{}

	res, err := r.es.Search().Index(CELEBRITY_INDEX).Query(&q).Sort(sort...).MinScore(15).Scroll("3m").RestTotalHitsAsInt(true).Source_(false).Do(ctx)
	if err != nil {
		return searchResult, err
	}

	r.populateCelebrities(&searchResult, res)

	return searchResult, nil
}

func (r esRepo) Subject(ctx context.Context, query string, typ string, tags []string, dateRange map[string]string) (model.SearchResult, error) {
	searchResult := model.SearchResult{}

	var shouldQuery []types.Query
	if query != "" {
		shouldQuery = []types.Query{
			{
				Prefix: map[string]types.PrefixQuery{
					"name.keyword": {
						Value: query,
						Boost: null.NilFloat32(10),
					},
				},
			},
			{
				Prefix: map[string]types.PrefixQuery{
					"name_cn.keyword": {
						Value: query,
						Boost: null.NilFloat32(10),
					},
				},
			},
			{
				Prefix: map[string]types.PrefixQuery{
					"alias.keyword": {
						Value: query,
						Boost: null.NilFloat32(10),
					},
				},
			},
			{
				Nested: &types.NestedQuery{
					Query: &types.Query{
						Bool: &types.BoolQuery{
							Should: []types.Query{
								{
									Match: map[string]types.MatchQuery{
										"tags.tags": {
											Query: query,
										},
									},
								},
								{
									Term: map[string]types.TermQuery{
										"tags.tags.keyword": {
											Value: query,
										},
									},
								},
							},
						},
					},
					ScoreMode: &childscoremode.Max,
					Path:      "tags",
				},
			},
			{
				MultiMatch: &types.MultiMatchQuery{
					Query:  query,
					Type:   &textquerytype.Phrase,
					Fields: []string{"name", "name_cn", "alias", "name.jp"},
					Boost:  null.NilFloat32(5),
				},
			},
			{
				MultiMatch: &types.MultiMatchQuery{
					Query:  query,
					Type:   &textquerytype.Phrase,
					Fields: []string{"name.pinyin", "name_cn.pinyin", "alias.pinyin", "name.romaji"},
					Boost:  null.NilFloat32(2),
				},
			},
			{
				MultiMatch: &types.MultiMatchQuery{
					Query:         query,
					Fuzziness:     "AUTO",
					MaxExpansions: null.NilInt(10),
					PrefixLength:  null.NilInt(3),
					Fields:        []string{"name", "name_cn", "alias"},
				},
			},
			{
				MultiMatch: &types.MultiMatchQuery{
					Query:  query,
					Fields: []string{"summary", "summary.cn", "summary.jp"},
					Boost:  null.NilFloat32(0.1),
				},
			},
		}
	}

	typeTerm := types.NewTermQuery()
	typeTerm.Value = typ
	typeFilter := types.NewQuery()
	typeFilter.Term["type"] = *typeTerm
	dateRangeQuery := types.NewDateRangeQuery()
	if lte, ok := dateRange["lte"]; ok {
		dateRangeQuery.Lte = new(string)
		*dateRangeQuery.Lte = lte
	}
	if gte, ok := dateRange["gte"]; ok {
		dateRangeQuery.Gte = new(string)
		*dateRangeQuery.Gte = gte
	}
	dateFilter := types.NewQuery()
	dateFilter.Range["date"] = *dateRangeQuery

	mustQuery := slice.Map(tags, func(t string) types.Query {
		return types.Query{
			Nested: &types.NestedQuery{
				Path: "tags",
				Query: &types.Query{
					MatchPhrase: map[string]types.MatchPhraseQuery{
						"tags.tags": {
							Query: t,
						},
					},
				},
			},
		}
	})

	boolQuery := types.NewBoolQuery()
	if len(tags) > 0 {
		boolQuery.Must = mustQuery
	}
	boolQuery.Should = shouldQuery
	if len(dateRange) > 0 {
		boolQuery.Filter = append(boolQuery.Filter, *dateFilter)
	}
	if typ != "subject" {
		boolQuery.Filter = append(boolQuery.Filter, *typeFilter)
	}
	q := types.NewQuery()
	q.Bool = boolQuery

	sort := types.Sort{types.SortOptions{
		Score_: &types.ScoreSort{
			Order: &sortorder.Desc,
		},
	}}
	if len(tags) > 0 {
		escapedTags := slice.Map(tags, escapeSpecialCharacters)
		tagField := types.FieldSort{
			Order: &sortorder.Desc,
			Mode:  &sortmode.Avg,
			Nested: &types.NestedSortValue{
				Path: "tags",
				Filter: &types.Query{
					QueryString: &types.QueryStringQuery{
						Query:        strings.Join(escapedTags, " OR "),
						DefaultField: null.NilString("tags.tags"),
					},
				},
			},
		}

		popularityCond := types.SortOptions{
			SortOptions: map[string]types.FieldSort{
				"tags.tag_cnt": tagField,
			},
		}
		confidenceCond := types.SortOptions{
			SortOptions: map[string]types.FieldSort{
				"tags.confidence": tagField,
			},
		}
		sort = append([]types.SortCombinations{popularityCond, confidenceCond}, sort...)
	}

	res, err := r.es.Search().Index(SUBJECT_INDEX).Query(q).Sort(sort...).MinScore(1).Scroll("3m").RestTotalHitsAsInt(true).Source_(false).Do(ctx)

	if err != nil {
		return searchResult, err
	}

	r.populateSubjects(&searchResult, res)

	return searchResult, nil
}

func (r esRepo) Scroll(ctx context.Context, scrollId string) (model.SearchResult, error) {
	searchResult := model.SearchResult{}
	res, err := r.es.Scroll().ScrollId(scrollId).Scroll("3m").RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		return searchResult, err
	}

	if len(res.Hits.Hits) > 0 {
		index := res.Hits.Hits[0].Index_
		if index == CELEBRITY_INDEX {
			r.populateCelebrities(&searchResult, &search.Response{
				Hits:      res.Hits,
				Took:      res.Took,
				TimedOut:  res.TimedOut,
				ScrollId_: res.ScrollId_,
			})
		} else {
			r.populateSubjects(&searchResult, &search.Response{
				Hits:      res.Hits,
				Took:      res.Took,
				TimedOut:  res.TimedOut,
				ScrollId_: res.ScrollId_,
			})
		}
	} else {
		searchResult.Metadata = model.SearchMetaData{
			Total:    res.Hits.Total.Value,
			Took:     res.Took,
			Timedout: res.TimedOut,
			ScrollID: *res.ScrollId_,
		}
	}
	return searchResult, nil
}

func escapeSpecialCharacters(input string) string {
	// Create a regular expression to match special characters.
	specialCharactersRegex := regexp.MustCompile(`[+\-=&|><!()\[\]{}^"~*?:/]`)

	// Replace all special characters with their escaped versions.
	output := specialCharactersRegex.ReplaceAllString(input, "\\$1")
	return output
}

func (r esRepo) populateSubjects(serd *model.SearchResult, res *search.Response) {
	serd.Metadata = model.SearchMetaData{
		Total:    res.Hits.Total.Value,
		Took:     res.Took,
		Timedout: res.TimedOut,
		ScrollID: *res.ScrollId_,
	}
	subjects := []model.Subject{}
	if len(res.Hits.Hits) > 0 {
		for i := 0; i < len(res.Hits.Hits); i++ {
			idstr := res.Hits.Hits[i].Id_
			id, err := strconv.ParseUint(idstr, 10, 32)
			if err != nil {
				r.log.Warn(fmt.Sprintf("Cannot parse subject id %s", idstr))
				continue
			}
			subjects = append(subjects, model.Subject{ID: uint32(id)})
		}
	}
	serd.Subjects = subjects
}

func (r esRepo) populateCelebrities(serd *model.SearchResult, res *search.Response) {
	serd.Metadata = model.SearchMetaData{
		Total:    res.Hits.Total.Value,
		Took:     res.Took,
		Timedout: res.TimedOut,
		ScrollID: *res.ScrollId_,
	}
	celebrities := []model.Celebrity{}
	if len(res.Hits.Hits) > 0 {
		for i := 0; i < len(res.Hits.Hits); i++ {
			idstr := res.Hits.Hits[i].Id_
			id, err := strconv.ParseUint(strings.Split(idstr, "_")[1], 10, 32)
			if err != nil {
				r.log.Warn(fmt.Sprintf("Cannot parse celebrity id %s", idstr))
				continue
			}
			if strings.Split(idstr, "_")[0] == "person" {
				celebrities = append(celebrities, model.Celebrity{Category: 0, ID: uint32(id)})
			} else {
				celebrities = append(celebrities, model.Celebrity{Category: 0, ID: uint32(id)})
			}
		}
	}
	serd.Celebrities = celebrities
}
