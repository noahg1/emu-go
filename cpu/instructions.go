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
	ADC   = "ADC"
	SUB   = "SUB"
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

func (cpu *CPU) execute(instruct Instruction) error {
	registers := map[Target]*uint8{
		A: &cpu.registers.a,
		B: &cpu.registers.b,
		C: &cpu.registers.c,
		D: &cpu.registers.d,
		E: &cpu.registers.e,
		H: &cpu.registers.h,
		L: &cpu.registers.l,
	}

	operation := map[string]func(uint8) uint8{
		ADD:  cpu.add,
		ADC:  cpu.adc,
		SUB:  cpu.sub,
	}

	if op, ok := operation[instruct.opcode]; ok {
		target := *registers[instruct.target]
		cpu.registers.a = op(target)
		return nil
	}
	switch instruct.opcode{
	case "ADDHL" :
		switch instruct.target {
		case BC :
			cpu.addhl(cpu.registers.get_bc())
			return nil
		case DE :
			cpu.addhl(cpu.registers.get_de())
			return nil
		case HL :
			cpu.addhl(cpu.registers.get_hl())
			return nil
		default :
			return fmt.Errorf("unsupported target: %d for instruction: addhl", instruct.target)
		}
	}

	return fmt.Errorf("unsupported opcode: %s", instruct.opcode)
}

// Add performs the addition operation on the A register and another value
func (cpu *CPU) add(value uint8) uint8 {
	a := cpu.registers.a
	newValue := uint16(a) + uint16(value)

	cpu.flags.carry = newValue > 0xFF
	cpu.flags.half_carry = (a & 0xF) + (value & 0xF) >= 0x10
	cpu.flags.subtract = false
	cpu.flags.zero = newValue == 0

	return uint8(newValue)
}

func (cpu *CPU) addhl(value uint16) {
	hl := cpu.registers.get_hl()
	newValue := uint32(hl) + uint32(value)

	cpu.flags.carry = newValue > 0xFFFF
	cpu.flags.half_carry = ((hl & 0xFF) + (value & 0xFF) >= 0x100)
	cpu.flags.subtract = false
	cpu.flags.zero = newValue == 0

	cpu.registers.set_hl(uint16(newValue))
}

func (cpu *CPU) adc(value uint8) uint8 {
	var carryValue uint8
	if cpu.flags.carry {
		carryValue = 1
	} else {
		carryValue = 0
	}
	a := cpu.registers.a

	newValue := uint16(a) + uint16(value) + uint16(carryValue)

	cpu.flags.carry = uint16(a) + uint16(value) + uint16(carryValue) > 0xFF
	cpu.flags.half_carry = (a & 0x0F) + (value & 0x0F) + uint8(carryValue) > 0xF
	cpu.flags.subtract = false
	cpu.flags.zero = newValue == 0

	return uint8(newValue)
}

func (cpu *CPU) sub(value uint8) uint8 {
	a := cpu.registers.a
	newValue := uint16(a) - uint16(value)

	cpu.flags.carry = int16(a) - int16(value) < 0
	cpu.flags.half_carry = (int(a) & 0xF) - (int(value) & 0xF) < 0
	cpu.flags.subtract = true
	cpu.flags.zero = newValue == 0

	return uint8(newValue)
}

func testADD() {
	// Create a new CPU and set some initial values
	cpu := NewCPU()
	cpu.registers.set_af(0x6400) // 100 in A
	cpu.registers.set_bc(0x00FF) // 255 in C

	// Create an ADD instruction targeting the C register
	instruction := NewInstruction(ADD, C)

	// Execute the instruction
	err := cpu.execute(instruction)
	if err != nil {
		panic(err)
	}

	// Print the result of the addition
	fmt.Printf("\nA register after ADD: %d\n", cpu.registers.a) // Should be 99
	fmt.Printf("Carry flag: %v\n", cpu.flags.carry) // Should be true
	fmt.Printf("Half Carry Flag: %v\n\n", cpu.flags.half_carry) // Should be true
}

func testADDHL() {
	// Create a new CPU and set some initial values
	cpu := NewCPU()
	cpu.registers.set_hl(0x00FF) // 255 in HL
	cpu.registers.set_bc(0x00FF) // 255 in C

	// Create an ADD instruction targeting the C register
	instruction := NewInstruction(ADDHL, BC)

	// Execute the instruction
	err := cpu.execute(instruction)
	if err != nil {
		panic(err)
	}

	// Print the result of the addition
	fmt.Printf("HL register after ADDHL: %d\n", cpu.registers.get_hl()) // Should be 510
	fmt.Printf("Carry flag: %v\n", cpu.flags.carry) // Should be false
	fmt.Printf("Half Carry Flag: %v\n\n", cpu.flags.half_carry) // Should be true
}

func testADC() {
	// Create a new CPU and set some initial values
	cpu := NewCPU()
	cpu.registers.a = 0xFF // 255 in A
	cpu.registers.b = 0xE1 // 225 in B

	// Create an ADD instruction targeting the B register
	instruction := NewInstruction(ADC, B)

	// Execute the instruction
	err := cpu.execute(instruction)
	if err != nil {
		panic(err)
	}

	// Print the result of the addition
	fmt.Printf("A register after ADC: %d\n", cpu.registers.a) // Should be 224
	fmt.Printf("Carry flag: %v\n", cpu.flags.carry) // Should be true
	fmt.Printf("Half Carry Flag: %v\n\n", cpu.flags.half_carry) // Should be true
}

func testSUB() {
	cpu := NewCPU()
	cpu.registers.a = 0xE1 // 225 in A
	cpu.registers.b = 0xFD // 253 in C

	// Create an ADD instruction targeting the C register
	instruction := NewInstruction(SUB, B)

	// Execute the instruction
	err := cpu.execute(instruction)
	if err != nil {
		panic(err)
	}

	// Print the result of the addition
	fmt.Printf("A register after SUB: %d\n", cpu.registers.a) // Should be 228
	fmt.Printf("Carry flag: %v\n", cpu.flags.carry) // Should be true
	fmt.Printf("Half Carry Flag: %v\n\n", cpu.flags.half_carry) // Should be true
}

func main() {
	testADD()
	testADDHL()
	testADC()
	testSUB()
}