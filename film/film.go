package film

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

//Our Film object
type Film struct {
	Title        string `json:"title"`
	Episode      int    `json:"episode_id"`
	OpeningCrawl string `json:"opening_crawl"`
}

//Get a film based on ID
func (f *Film) GetFilm(id int) Film {
	//Base uri
	call, err := url.Parse("http://swapi.co/api/films")
	if err != nil {
		log.Fatal(err)
	}

	//Add ID to search path
	call.Path = path.Join(call.Path, strconv.Itoa(id))

	//Perform query
	resp, err := http.Get(call.String())
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//Save query into object1
	json.Unmarshal(content, &f)

	return *f
}

//Get Film title
func (f Film) GetTitle() string {
	return f.Title
}

//Get Film episode
func (f Film) GetEpisode() int {
	return f.Episode
}

//Get Film opening crawl
func (f Film) GetOpeningCrawl() string {
	return f.OpeningCrawl
}
