package xbmctoollib

//The searcher interface is just in case I
// a) test this code
// b) Want to switch out, create a different search provider
type Searcher interface {
	Search(q string) (ResultSet, error)
}

//the resultset interface makes it easier to work with results
//that may come in different forms.
type ResultSet interface {
	Hits() int
	GetResult(i int) (url, title, release string, ok bool)
}
