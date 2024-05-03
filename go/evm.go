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

// Run runs the EVM code and returns the stack and a success indicator.
func Evm(code []byte) ([]*big.Int, bool) {
	var stack []*big.Int
	pc := 0
	// print(len(code))
	for pc < len(code) {
		op := code[pc]
		pc++

		switch {
		case op == 0x00:
		case op == 0x5f:
			stack = append([]*big.Int{big.NewInt(0x00)}, stack...)
		case op > 95 && op < 128:
			size := int(op) - 95
			operand := big.NewInt(0)
			for i := 0; i < size; i++ {
				/* concatenate the bytes of the operand */
				operand.Lsh(operand, 8)
				operand.Or(operand, big.NewInt(int64(code[pc])))
				pc++
			}
			stack = append([]*big.Int{operand}, stack...)
		default:
		}
	}

	return stack, true
}
