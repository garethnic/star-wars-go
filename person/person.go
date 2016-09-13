package person

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

//Our Person object
type Person struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

//Get a person based on id
func (p *Person) GetPerson(id int) Person {
	//Base uri
	call, err := url.Parse("http://swapi.co/api/people/")
	if err != nil {
		log.Fatal(err)
	}

	//Add ID to search for
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
	json.Unmarshal(content, &p)

	return *p
}

//Get Person name
func (p Person) GetName() string {
	return p.Name
}

//Get Person gender
func (p Person) GetGender() string {
	return p.Gender
}
