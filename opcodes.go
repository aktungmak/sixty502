package main

// look up opcode based on name and addressing type
// 0xff represents and illegal opcode
var opcodes = map[string][12]byte{
	//op,   Imm,  ZP,   ZPx,  ZPY,  ABS,  ABSx, ABSY, IND,  INDx, INDY, SNGL, BRA
	"ADC": {0x69, 0x65, 0x75, 0xFF, 0x6D, 0x7D, 0x79, 0xFF, 0x61, 0x71, 0xFF, 0xFF},
	"AND": {0x29, 0x25, 0x35, 0xFF, 0x2D, 0x3D, 0x39, 0xFF, 0x21, 0x31, 0xFF, 0xFF},
	"ASL": {0xFF, 0x06, 0x16, 0xFF, 0x0E, 0x1E, 0xFF, 0xFF, 0xFF, 0xFF, 0x0A, 0xFF},
	"BCC": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x9},
	"BCS": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xB},
	"BEQ": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF},
	"BIT": {0xFF, 0x24, 0xFF, 0xFF, 0x2C, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	"BMI": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x3},
	"BNE": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xD},
	"BPL": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x1},
	"BRK": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0xFF},
	"BVC": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x5},
	"BVS": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7},
	"CLC": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x18, 0xFF},
	"CLD": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xD8, 0xFF},
	"CLI": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x58, 0xFF},
	"CLV": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xB8, 0xFF},
	"CMP": {0xC9, 0xC5, 0xD5, 0xFF, 0xCD, 0xDD, 0xD9, 0xFF, 0xC1, 0xD1, 0xFF, 0xFF},
	"CPX": {0xE0, 0xE4, 0xFF, 0xFF, 0xEC, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	"CPY": {0xC0, 0xC4, 0xFF, 0xFF, 0xCC, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	"DEC": {0xFF, 0xC6, 0xD6, 0xFF, 0xCE, 0xDE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	"DEX": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xCA, 0xFF},
	"DEY": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x88, 0xFF},
	"EOR": {0x49, 0x45, 0x55, 0xFF, 0x4D, 0x5D, 0x59, 0xFF, 0x41, 0x51, 0xFF, 0xFF},
	"INC": {0xFF, 0xE6, 0xF6, 0xFF, 0xEE, 0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	"INX": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE8, 0xFF},
	"INY": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xC8, 0xFF},
	"JMP": {0xFF, 0xFF, 0xFF, 0xFF, 0x4C, 0xFF, 0xFF, 0x6C, 0xFF, 0xFF, 0xFF, 0xFF},
	"JSR": {0xFF, 0xFF, 0xFF, 0xFF, 0x20, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	"LDA": {0xA9, 0xA5, 0xB5, 0xFF, 0xAD, 0xBD, 0xB9, 0xFF, 0xA1, 0xB1, 0xFF, 0xFF},
	"LDX": {0xA2, 0xA6, 0xFF, 0xB6, 0xAE, 0xFF, 0xBE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	"LDY": {0xA0, 0xA4, 0xB4, 0xFF, 0xAC, 0xBC, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	"LSR": {0xFF, 0x46, 0x56, 0xFF, 0x4E, 0x5E, 0xFF, 0xFF, 0xFF, 0xFF, 0x4A, 0xFF},
	"NOP": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEA, 0xFF},
	"ORA": {0x09, 0x05, 0x15, 0xFF, 0x0D, 0x1D, 0x19, 0xFF, 0x01, 0x11, 0xFF, 0xFF},
	"PHA": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x48, 0xFF},
	"PHP": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x08, 0xFF},
	"PLA": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x68, 0xFF},
	"PLP": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x28, 0xFF},
	"ROL": {0xFF, 0x26, 0x36, 0xFF, 0x2E, 0x3E, 0xFF, 0xFF, 0xFF, 0xFF, 0x2A, 0xFF},
	"ROR": {0xFF, 0x66, 0x76, 0xFF, 0x6E, 0x7E, 0xFF, 0xFF, 0xFF, 0xFF, 0x6A, 0xFF},
	"RTI": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x40, 0xFF},
	"RTS": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x60, 0xFF},
	"SBC": {0xE9, 0xE5, 0xF5, 0xFF, 0xED, 0xFD, 0xF9, 0xFF, 0xE1, 0xF1, 0xFF, 0xFF},
	"SEC": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x38, 0xFF},
	"SED": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0xFF},
	"SEI": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x78, 0xFF},
	"STA": {0xFF, 0x85, 0x95, 0xFF, 0x8D, 0x9D, 0x99, 0xFF, 0x81, 0x91, 0xFF, 0xFF},
	"STX": {0xFF, 0x86, 0xFF, 0x96, 0x8E, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	"STY": {0xFF, 0x84, 0x94, 0xFF, 0x8C, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	"TAX": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xAA, 0xFF},
	"TAY": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xA8, 0xFF},
	"TSX": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xBA, 0xFF},
	"TXA": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x8A, 0xFF},
	"TXS": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x9A, 0xFF},
	"TYA": {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x98, 0xFF},
}

// addressing mode index into arrays above
const (
	amIMM = iota
	amZP
	amZPX
	amZPY
	amABS
	amABSX
	amABSY
	amIND
	amINDX
	amINDY
	amSNGL
	amBRA
)

func isOpcode(tok string) bool {
	_, ok := opcodes[tok]
	return ok
}
