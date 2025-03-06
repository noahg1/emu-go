package main

type FlagRegister struct {
	zero, subtract, half_carry, carry bool
}

const ZERO_FLAG_BYTE_POSITION uint8       = 7
const SUBTRACT_FLAG_BYTE_POSITION uint8   = 6;
const HALF_CARRY_FLAG_BYTE_POSITION uint8 = 5;
const CARRY_FLAG_BYTE_POSITION uint8      = 4;

func (fr *FlagRegister) toRegister() uint8 {
	var res uint8
	
	if fr.zero {
		res |= (1 << ZERO_FLAG_BYTE_POSITION) 
	} else { 
		res |= (0 << ZERO_FLAG_BYTE_POSITION) 
	}

	if fr.subtract {
		res |= (1 << SUBTRACT_FLAG_BYTE_POSITION)
	} else { 
		res |= (0 << SUBTRACT_FLAG_BYTE_POSITION) 
	}

	if fr.half_carry {
		res |= (1 << HALF_CARRY_FLAG_BYTE_POSITION)
	} else { 
		res |= (0 << HALF_CARRY_FLAG_BYTE_POSITION) 
	}

	if fr.carry {
		res |= (1 << CARRY_FLAG_BYTE_POSITION)
	} else { 
		res |= (0 << CARRY_FLAG_BYTE_POSITION) 
	}

	return res
}

func toFlagRegister(register uint8) FlagRegister {
	var zero bool = ((register >> ZERO_FLAG_BYTE_POSITION) & 0b1) != 0;
	var subtract bool = ((register >> SUBTRACT_FLAG_BYTE_POSITION) & 0b1) != 0;
	var half_carry bool = ((register >> HALF_CARRY_FLAG_BYTE_POSITION) & 0b1) != 0;
	var carry bool = ((register >> CARRY_FLAG_BYTE_POSITION) & 0b1) != 0;

	return FlagRegister{
		zero, 
		subtract, 
		half_carry, 
		carry,
	}
}
