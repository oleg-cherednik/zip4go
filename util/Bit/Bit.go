package Bit

const (
	Bit0 = 0b00000001
	Bit1 = 0b00000010
	Bit2 = 0b00000100
	Bit3 = 0b00001000
	Bit4 = 0b00010000
	Bit5 = 0b00100000
	Bit6 = 0b01000000
	Bit7 = 0b10000000

	Bit8  = Bit0 << 8
	Bit9  = Bit1 << 8
	Bit10 = Bit2 << 8
	Bit11 = Bit3 << 8
	Bit12 = Bit4 << 8
	Bit13 = Bit5 << 8
	Bit14 = Bit6 << 8
	Bit15 = Bit7 << 8
)

func IsBitSet(val uint, bits uint) bool {
	return (val & bits) == bits
}

func IsBitClear(val uint, bits uint) bool {
	return (val & bits) == 0
}
