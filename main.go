package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	TOKEN_SHIFT_RIGHT = '>'
	TOKEN_SHIFT_LEFT  = '<'
	TOKEN_INCREMENT   = '+'
	TOKEN_DECREMENT   = '-'
	TOKEN_PRINT_CHAR  = '.'
	TOKEN_GET_CHAR    = ','
	TOKEN_START_LOOP  = '['
	TOKEN_END_LOOP    = ']'
)

type BrainFuck struct {
	memory   [30000]byte
	pointer  uint
	jmpStack []int
}

func NewBrainFuck() *BrainFuck {
	return &BrainFuck{
		memory:   [30000]byte{},
		pointer:  0,
		jmpStack: []int{},
	}
}

func (b *BrainFuck) Run(input string) {
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case TOKEN_INCREMENT:
			b.memory[b.pointer]++
		case TOKEN_DECREMENT:
			b.memory[b.pointer]--
		case TOKEN_SHIFT_RIGHT:
			b.pointer++
		case TOKEN_SHIFT_LEFT:
			b.pointer--
		case TOKEN_PRINT_CHAR:
			fmt.Print(string(b.memory[b.pointer]))
		case TOKEN_GET_CHAR:
			fmt.Scanf("%c", &b.memory[b.pointer])
		case TOKEN_START_LOOP:
			if b.memory[b.pointer] == 0 {
				for {
					i++
					if input[i] == TOKEN_END_LOOP {
						break
					}
				}
			} else {
				b.jmpStack = append(b.jmpStack, i)
			}
		case TOKEN_END_LOOP:
			if b.memory[b.pointer] != 0 {
				i = b.jmpStack[len(b.jmpStack)-1] - 1
				b.jmpStack = b.jmpStack[:len(b.jmpStack)-1]
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	NewBrainFuck().Run(text)
}
