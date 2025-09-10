package main

import (
	"fmt"
)

// packets are 1–4 bytes long and the first few bits define the packet size (like UTF-8 encoding). Let’s assume rules like these (example only, we can adjust to your spec):

// 1-byte packet → starts with 0xxxxxxx
//
// 2-byte packet → starts with 110xxxxx 10xxxxxx
//
// 3-byte packet → starts with 1110xxxx 10xxxxxx 10xxxxxx
//
// 4-byte packet → starts with 11110xxx 10xxxxxx 10xxxxxx 10xxxxxxpackets are 1–4 bytes long and the first few bits define the packet size (like UTF-8 encoding). Let’s assume rules like these (example only, we can adjust to your spec):
func validatePacketStructure(data []int) bool {
	firstElement := data[0]

	a := true
	i := 0
	for a {
		da := (1 << (7 - i)) & firstElement

		if da == 0 {
			a = false
		} else {
			i += 1
		}
	}

	// if the number of 1's greater than 4 then it is invalid by default'
	if i > 4 {
		return false
	}

	// assume isValid is true
	isValid := true

	for _, d := range data[1:] {
		firstByteOne := (1 << 7) & d
		if firstByteOne > 0 {
			i -= 1
		}
	}

	if i != 1 {
		isValid = false
	}

	return isValid
}

func Shift() {
	// Example - 1
	data := []int{193, 129, 3}
	fmt.Println(validatePacketStructure(data))

	// Example - 2
	data = []int{230, 129, 7}
	fmt.Println(validatePacketStructure(data))

	// Example - 3
	data = []int{248, 129, 129, 129, 129}
	fmt.Println(validatePacketStructure(data))
}
