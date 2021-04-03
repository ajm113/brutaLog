package main

import "math/rand"

const (
	passwordLowerCharSet   = "abcdedfghijklmnopqrst"
	passwordUpperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	passwordSpecialCharSet = "!@#$"
	passwordNumberSet      = "0123456789"
	passwordAllCharSet     = passwordLowerCharSet + passwordUpperCharSet + passwordSpecialCharSet + passwordNumberSet
)

var (
	randomWords = RainbowTable{
		"car",
		"battery",
		"horse",
		"train",
		"plane",
		"house",
		"grandma",
		"grandpa",
		"crazy",
		"everyone",
		"zippy",
		"password",
		"geek",
		"squad",
	}
)

// generatePassword Generates completely random passwords for phishing sites.
// Uses mix between completely random to using key words.
func generatePassword() (outputBuffer string) {
	numberOfRandomWords := rand.Intn(5)
	inBetweenRandomWordRandomCharacterLength := rand.Intn(5)
	prefixLength := rand.Intn(5)
	suffixLength := rand.Intn(5)
	minLength := rand.Intn(24)

	for i := 0; i < prefixLength; i++ {
		outputBuffer += string(passwordAllCharSet[rand.Intn(len(passwordAllCharSet))])
	}

	for i := 0; i < numberOfRandomWords; i++ {
		outputBuffer += randomWords.GetRandomElement()

		if (i + 1) >= numberOfRandomWords {
			break
		}

		for j := 0; j < inBetweenRandomWordRandomCharacterLength; j++ {
			outputBuffer += string(passwordAllCharSet[rand.Intn(len(passwordAllCharSet))])
		}
	}

	for i := 0; i < suffixLength; i++ {
		outputBuffer += string(passwordAllCharSet[rand.Intn(len(passwordAllCharSet))])
	}

	if minLength > len(outputBuffer) {
		for i := 0; i < (minLength - len(outputBuffer)); i++ {
			outputBuffer += string(passwordAllCharSet[rand.Intn(len(passwordAllCharSet))])
		}
	}

	return
}
