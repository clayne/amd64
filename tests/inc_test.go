package tests

import . "github.com/modern-go/amd64"

func init() {
	testCases = append(testCases, []testCase{{
		input: input{
			INC, EAX,
		},
		comment: "rex prefix not required because eax is 32 bit",
		output: []uint8{
			0xff, 0xc0,
		},
	}, {
		input: input{
			INC, RAX,
		},
		comment: "rax is 64 bit, requires rex prefix",
		output: []uint8{
			0x48, 0xff, 0xc0,
		},
	}, {
		input: input{
			INC, AL,
		},
		comment: "al is 8 bit, has a different opcode",
		output: []uint8{
			0xfe, 0xc0,
		},
	}, {
		input: input{
			INC, AX,
		},
		comment: "ax is 16 bit, need 0x66 prefix",
		output: []uint8{
			0x66, 0xff, 0xc0,
		},
	}}...)
}