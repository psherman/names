package main

import (
	"fmt"
	"github.com/psherman/names/namegen"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Println(namegen.Male())
	fmt.Println(namegen.Female())
	fmt.Println(namegen.Random())
}
