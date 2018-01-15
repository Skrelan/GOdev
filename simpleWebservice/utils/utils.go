package utils

import (
	"github.com/skrelan/GOdev/simpleWebservice/models"
)

func createPerson(s []string) *models.Person {
	return &models.Person{
		ID:        s[0],
		FirstName: s[1],
		LastName:  s[2],
		Address: &models.Address{
			City:  s[3],
			State: s[4],
		}}
}

func CreateData(p *[]models.Person) {
	rawData := [][]string{
		{"1", "James", "Bond", "San Jose", "CA"},
		{"2", "Super", "Man", "Los Angeles", "CA"},
		{"3", "Jack", "Daniels", "Kentukky", "TN"}}
	for _, data := range rawData {
		*p = append(*p, *createPerson(data))
	}
}
