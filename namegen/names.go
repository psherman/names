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

func MaleFirst() string {
	return maleNames[rand.Intn(len(maleNames))].name
}

func FemaleFirst() string {
	return femaleNames[rand.Intn(len(femaleNames))].name
}

func Surname() string {
	return surnames[rand.Intn(len(surnames))].name
}

func MaleFirstWeighted() string {
	value := rand.Float64() * 1000
	pos := sort.Search(len(maleNames),
		func(i int) bool { return maleNames[i].threshold > value })
	return maleNames[pos].name
}

func FemaleFirstWeighted() string {
	value := rand.Float64() * 1000
	pos := sort.Search(len(femaleNames),
		func(i int) bool { return femaleNames[i].threshold > value })
	return femaleNames[pos].name
}

func SurnameWeighted() string {
	value := rand.Float64() * 1000
	pos := sort.Search(len(surnames),
		func(i int) bool { return surnames[i].threshold > value })
	return surnames[pos].name
}

func Male() string {
	return fmt.Sprintf("%s %s", MaleFirst(), Surname())
}

func Female() string {
	return fmt.Sprintf("%s %s", FemaleFirst(), Surname())
}

func Random() string {
	if rand.Float64() < 0.5 {
		return Male()
	} else {
		return Female()
	}
}

func MaleWeighted() string {
	return fmt.Sprintf("%s %s", MaleFirstWeighted(), SurnameWeighted())
}

func FemaleWeighted() string {
	return fmt.Sprintf("%s %s", FemaleFirstWeighted(), SurnameWeighted())
}

func RandomWeighted() string {
	if rand.Float64() < 0.5 {
		return MaleWeighted()
	} else {
		return FemaleWeighted()
	}
}
