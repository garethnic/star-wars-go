package person

import (
	"encoding/json"
	"fmt"
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
func GetPerson(id int) {
	var person Person

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
	json.Unmarshal(content, &person)

	fmt.Println("Found: ")
	fmt.Println(person.Name, person.Gender)
}
