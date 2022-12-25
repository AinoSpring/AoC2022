package main

import (
	"log"
	"strconv"
	"strings"
)

type Cleaner struct {
	from int
	to   int
}

func NewCleaner(cleaner string) Cleaner {
	from, err := strconv.Atoi(strings.Split(cleaner, "-")[0])
	if err != nil {
		log.Panic(err)
	}
	to, err := strconv.Atoi(strings.Split(cleaner, "-")[1])
	if err != nil {
		log.Panic(err)
	}
	return Cleaner{from: from, to: to}
}

func (cleaner *Cleaner) Contains(other Cleaner) bool {
	return cleaner.from <= other.from && cleaner.to >= other.to
}

func (cleaner *Cleaner) Between(other Cleaner) bool {
	return Between(other, cleaner.from) || Between(other, cleaner.to)
}
