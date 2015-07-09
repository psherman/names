// Generate a random male or female (or either) name
// Uses the top 100 male/female names from 1915-2014
// from http://www.ssa.gov/oact/babynames/decades/century.html
// and the top 100 surnames in the United States
// from https://en.wikipedia.org/wiki/List_of_most_common_surnames_in_North_America
package namegen

import (
	"fmt"
	"math/rand"
)

type Name struct {
	name      string
	threshold float64
}

func (n Name) String() string {
	return n.name
}

func Male() string {
	firstPos := rand.Intn(len(maleNames))
	surnamePos := rand.Intn(len(surnames))
	return fmt.Sprintf("%s %s", maleNames[firstPos], surnames[surnamePos])
}
