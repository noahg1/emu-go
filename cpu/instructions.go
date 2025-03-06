package main

import (
	"fmt"
)

// ArithmeticTarget represents the registers
type ArithmeticTarget int

const (
	A ArithmeticTarget = iota
	B
	C
	D
	E
	H
	L
)

// Instruction represents an instruction with an arithmetic target
type Instruction struct {
	opcode       string
	arithmeticTarget ArithmeticTarget
}

// Constants for instruction types
const (
	ADD = "ADD"
)

// CPU represents the CPU with a set of registers
type CPU struct {
	registers Registers
	flags     FlagRegister
}

// NewCPU creates a new CPU instance
func NewCPU() *CPU {
	registers := emptyRegisters()
	flags := toFlagRegister(0x0000)
	return &CPU{ registers, flags}
}

// NewInstruction creates a new instruction with the given opcode and target
func NewInstruction(opcode string, target ArithmeticTarget) Instruction {
	return Instruction{
		opcode:          opcode,
		arithmeticTarget: target,
	}
}

func (cpu *CPU) execute(instruct Instruction) {
	switch instruct.opcode{
	case "ADD" :
		switch instruct.arithmeticTarget {
		case A :
			value := cpu.registers.a
			newVal := cpu.add(value)
			cpu.registers.a = newVal
		case B :
			value := cpu.registers.b
			newVal := cpu.add(value)
			cpu.registers.a = newVal
		case C :
			value := cpu.registers.c
			newVal := cpu.add(value)
			cpu.registers.a = newVal
		case D :
			value := cpu.registers.d
			newVal := cpu.add(value)
			cpu.registers.a = newVal
		case E :
			value := cpu.registers.e
			newVal := cpu.add(value)
			cpu.registers.a = newVal
		case H :
			value := cpu.registers.h
			newVal := cpu.add(value)
			cpu.registers.a = newVal
		case L :
			value := cpu.registers.l
			newVal := cpu.add(value)
			cpu.registers.a = newVal
		}
	}
}

// Add performs the addition operation on the A register and another value
func (cpu *CPU) add(value uint8) uint8 {
	newValue := uint16(cpu.registers.a) + uint16(value)
	overflow := newValue > 0xFF // Check for overflow
	if overflow {
		// Set the carry flag if there is an overflow
		cpu.flags.carry = true
	} else {
		cpu.flags.carry = false
	}

	// Since Go doesn't have overflowing_add like Rust, we manually return the lower byte
	return uint8(newValue)
}

func main() {
	// Create a new CPU and set some initial values
	cpu := NewCPU()
	cpu.registers.set_af(0x6400) // 100 in A
	cpu.registers.set_bc(0x00FF) // 255 in C

	// Create an ADD instruction targeting the C register
	instruction := NewInstruction(ADD, C)

	// Execute the instruction
	cpu.execute(instruction)

	// Print the result of the addition
	fmt.Printf("A register after ADD: %d\n", cpu.registers.a) // Should be 99
	fmt.Printf("Carry flag: %v\n", cpu.flags.carry) // Should be true
}