package req

type SubjectQuery struct {
	Query     string       `json:"query"`
	Type      string       `json:"type"`
	Tags      []string     `json:"tags"`
	DateRange DateRangeDTO `json:"dateRange"`
}

type DateRangeDTO struct {
	Gte string `json:"gte"`
	Lte string `json:"lte"`
}

type CelebrityQuery struct {
	Query string `json:"query"`
	Type  string `json:"type"`
}
