package main

type Registers struct {
	a, b, c, d, e, f, h, l uint8
}

func (r *Registers) get_bc() uint16 {
	// Cast B to a 16 bit int and shift bits left 8, the OR it with C 
	return (uint16(r.b) << 8) | uint16(r.c) // This does a bitwise combination register C and B to a single 16 bit uint
}

func (r *Registers) set_bc(value uint16) {
	r.b = uint8((value & 0xFF00) >> 8) // Extract the high byte and store in r.b
	r.c = uint8(value & 0xFF) // Extract the low byte and store in r.c
}

func (r *Registers) get_af() uint16 {
	return (uint16(r.a) << 8) | uint16(r.f) 
}

func (r *Registers) set_af(value uint16) {
	r.a = uint8((value & 0xFF00) >> 8)
	r.f = uint8(value & 0xFF)
}

func (r *Registers) get_de() uint16 {
	return (uint16(r.d) << 8) | uint16(r.e) 
}

func (r *Registers) set_de(value uint16) {
	r.d = uint8((value & 0xFF00) >> 8)
	r.e = uint8(value & 0xFF)
}

func (r *Registers) get_hl() uint16 {
	return (uint16(r.h) << 8) | uint16(r.l) 
}

func (r *Registers) set_hl(value uint16) {
	r.h = uint8((value & 0xFF00) >> 8)
	r.l = uint8(value & 0xFF)
}

