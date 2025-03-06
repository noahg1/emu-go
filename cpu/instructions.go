package main

import (
	"fmt"
)

// ArithmeticTarget represents the registers
type Target int

const (
	A Target = iota
	B
	C
	D
	E
	H
	L
	BC
	DE
	HL
)

// Instruction represents an instruction with an arithmetic target
type Instruction struct {
	opcode string
	target Target
}

// Constants for instruction types
const (
	ADD   = "ADD"
	ADDHL = "ADDHL"
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
func NewInstruction(opcode string, target Target) Instruction {
	return Instruction{
		opcode: opcode,
		target: target,
	}
}

func (cpu *CPU) execute(instruct Instruction) (string, error) {
	switch instruct.opcode{
	case "ADD" :
		switch instruct.target {
		case A :
			cpu.registers.a = cpu.add(cpu.registers.a)
			return "" , nil
		case B :
			cpu.registers.a = cpu.add(cpu.registers.b)
			return "" , nil
		case C :
			cpu.registers.a = cpu.add(cpu.registers.c)
			return "" , nil
		case D :
			cpu.registers.a = cpu.add(cpu.registers.d)
			return "" , nil
		case E :
			cpu.registers.a = cpu.add(cpu.registers.e)
			return "" , nil
		case H :
			cpu.registers.a = cpu.add(cpu.registers.h)
			return "" , nil
		case L :
			cpu.registers.a = cpu.add(cpu.registers.l)
			return "" , nil
		}
		return "", fmt.Errorf("Unsupported target: %d for instruction: ADD", instruct.target)
	case "ADDHL" :
		switch instruct.target {
		case BC :
			cpu.registers.set_hl(cpu.addhl(cpu.registers.get_bc()))
			return "", nil
		case DE :
			cpu.registers.set_hl(cpu.addhl(cpu.registers.get_de()))
			return "", nil
		case HL :
			cpu.registers.set_hl(cpu.addhl(cpu.registers.get_hl()))
			return "", nil
		}
		return "", fmt.Errorf("Unsupported target: %d for instruction: ADDHL", instruct.target)
	}
	return "", nil
}

// Add performs the addition operation on the A register and another value
func (cpu *CPU) add(value uint8) uint8 {
	newValue := uint16(cpu.registers.a) + uint16(value)

	cpu.flags.carry = newValue > 0xFF
	cpu.flags.half_carry = (cpu.registers.a & 0xF) + (value & 0xF) >= 0x10
	cpu.flags.subtract = false

	return uint8(newValue)
}

func (cpu *CPU) addhl(value uint16) uint16 {
	newValue := uint16(cpu.registers.get_hl()) + uint16(value)

	cpu.flags.carry = newValue > 0xFF
	cpu.flags.half_carry = ((cpu.registers.get_hl() & 0xFFF) + (value & 0xFFF) >= 0x1000)
	cpu.flags.subtract = false

	return uint16(newValue)
}

func main() {
	// Create a new CPU and set some initial values
	cpu := NewCPU()
	cpu.registers.set_hl(0x0000)
	cpu.registers.set_af(0x6400) // 100 in A
	cpu.registers.set_bc(0x00FF) // 255 in C

	// Create an ADD instruction targeting the C register
	instruction := NewInstruction(ADDHL, BC)

	// Execute the instruction
	_, err := cpu.execute(instruction)
	if err != nil {
		panic(err)
	}

	// Print the result of the addition
	fmt.Printf("HL register after ADDHL: %d\n", cpu.registers.get_hl()) // Should be 99
	fmt.Printf("Carry flag: %v\n", cpu.flags.carry) // Should be true
	fmt.Printf("Half Carry Flag: %v\n", cpu.flags.half_carry)
}