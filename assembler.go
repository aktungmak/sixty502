package main

import (
	"regexp"
	"log"
	"fmt"
	"strings"
)

const opcodeNames = [...]string{
	"ADC", "AND", "ASL", "BCC", "BCS", "BEQ", "BIT",
	"BMI", "BNE", "BPL", "BRK", "BVC", "BVS", "CLC",
	"CLD", "CLI", "CLV", "CMP", "CPX", "CPY", "DEC",
	"DEX", "DEY", "EOR", "INC", "INX", "INY", "JMP",
	"JSR", "LDA", "LDX", "LDY", "LSR", "NOP", "ORA",
	"PHA", "PHP", "PLA", "PLP", "ROL", "ROR", "RTI",
	"RTS", "SBC", "SEC", "SED", "SEI", "STA", "STX",
	"STY", "TAX", "TAY", "TSX", "TXA", "TXS", "TYA",
}

const (
	absolute    = regexp.MustCompile("\\$[[:xdigit:]]{4}")
	absoluteX   = regexp.MustCompile("\\$[[:xdigit:]]{4},X")
	absoluteY   = regexp.MustCompile("\\$[[:xdigit:]]{4},Y")
	accumulator = regexp.MustCompile("A")
	immediate   = regexp.MustCompile("#\\$[[:xdigit:]]{2}")
	indirect    = regexp.MustCompile("\\(\\$[[:xdigit:]]{4}\\)")
	indirectX   = regexp.MustCompile("\\(\\$[[:xdigit:]]{4},X\\)")
	indirectY   = regexp.MustCompile("\\(\\$[[:xdigit:]]{4}\\),Y")
	zeropage    = regexp.MustCompile("\\$[[:xdigit:]]{2}")
	zeropageX   = regexp.MustCompile("\\$[[:xdigit:]]{2},X")
	zeropageY   = regexp.MustCompile("\\$[[:xdigit:]]{2},Y")
	// relative is calculated based on the branch loc
	// implied mode is implied, so no regex needed 
)

// this keeps track of the assembler state
type Assembler struct {
	StartAddr uint16
	PC        uint16
	Labels    map[string]uint16
	Src       string
	IrLines   []string
}

func (a *Assembler) Assemble() {
	fmt.Print("assembling")
}

// choose how to parse the tokens based on the content
func (a *Assembler) FirstPass(tokens []string) (bytelen int, err error) {
	switch len(tokens) {
	case 0:
		err = "empty line"
	case 1:
		// see if the token is an opcode
		isOpcode := false
		for opcode := range opcodeNames {
			if tokens[0] == opcode {
				isOpcode = true
				break
			}
		}
		if isOpcode {
			// OPCODE
			// replace opcode with correct hex val (implied mode)
			// increment pc
			a.PC++
		} else {
			// LABELDECL (done)
			// add to label map
			a.AddLabel(tokens[0])
		}
	case 2:
		// see if the first token is an opcode
		firstIsOpcode := false
		for opcode := range opcodeNames {
			if tokens[0] == opcode {
				firstIsOpcode = true
				break
			}
		}
		if firstIsOpcode {
			// OPCODE OPERAND
			// parse operand to find addressing type
			if match {

			} else if match2 {

			} 
			// replace opcode with correct hex val
			// increment pc appropriately
		} else {
			// LABEL OPCODE
			addrMode := implied
			// add to label to label map
			a.AddLabel(tokens[0])
			// replace opcode with correct hex val
			// increment pc appropriately
		}

	case 3:
		// LABEL OPCODE OPERAND
		// add to label to label map
		a.AddLabel(tokens[0])
		// work out addressing mode of tokens[2]
		// replace opcode with hex value

	default:
		err = "couldn't parse line"
	}

}

// put the label into the label map, if its not there already
func (a *Assembler) AddLabel(label string) {
	addr, present := a.Labels[tokens[0]]
	if present {
		log.Printf("Warning: label %s already defined at %x", label, addr)
	} else {
		a.Labels[label] = a.PC
	}
}

//split a line into individual tokens
func Tokenize(line string) []string {
	return strings.Split(line, " ")
}

// eliminate all tokens after a comment
func StripComments(tokens []string) (ret []string) {
	for _, tok := range tokens {
		if strings.Contains(tok, ";") {
			//its a comment, stop processing
			break
		}
		ret = append(ret, tok)
	}

	return ret
}

func 
