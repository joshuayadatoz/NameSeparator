package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	Title     string
	FirstName string
	Initial   string
	LastName  string
}

func ParseName(name string) []Person {
	titles := []string{"Mr", "Mrs", "Mister", "Dr", "Ms", "Prof"}
	nameSplitters := []string{" and ", " & "}

	var persons []Person

	// Check for multiple homeowners
	for _, splitter := range nameSplitters {
		if strings.Contains(name, splitter) {
			singleName := strings.Split(name, splitter)
			sharedSurname := ""
			var tempPersons []Person
			//Start from the end to catch the possible surname from n+1 person
			for i := len(singleName) - 1; i >= 0; i-- {
				//Check if the first person has a lastname, if not assign the second person's last name
				wordCountInName := len(strings.Split(singleName[i], " "))
				if wordCountInName > 1 {
					sharedSurname = strings.Split(singleName[i], " ")[wordCountInName-1]
				} else {
					singleName[i] = singleName[i] + " " + sharedSurname
				}
				tempPersons = append(persons, ParseName(strings.TrimSpace(singleName[i]))...) // Process each homeowner separately
			}
			for i := len(tempPersons) - 1; i >= 0; i-- {
				persons = append(persons, tempPersons[i]) // Process each homeowner separately
			}

			return persons
		}
	}

	// Processing a single homeowner
	for _, title := range titles {
		if strings.HasPrefix(name, title+" ") {
			var person Person
			person.Title = title
			splitName := strings.SplitN(name, title, 2)
			if len(splitName) > 1 {
				name = strings.TrimSpace(splitName[1])
			}

			if name == "" {
				persons = append(persons, person)
				continue
			}

			if strings.Contains(name, ".") {
				splits := strings.Split(name, ".")
				if len(splits) > 0 {
					person.Initial = strings.TrimSpace(splits[0])
				}
				if len(splits) > 1 {
					person.LastName = strings.TrimSpace(splits[1])
				}
			} else {
				splits := strings.Fields(name)
				if len(splits) > 0 {
					switch len(splits) {
					case 1:
						person.LastName = splits[0]
					case 2:
						person.FirstName = splits[0]
						person.LastName = splits[1]
					default:
						person.FirstName = splits[0]
						person.LastName = strings.Join(splits[1:], " ")
					}
				}
			}
			persons = append(persons, person)
			break
		}
	}

	return persons
}

func main() {
	csvfile, err := os.Open("homeowner.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	r.FieldsPerRecord = -1

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		if len(record) > 0 {
			persons := ParseName(record[0])
			for _, person := range persons {
				fmt.Printf("Title: %s, FirstName: %s, Initial: %s, LastName: %s\n", person.Title, person.FirstName, person.Initial, person.LastName)
			}
		}
	}
}
