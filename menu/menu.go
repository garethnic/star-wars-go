package menu

import (
	"fmt"
	"os"
	"star-wars-go/film"
	"star-wars-go/person"
)

func Menu() {
	var category int
	var choice int

	fmt.Println("OPTIONS:\n1. Films\n2. People")
	fmt.Scan(&category)

	switch category {
	case 1:
		fmt.Println("Enter the film ID")
		fmt.Scan(&choice)
		film := film.Film{}
		film.GetFilm(choice)
		fmt.Printf("Found: \n%s, %d,\n%s\n", film.GetTitle(), film.GetEpisode(), film.GetOpeningCrawl())
		Menu()
	case 2:
		fmt.Println("Enter the person ID")
		fmt.Scan(&choice)
		person := person.Person{}
		person.GetPerson(choice)
		fmt.Printf("Found: \n%s, %s\n", person.GetName(), person.GetGender())
		Menu()
	case -1:
		fmt.Println("Goodbye")
		os.Exit(1)
	default:
		Menu()
	}
}
