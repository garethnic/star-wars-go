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

	fmt.Println("OPTIONS:")
	fmt.Println("1. Films")
	fmt.Println("2. People")
	fmt.Scan(&category)

	switch category {
	case 1:
		fmt.Println("Enter the film ID")
		fmt.Scan(&choice)
		film.GetFilm(choice)
		Menu()
	case 2:
		fmt.Println("Enter the person ID")
		fmt.Scan(&choice)
		person.GetPerson(choice)
		Menu()
	case -1:
		fmt.Println("Goodbye")
		os.Exit(1)
	default:
		Menu()
	}
}
