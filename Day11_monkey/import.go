package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	operationWrapper = map[string]func(int) func(uint64) uint64{
		"+": func(x int) func(uint64) uint64 {
			if x == -1 {
				return func(a uint64) uint64 {
					return a + a
				}
			}
			return func(a uint64) uint64 {
				return a + uint64(x)
			}
		},
		"*": func(x int) func(uint64) uint64 {
			if x == -1 {
				return func(a uint64) uint64 {
					return a * a
				}
			}
			return func(a uint64) uint64 {
				return a * uint64(x)
			}
		},
	}
)

func Import() (monkeys Monkeys) {
	monkeys = NewMonkeys()
	contentBytes, err := os.ReadFile(input)
	Handle(err)
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	for _, line := range strings.Split(content, "\n") {
		line = strings.Replace(line, "  ", "", -1)
		var command = strings.Split(line, " ")
		switch strings.ToLower(command[0]) {
		case "monkey":
			monkeys.Append(strings.Replace(command[1], ":", "", 1), NewMonkey())
		case "starting":
			for _, item := range command[2:] {
				item = strings.Replace(item, ",", "", -1)
				itemInt, err := strconv.Atoi(item)
				Handle(err)
				monkeys.Last().items = append(monkeys.Last().items, uint64(itemInt))
			}
		case "operation:":
			var factor = -1
			if command[5] != "old" {
				factor, err = strconv.Atoi(command[5])
				Handle(err)
			}
			monkeys.Last().operation = operationWrapper[command[4]](factor)
		case "test:":
			divisor, err := strconv.Atoi(command[3])
			Handle(err)
			monkeys.lowestCommonDivider *= uint64(divisor)
			monkeys.Last().test = uint64(divisor)
		case "if":
			monkeys.Last().throw[command[1] == "true:"] = command[5]
		case "":
			continue
		default:
			log.Panicf("Invalid instruction: %v", strings.Join(command, " "))
		}
	}
	return
}
