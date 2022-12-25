package main

import "strings"

type INSTRUCTION int

const (
	NOOP INSTRUCTION = iota
	ADDX
)

type CPU struct {
	xHist []int
	xReg  int
	crt   CRT
}

func NewCPU() (cpu CPU) {
	cpu = CPU{make([]int, 0), 1, CRT{}}
	cpu.Cycle()
	return
}

func (cpu *CPU) Cycle() {
	cpu.xHist = append(cpu.xHist, cpu.xReg)
	cpu.crt.Cycle(len(cpu.xHist), cpu.xReg)
}

func (cpu *CPU) GetCycle(cycle int) int {
	return cpu.xHist[cycle-1]
}

func (cpu *CPU) Instruction(instruction INSTRUCTION, arg int) {
	switch instruction {
	case NOOP:
		cpu.Cycle()
	case ADDX:
		cpu.Cycle()
		cpu.xReg += arg
		cpu.Cycle()
	}
}

func (cpu *CPU) XRegSum(timestamps []int) (sum int) {
	for _, timestamp := range timestamps {
		sum += timestamp * cpu.GetCycle(timestamp)
	}
	return
}

func ParseInstruction(instruction string) INSTRUCTION {
	switch strings.ToLower(instruction) {
	case "noop":
		return NOOP
	case "addx":
		return ADDX
	}
	return -1
}
