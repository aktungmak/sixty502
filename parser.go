package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// represents a single operation or label
type Command struct {
	cmdType int
	tokens  []string
	label   string
	opcode  byte
	operand int
}

// identifiers for the 5 types of command syntax
const (
	opcode = iota
	label
	opcodeOperand
	labelOpcode
	labelOpcodeOperand
)

//split a line into individual tokens
func Parse(in io.Reader) {
	s := bufio.NewScanner(in)

	lnum := 0
	for s.Scan() {
		l := s.Text()
		lnum++
		cmd, err := parseLine(l, lnum)
		log.Printf("%v", cmd)
		if err != nil {
			log.Print(err)
			break
		}
	}
	if err := s.Err(); err != nil {
		log.Print(err)
	}
}

func parseLine(line string, lnum int) (Command, error) {
	var cmd Command
	var err error
	// strip comment
	line = strings.Split(line, ";")[0]
	// upcase
	line = strings.ToUpper(line)
	// split into tokens
	tokens := strings.Fields(line)
	// identify type
	switch len(tokens) {
	case 0:
		// empty line, skip
	case 1:
		// see if the first token is an opcode
		if isOpcode(tokens[0]) {
			// OPCODE
			cmd, err = parseOpcode(tokens)

		} else {
			// LABEL
			cmd, err = parseLabel(tokens)
		}
	case 2:
		// see if the first token is an opcode
		if isOpcode(tokens[0]) {
			// OPCODE OPERAND
			cmd, err = parseOpcodeOperand(tokens)
		} else {
			// LABEL OPCODE
			cmd, err = parseLabelOpcode(tokens)
		}
	case 3:
		// LABEL OPCODE OPERAND
		cmd, err = parseLabelOpcodeOperand(tokens)
	default:
		err = errors.New("unknown syntax")
	}
	return cmd, err
}

func parseOpcode(tokens []string) (Command, error) {
	cmd := Command{
		cmdType: opcode,
		tokens:  tokens,
		opcode:  opcodes[tokens[0]][amSNGL],
	}
	return cmd, nil
}

func parseLabel(tokens []string) (Command, error) {
	cmd := Command{
		cmdType: label,
		tokens:  tokens,
		label:   tokens[0],
	}
	return cmd, nil
}

func parseOpcodeOperand(tokens []string) (Command, error) {
	am, val, err := checkAddrMode(tokens[1])
	if err != nil {
		return Command{}, err
	}

	cmd := Command{
		cmdType: opcodeOperand,
		tokens:  tokens,
		opcode:  opcodes[tokens[0]][am],
		operand: val,
	}
	return cmd, nil
}

func parseLabelOpcode(tokens []string) (Command, error) {
	cmd := Command{
		cmdType: labelOpcode,
		tokens:  tokens,
		label:   tokens[0],
		opcode:  opcodes[tokens[1]][amSNGL],
	}
	return cmd, nil
}

func parseLabelOpcodeOperand(tokens []string) (Command, error) {
	return Command{}, nil
}

var absolute = regexp.MustCompile("^\\$[[:xdigit:]]{4}$")
var absoluteX = regexp.MustCompile("^\\$[[:xdigit:]]{4},X$")
var absoluteY = regexp.MustCompile("^\\$[[:xdigit:]]{4},Y$")
var accumulator = regexp.MustCompile("^A$")
var immediate = regexp.MustCompile("^#\\$[[:xdigit:]]{2}$")
var indirect = regexp.MustCompile("^\\(\\$[[:xdigit:]]{4}\\)$")
var indirectX = regexp.MustCompile("^\\(\\$[[:xdigit:]]{4},X\\)$")
var indirectY = regexp.MustCompile("^\\(\\$[[:xdigit:]]{4}\\),Y$")
var zeropage = regexp.MustCompile("^\\$[[:xdigit:]]{2}$")
var zeropageX = regexp.MustCompile("^\\$[[:xdigit:]]{2},X$")
var zeropageY = regexp.MustCompile("^\\$[[:xdigit:]]{2},Y$")

func checkAddrMode(tok string) (addrMode int, val int, err error) {
	switch {
	case absolute.MatchString(tok):
		addrMode = amABS
	case absoluteX.MatchString(tok):
		addrMode = amABSX
	case absoluteY.MatchString(tok):
		addrMode = amABSY
	case accumulator.MatchString(tok):
		addrMode = amSNGL
	case immediate.MatchString(tok):
		addrMode = amIMM
	case indirect.MatchString(tok):
		addrMode = amIND
	case indirectX.MatchString(tok):
		addrMode = amINDX
	case indirectY.MatchString(tok):
		addrMode = amINDY
	case zeropage.MatchString(tok):
		addrMode = amZP
	case zeropageX.MatchString(tok):
		addrMode = amZPX
	case zeropageY.MatchString(tok):
		addrMode = amZPY
	default:
		// no match, error!
		addrMode = -1
		err = errors.New("unknown addressing mode")
	}
	// if we could identify the mode, extract the value
	if addrMode >= 0 {
		r := regexp.MustCompile("\\$([[:xdigit:]]{2,4})")
		sv := r.FindStringSubmatch(tok)[1]
		var iv int64
		iv, err = strconv.ParseInt(sv, 16, 0)
		val = int(iv)
	}
	return
}
