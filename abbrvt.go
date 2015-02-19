package abbrvt

import (
	"errors"
	"regexp"
)

var vowel *regexp.Regexp = regexp.MustCompile(`[aeiou]`)
var space *regexp.Regexp = regexp.MustCompile(`\s`)
var capital *regexp.Regexp = regexp.MustCompile(`\b[A-Z]`)

/*
   Determines if the character string is a vowel
   Expected to pass only single characters here
*/
func isVowel(character string) bool {
	return vowel.MatchString(character)
}

/*
   Determines if the string contains multiple words
*/
func isMultiWord(input string) bool {
	return space.MatchString(input)
}

/*
   Determines the locations of the character which are capital letters
   and occur at the beginning of a word
*/
func getCapitals(input string) [][]int {
	return capital.FindAllStringIndex(input, -1)
}

/*
   Public function that abbreviates the given string
*/
func Get(original string) (string, error) {
	results := ""

	//First check if the string contains multiple words
	if isMultiWord(original) {
		//Now check if those words start with capital letters
		capLocations := getCapitals(original)

		if len(capLocations) < 3 {
			return "", errors.New("String contains too few capital words. Can't abbreviate.")
		}

		for _, location := range capLocations {
			results += string(original[location[0]])
		}
	} else {
		for i := 0; i < len(original); i++ {
			if !isVowel(string(original[i])) {
				results += string(original[i])
			}
		}
	}

	return results, nil
}
