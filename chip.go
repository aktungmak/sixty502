// +build ignore

package main

import (
	"fmt"
)

// cpu has 64K of memory
const MEM_SIZE = 0xFFFF

// the SP register is a 8-bit offset from here
// (stack grows downwards)
const STACK_BASE = 0x0100

type Memory [MEM_SIZE]byte

type Processor struct {
	A   byte    //accumulator
	X   byte    //index x
	Y   byte    //index y
	N   bool    //negative
	V   bool    //overflow
	B   bool    //break
	D   bool    //decimal
	I   bool    //interrupt disable
	Z   bool    //zero
	C   bool    //carry
	PC  uint16  //programme counter
	SP  byte    //stack pointer
	Mem *Memory //pointer to memory block
}

func NewProcessor() *Processor {
	fmt.Print("making new processor")
	return &Processor{
		PC:  MEM_SIZE,
		SP:  0xFF,
		Mem: &Memory{},
	}

}

func (p *Processor) setZNFlags(val byte) {
	if val == 0 {
		p.Z = true
	} else {
		p.Z = false
	}

	if (val & 0x80) > 0 {
		p.N = true
	} else {
		p.N = false
	}
}

func (p *Processor) setCarryFromBit0(val byte) {
	if (val & 1) == 0 {
		p.C = false
	} else {
		p.C = true
	}
}
func (p *Processor) setCarryFromBit7(val byte) {
	if ((val >> 7) & 1) == 0 {
		p.C = false
	} else {
		p.C = true
	}
}

// base functions --------------------------------- //
// these implement the actual functionality
// they assume a 16-bit address input
// (which may be ignored)
// they get called by the wrappers
// adr is a int16 address, val is a byte value and
// some functions have no arguments

func (p *Processor) ADC(val byte) {
	if (int(p.A) + int(val)) > 0xff {
		p.V = true
	}
	p.A += val
	p.setZNFlags(p.A)
}
func (p *Processor) AND(adr int16) {}
func (p *Processor) ASL(adr int16) {
	p.setCarryFromBit7(p.Mem[adr])
	p.Mem[adr] = p.Mem[adr] << 1
	p.setZNFlags(p.Mem[adr])
}
func (p *Processor) BCC(adr int16) {}
func (p *Processor) BCS(adr int16) {}
func (p *Processor) BEQ(adr int16) {}
func (p *Processor) BIT(adr int16) {}
func (p *Processor) BMI(adr int16) {}
func (p *Processor) BNE(adr int16) {}
func (p *Processor) BPL(adr int16) {}
func (p *Processor) BRK(adr int16) {}
func (p *Processor) BVC(adr int16) {}
func (p *Processor) BVS(adr int16) {}
func (p *Processor) CLC()          { p.C = false }
func (p *Processor) CLD()          { p.D = false }
func (p *Processor) CLI()          { p.I = false }
func (p *Processor) CLV()          { p.V = false }
func (p *Processor) CMP(adr int16) {}
func (p *Processor) CPX(adr int16) {}
func (p *Processor) CPY(adr int16) {}
func (p *Processor) DEC(adr int16) {
	p.Mem[adr] -= 1
	p.setZNFlags(p.Mem[adr])
}
func (p *Processor) DEX() {
	p.X -= 1
	p.setZNFlags(p.X)
}
func (p *Processor) DEY() {
	p.Y -= 1
	p.setZNFlags(p.Y)
}
func (p *Processor) EOR(adr int16) {}
func (p *Processor) INC(adr int16) {
	p.Mem[adr] += 1
	p.setZNFlags(p.Mem[adr])
}
func (p *Processor) INX() {
	p.X += 1
	p.setZNFlags(p.X)
}
func (p *Processor) INY() {
	p.Y += 1
	p.setZNFlags(p.Y)
}
func (p *Processor) JMP(adr int16) {}
func (p *Processor) JSR(adr int16) {}
func (p *Processor) LDA(adr int16) {
	p.A = p.Mem[adr]
	p.setZNFlags(p.A)
}
func (p *Processor) LDX(adr int16) {
	p.X = p.Mem[adr]
	p.setZNFlags(p.X)
}
func (p *Processor) LDY(adr int16) {
	p.Y = p.Mem[adr]
	p.setZNFlags(p.Y)
}
func (p *Processor) LSR(adr int16) {
	p.setCarryFromBit0(p.Mem[adr])
	p.Mem[adr] = p.Mem[adr] >> 1
	p.setZNFlags(p.Mem[adr])
}
func (p *Processor) NOP()          { /*nothing!*/ }
func (p *Processor) ORA(adr int16) {}
func (p *Processor) PHA(adr int16) {}
func (p *Processor) PHP(adr int16) {}
func (p *Processor) PLA(adr int16) {}
func (p *Processor) PLP(adr int16) {}
func (p *Processor) ROL(adr int16) {}
func (p *Processor) ROR(adr int16) {}
func (p *Processor) RTI(adr int16) {}
func (p *Processor) RTS(adr int16) {}
func (p *Processor) SBC(adr int16) {}
func (p *Processor) SEC()          { p.C = true }
func (p *Processor) SED()          { p.D = true }
func (p *Processor) SEI()          { p.I = true }
func (p *Processor) STA(adr int16) { p.Mem[adr] = p.A }
func (p *Processor) STX(adr int16) { p.Mem[adr] = p.X }
func (p *Processor) STY(adr int16) { p.Mem[adr] = p.Y }
func (p *Processor) TAX()          { p.X = p.A }
func (p *Processor) TAY()          { p.Y = p.A }
func (p *Processor) TSX()          { p.X = p.SP }
func (p *Processor) TXA()          { p.A = p.X }
func (p *Processor) TXS()          { p.SP = p.X }
func (p *Processor) TYA()          { p.A = p.Y }

// end of base functions -------------------------- //

// opcode map ------------------------------------- //
// this map relates a byte value to a specific
// addressing mode and opcode. note not all byte
// values are used!

var OPCODE_MAP = map[byte]func(int16){}
