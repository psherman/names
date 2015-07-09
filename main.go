package main

import (
	"fmt"
	"github.com/psherman/names/namegen"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Printf("Random Names\n-------------------\n")
	fmt.Printf("Male:\t%s\n", namegen.Male())
	fmt.Printf("Female:\t%s\n", namegen.Female())
	fmt.Printf("Random:\t%s\n", namegen.Random())

	fmt.Printf("\nWeighted Names\n-------------------\n")
	fmt.Printf("Male:\t%s\n", namegen.WeightedMale())
	fmt.Printf("Female:\t%s\n", namegen.WeightedFemale())
	fmt.Printf("Random:\t%s\n", namegen.WeightedRandom())
}
