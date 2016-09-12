package film

import (
	"net/url"
	"log"
	"path"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

//Our Film object
type Film struct {
	Title string `json:"title"`
	Episode int `json:"episode_id"`
	OpeningCrawl string `json:"opening_crawl"`
}

//Get a film based on ID
func GetFilm(id int) {
	var film Film

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

	//Save query into object
	json.Unmarshal(content, &film)

	fmt.Println("Found: ")
	fmt.Println(film.Title, film.Episode)
	fmt.Println(film.OpeningCrawl)
}
