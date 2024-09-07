package model

type SearchMetaData struct {
	Total    int64
	Took     int64
	Timedout bool
	ScrollID string
}

type SearchResult struct {
	Metadata    SearchMetaData
	Celebrities []Celebrity
	Subjects    []Subject
}
