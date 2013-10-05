package xbmctoollib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	TMDB_URL_PATTERN = "http://api.themoviedb.org/3/search/movie?api_key=%s&query=%s"
	TMDB_MOVIE_URL   = "http://themoviedb.org/movie/%d"
)

type TmdbResult struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Release string `json:"release_date,omitempty"`
}

type TmdbResponse struct {
	Results []*TmdbResult `json:"results"`
}

type TmdbMovieSearch struct {
	api_key string
}

func NewTmdbMovieSearch(key string) *TmdbMovieSearch {
	return &TmdbMovieSearch{api_key: key}
}

// this satisfies the Searcher interface
func (t *TmdbMovieSearch) Search(q string) (res ResultSet, err error) {
	var r *http.Response
	url := fmt.Sprintf(TMDB_URL_PATTERN, t.api_key, url.QueryEscape(q))
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("accept", "application/json")
	if r, err = http.DefaultClient.Do(req); err != nil {
		return nil, fmt.Errorf("HTTPError: %s", err)
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadResponseError: %s", err)
	}
	res = &TmdbResponse{}
	if err = json.Unmarshal(b, res); err != nil {
		return nil, fmt.Errorf("JSONError: %s", err)
	}
	return res, nil
}

//this satisfies the ResultSet interface
func (tr *TmdbResponse) Hits() int {
	return len(tr.Results)
}

//this satisfies the ResultSet interface
func (tr *TmdbResponse) GetResult(i int) (url, title, release string, ok bool) {
	if i < 0 || i >= len(tr.Results) {
		panic("Result Index Out of Bounds!")
	}
	r := tr.Results[i]
	if r == nil {
		return "", "", "", false
	}
	return fmt.Sprintf(TMDB_MOVIE_URL, r.Id), r.Title, r.Release, true
}
