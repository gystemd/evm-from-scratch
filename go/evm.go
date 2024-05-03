// Package evm is an **incomplete** implementation of the Ethereum Virtual
// Machine for the "EVM From Scratch" course:
// https://github.com/w1nt3r-eth/evm-from-scratch
//
// To work on EVM From Scratch In Go:
//
// - Install Golang: https://golang.org/doc/install
// - Go to the `go` directory: `cd go`
// - Edit `evm.go` (this file!), see TODO below
// - Run `go test ./...` to run the tests
package evm

import (
	"math/big"
)

const STOP = 0
const ADD = 1
const MUL = 2
const SUB = 3
const DIV = 4
const POP = 80
const PUSH0 = 95

var bitMask256 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))

// Run runs the EVM code and returns the stack and a success indicator.
func Evm(code []byte) ([]*big.Int, bool) {
	var stack []*big.Int
	pc := 0
	// print(len(code))
	for pc < len(code) {
		op := code[pc]
		pc++

		switch {
		case op == STOP:
			pc = len(code)
		case op == ADD:
			sum := new(big.Int).Add(stack[0], stack[1])
			sum.And(sum, bitMask256)
			stack = append([]*big.Int{sum}, stack[2:]...)
		case op == MUL:
			product := new(big.Int).Mul(stack[0], stack[1])
			product.And(product, bitMask256)
			stack = append([]*big.Int{product}, stack[2:]...)
		case op == SUB:
			diff := new(big.Int).Sub(stack[0], stack[1])
			diff.And(diff, bitMask256)
			stack = append([]*big.Int{diff}, stack[2:]...)
		case op == DIV:
			var quotient *big.Int
			if stack[1].Cmp(big.NewInt(0)) != 0 {
				quotient = new(big.Int).Div(stack[0], stack[1])
			} else {
				quotient = big.NewInt(0)
			}
			quotient.And(quotient, bitMask256)
			stack = append([]*big.Int{quotient}, stack[2:]...)
		case op == PUSH0:
			stack = append([]*big.Int{big.NewInt(0x00)}, stack...)
		case op > PUSH0 && op < 128:
			size := int(op) - 95
			operand := big.NewInt(0)
			for i := 0; i < size; i++ {
				/* concatenate the bytes of the operand */
				operand.Lsh(operand, 8)
				operand.Or(operand, big.NewInt(int64(code[pc])))
				pc++
			}
			stack = append([]*big.Int{operand}, stack...)
		case op == POP:
			stack = stack[1:]
		default:
		}
	}

	return stack, true
}
