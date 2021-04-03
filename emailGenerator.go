package main

import "fmt"

const (
	emailSpecialCharSet = ".-"
	emailNumberSet      = "0123456789"
	emailAllCharSet     = emailSpecialCharSet + emailNumberSet
)

var (
	firstNames = RainbowTable{
		"andrew",
		"tim",
		"bob",
		"joe",
		"kimmy",
		"margaret",
		"zach",
		"susuan",
		"mike",
		"micheal",
		"pykhul",
		"cyber",
		"punk",
		"zoe",
		"bane",
		"brain",
		"bill",
		"ryan",
		"bart",
		"lisa",
		"homer",
		"betty",
		"chris",
		"matt",
		"tiff",
		"tiffany",
		"bree",
		"james",
	}

	lastNames = RainbowTable{
		"simpson",
		"rammstein",
		"moore",
		"cosby",
		"trump",
		"smith",
		"johnson",
		"asshole",
		"garcia",
		"miller",
		"davis",
		"anderson",
		"thomas",
		"haris",
		"robinson",
		"lewis",
		"flores",
		"green",
		"cook",
		"cox",
		"wood",
		"foster",
		"james",
	}

	emailDomains = RainbowTable{
		"gmail.com",
		"hotmail.com",
		"aol.com",
		"yahoo.com",
		"protonmail.com",
		"att.net",
		"myspace.com",
	}
)

func generateEmail() string {
	return fmt.Sprintf(
		"%s%s@%s",
		firstNames.GetRandomElement(),
		lastNames.GetRandomElement(),
		emailDomains.GetRandomElement(),
	)
}
