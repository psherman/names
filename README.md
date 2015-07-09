##namegen

Generate a random male/female name using most popular names in the US.

    package main

    import (
        "fmt"
        "github.com/psherman/names/namegen"
    )

    func main() {
        male := namegen.Male()
        female := namegen.Female()

        fmt.Printf("%s and %s, the dynamic duo!\n", male, female)
    }

#####`namegen.MaleFirst`

Return a random male first name.

#####`namegen.FemaleFirst`

Return a random female first name.

#####`namegen.Surname`

Return a random surname.

#####`namegen.WeightedMaleFirst`

Return a random male first name. The names are weighted out of 1000, with more popular names having a higher chance of being selected.

#####`namegen.WeightedFemaleFirst`

Return a random female first name. The names are weighted out of 1000, with more popular names having a higher chance of being selected.

#####`namegen.WeightedSurname`

Return a random surname. The names are weighted out of 1000, with more popular names having a higher chance of being selected.

#####`namegen.Male`

Return a male name using a male name randomly selected from the male name array and a surname randomly selected from the surname array

#####`namegen.Female`

Return a female name using a male name randomly selected from the female name array and a surname randomly selected from the surname array

#####`namegen.Random`

50/50 chance to return either a male or female name using either `namegen.Male` or `namegen.Female`

#####`namegen.WeightedMale`

Return a male name using a male name randomly selected from the male name array and a surname randomly selected from the surname array. The names are weighted out of 1000, with more popular names having a higher chance of being selected.

#####`namegen.WeightedFemale`

Return a female name using a male name randomly selected from the female name array and a surname randomly selected from the surname array. The names are weighted out of 1000, with more popular names having a higher chance of being selected.

#####`namegen.WeightedRandom`

50/50 chance to return either a male or female name using either `namegen.WeightedMale` or `namegen.WeightedFemale`