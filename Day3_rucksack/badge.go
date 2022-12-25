package main

import (
	"strings"
)

func FindBadge(group []string) string {
	for _, character := range strings.Split(group[0], "") {
		if strings.Contains(group[1], character) && strings.Contains(group[2], character) {
			return character
		}
	}
	return ""
}

func FindAllBadges(groups [][]string) (badges string) {
	for _, group := range groups {
		badges += FindBadge(group)
	}
	return
}
