package main

import (
	"fmt"
)

const MEM_SIZE = 0x600

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

func (p *Processor) ADC(arg int) {
	if (int(p.A) + arg) > 0xff {
		p.V = true
	}
	p.A += byte(arg)
	p.setZNFlags(p.A)
}
func (p *Processor) AND(arg int) {}
func (p *Processor) ASL(arg int) {}
func (p *Processor) BCC(arg int) {}
func (p *Processor) BCS(arg int) {}
func (p *Processor) BEQ(arg int) {}
func (p *Processor) BIT(arg int) {}
func (p *Processor) BMI(arg int) {}
func (p *Processor) BNE(arg int) {}
func (p *Processor) BPL(arg int) {}
func (p *Processor) BRK(arg int) {}
func (p *Processor) BVC(arg int) {}
func (p *Processor) BVS(arg int) {}
func (p *Processor) CLC(arg int) { p.C = false }
func (p *Processor) CLD(arg int) { p.D = false }
func (p *Processor) CLI(arg int) { p.I = false }
func (p *Processor) CLV(arg int) { p.V = false }
func (p *Processor) CMP(arg int) {}
func (p *Processor) CPX(arg int) {}
func (p *Processor) CPY(arg int) {}
func (p *Processor) DEC(arg int) {
	p.Mem[arg] -= 1
	p.setZNFlags(p.Mem[arg])
}
func (p *Processor) DEX(arg int) {
	p.X -= 1
	p.setZNFlags(p.X)

}
func (p *Processor) DEY(arg int) {
	p.Y -= 1
	p.setZNFlags(p.Y)

}
func (p *Processor) EOR(arg int) {}
func (p *Processor) INC(arg int) {
	p.Mem[arg] += 1
	p.setZNFlags(p.Mem[arg])
}
func (p *Processor) INX(arg int) {
	p.X += 1
	p.setZNFlags(p.X)
}
func (p *Processor) INY(arg int) {
	p.Y += 1
	p.setZNFlags(p.Y)

}
func (p *Processor) JMP(arg int) {}
func (p *Processor) JSR(arg int) {}
func (p *Processor) LDA(arg int) {
	p.A = p.Mem[arg]
}
func (p *Processor) LDX(arg int) {
	p.X = p.Mem[arg]

}
func (p *Processor) LDY(arg int) {
	p.Y = p.Mem[arg]

}
func (p *Processor) LSR(arg int) {}
func (p *Processor) NOP(arg int) { /*nothing!*/ }
func (p *Processor) ORA(arg int) {}
func (p *Processor) PHA(arg int) {}
func (p *Processor) PHP(arg int) {}
func (p *Processor) PLA(arg int) {}
func (p *Processor) PLP(arg int) {}
func (p *Processor) ROL(arg int) {}
func (p *Processor) ROR(arg int) {}
func (p *Processor) RTI(arg int) {}
func (p *Processor) RTS(arg int) {}
func (p *Processor) SBC(arg int) {}
func (p *Processor) SEC(arg int) { p.C = true }
func (p *Processor) SED(arg int) { p.D = true }
func (p *Processor) SEI(arg int) { p.I = true }
func (p *Processor) STA(arg int) {}
func (p *Processor) STX(arg int) {}
func (p *Processor) STY(arg int) {}
func (p *Processor) TAX(arg int) { p.X = p.A }
func (p *Processor) TAY(arg int) { p.Y = p.A }
func (p *Processor) TSX(arg int) { p.X = p.SP }
func (p *Processor) TXA(arg int) { p.A = p.X }
func (p *Processor) TXS(arg int) { p.SP = p.X }
func (p *Processor) TYA(arg int) { p.A = p.Y }

func main() {
	p := NewProcessor()
	fmt.Printf("%v\n", p)
	p.INX(0)
	p.INX(0)
	p.TXA(0)
	p.ADC(0xff)
	fmt.Printf("%v\n", p)
}
