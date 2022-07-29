package main

import "fmt"

type Sports struct {
	firstName string
	lastName  string
	team      string
}

func GetPlayersFromBothSports(s1, s2 []Sports) []string {
	// iterate through s1
	// -- create a map of firstName + lastName for s1, setting to true
	// -- end iter
	// init an empty `names` []string
	// iterate through s2
	// -- for each firstName + lastName, access map,
	// -- if exists, append to names, else ignore
	// -- end iteration
	// return `names`

	myMap := make(map[string]bool)

	for _, s := range s1 {
		fullName := s.firstName + " " + s.lastName
		myMap[fullName] = true
	}

	names := make([]string, 0, len(s2)) // optimized cap
	for _, s := range s2 {
		fullName := s.firstName + " " + s.lastName
		if _, ok := myMap[fullName]; ok {
			names = append(names, fullName)
		}
	}
	return names
} // O(N + M)

func main() {
	basketballPlayers := []Sports{
		{firstName: "Jill", lastName: "Huang", team: "Gators"}, // Sports{} can be removed
		{firstName: "Janko", lastName: "Barton", team: "Sharks"},
		{firstName: "Wanda", lastName: "Valkulskas", team: "Sharks"},
		{firstName: "Jill", lastName: "Moloney", team: "Gators"},
		{firstName: "Luuk", lastName: "Watkins", team: "Gators"},
	}

	footballPlayers := []Sports{
		{firstName: "Hanzla", lastName: "Radozi", team: "Gators"},
		{firstName: "Tina", lastName: "Watkins", team: "Sharks"},
		{firstName: "Alex", lastName: "Patel", team: "Sharks"},
		{firstName: "Jill", lastName: "Huang", team: "Gators"},
		{firstName: "Wanda", lastName: "Valkulskas", team: "Gators"},
	}

	fmt.Println(GetPlayersFromBothSports(basketballPlayers, footballPlayers))
}
