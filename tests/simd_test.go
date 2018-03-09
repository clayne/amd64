package tests

import (
	"testing"
	"github.com/modern-go/test"
	"context"
	. "github.com/modern-go/amd64"
	"github.com/modern-go/test/must"
)

func Test_simd(t *testing.T) {
	t.Run("end to end", test.Case(func(ctx context.Context) {
		asm := &Assembler{}
		asm.Assemble(MOVD, XMM0, EDI)
		asm.Assemble(VPBROADCASTD, XMM0, XMM0)
		asm.Assemble(VPCMPEQD, XMM1, XMM0, XMMWORD(RSI, 0))
		asm.Assemble(VPACKSSDW, XMM1, XMM1, XMM2)
		must.Nil(asm.Error)
		must.Equal([]byte{
			0xc5, 0xf9, 0x6e, 0xc7, // vmovd xmm0,edi
			0xc4, 0xe2, 0x79, 0x58, 0xc0, // vpbroadcastd xmm0,xmm0
			0xc5, 0xf9, 0x76, 0x0e, // vpcmpeqd xmm1, xmm0, xmmword ptr [rsi]
			0xc5, 0xf1, 0x6b, 0xca, // vpackssdw xmm1, xmm1, xmm2
		},
			asm.Buffer)
	}))
}
