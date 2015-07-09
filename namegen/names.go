// Generate a random male or female (or either) name
// Uses the top 100 male/female names from 1915-2014
// from http://www.ssa.gov/oact/babynames/decades/century.html
// and the top 100 surnames in the United States
// from https://en.wikipedia.org/wiki/List_of_most_common_surnames_in_North_America
package namegen

import (
	"fmt"
	"math/rand"
	"sort"
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

func Female() string {
	firstPos := rand.Intn(len(femaleNames))
	surnamePos := rand.Intn(len(surnames))
	return fmt.Sprintf("%s %s", femaleNames[firstPos], surnames[surnamePos])
}

func Random() string {
	if rand.Float64() < 0.5 {
		return Male()
	} else {
		return Female()
	}
}

func WeightedMale() string {
	firstValue := rand.Float64() * 1000
	surnameValue := rand.Float64() * 1000
	firstPos := sort.Search(len(maleNames),
		func(i int) bool { return maleNames[i].threshold > firstValue })
	surnamePos := sort.Search(len(surnames),
		func(i int) bool { return surnames[i].threshold > surnameValue })
	return fmt.Sprintf("%s %s", maleNames[firstPos], surnames[surnamePos])
}

func WeightedFemale() string {
	firstValue := rand.Float64() * 1000
	surnameValue := rand.Float64() * 1000
	firstPos := sort.Search(len(femaleNames),
		func(i int) bool { return femaleNames[i].threshold > firstValue })
	surnamePos := sort.Search(len(surnames),
		func(i int) bool { return surnames[i].threshold > surnameValue })
	return fmt.Sprintf("%s %s", femaleNames[firstPos], surnames[surnamePos])
}

func WeightedRandom() string {
	if rand.Float64() < 0.5 {
		return WeightedMale()
	} else {
		return WeightedFemale()
	}
}
