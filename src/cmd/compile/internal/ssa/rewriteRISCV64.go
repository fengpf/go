// Code generated from gen/RISCV64.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "math"
import "cmd/compile/internal/types"

func rewriteValueRISCV64(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		v.Op = OpRISCV64ADD
		return true
	case OpAdd32:
		v.Op = OpRISCV64ADD
		return true
	case OpAdd32F:
		v.Op = OpRISCV64FADDS
		return true
	case OpAdd64:
		v.Op = OpRISCV64ADD
		return true
	case OpAdd64F:
		v.Op = OpRISCV64FADDD
		return true
	case OpAdd8:
		v.Op = OpRISCV64ADD
		return true
	case OpAddPtr:
		v.Op = OpRISCV64ADD
		return true
	case OpAddr:
		v.Op = OpRISCV64MOVaddr
		return true
	case OpAnd16:
		v.Op = OpRISCV64AND
		return true
	case OpAnd32:
		v.Op = OpRISCV64AND
		return true
	case OpAnd64:
		v.Op = OpRISCV64AND
		return true
	case OpAnd8:
		v.Op = OpRISCV64AND
		return true
	case OpAndB:
		v.Op = OpRISCV64AND
		return true
	case OpAvg64u:
		return rewriteValueRISCV64_OpAvg64u(v)
	case OpClosureCall:
		v.Op = OpRISCV64CALLclosure
		return true
	case OpCom16:
		v.Op = OpRISCV64NOT
		return true
	case OpCom32:
		v.Op = OpRISCV64NOT
		return true
	case OpCom64:
		v.Op = OpRISCV64NOT
		return true
	case OpCom8:
		v.Op = OpRISCV64NOT
		return true
	case OpConst16:
		v.Op = OpRISCV64MOVHconst
		return true
	case OpConst32:
		v.Op = OpRISCV64MOVWconst
		return true
	case OpConst32F:
		return rewriteValueRISCV64_OpConst32F(v)
	case OpConst64:
		v.Op = OpRISCV64MOVDconst
		return true
	case OpConst64F:
		return rewriteValueRISCV64_OpConst64F(v)
	case OpConst8:
		v.Op = OpRISCV64MOVBconst
		return true
	case OpConstBool:
		v.Op = OpRISCV64MOVBconst
		return true
	case OpConstNil:
		return rewriteValueRISCV64_OpConstNil(v)
	case OpConvert:
		v.Op = OpRISCV64MOVconvert
		return true
	case OpCvt32Fto32:
		v.Op = OpRISCV64FCVTWS
		return true
	case OpCvt32Fto64:
		v.Op = OpRISCV64FCVTLS
		return true
	case OpCvt32Fto64F:
		v.Op = OpRISCV64FCVTDS
		return true
	case OpCvt32to32F:
		v.Op = OpRISCV64FCVTSW
		return true
	case OpCvt32to64F:
		v.Op = OpRISCV64FCVTDW
		return true
	case OpCvt64Fto32:
		v.Op = OpRISCV64FCVTWD
		return true
	case OpCvt64Fto32F:
		v.Op = OpRISCV64FCVTSD
		return true
	case OpCvt64Fto64:
		v.Op = OpRISCV64FCVTLD
		return true
	case OpCvt64to32F:
		v.Op = OpRISCV64FCVTSL
		return true
	case OpCvt64to64F:
		v.Op = OpRISCV64FCVTDL
		return true
	case OpCvtBoolToUint8:
		v.Op = OpCopy
		return true
	case OpDiv16:
		return rewriteValueRISCV64_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueRISCV64_OpDiv16u(v)
	case OpDiv32:
		v.Op = OpRISCV64DIVW
		return true
	case OpDiv32F:
		v.Op = OpRISCV64FDIVS
		return true
	case OpDiv32u:
		v.Op = OpRISCV64DIVUW
		return true
	case OpDiv64:
		v.Op = OpRISCV64DIV
		return true
	case OpDiv64F:
		v.Op = OpRISCV64FDIVD
		return true
	case OpDiv64u:
		v.Op = OpRISCV64DIVU
		return true
	case OpDiv8:
		return rewriteValueRISCV64_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueRISCV64_OpDiv8u(v)
	case OpEq16:
		return rewriteValueRISCV64_OpEq16(v)
	case OpEq32:
		return rewriteValueRISCV64_OpEq32(v)
	case OpEq32F:
		v.Op = OpRISCV64FEQS
		return true
	case OpEq64:
		return rewriteValueRISCV64_OpEq64(v)
	case OpEq64F:
		v.Op = OpRISCV64FEQD
		return true
	case OpEq8:
		return rewriteValueRISCV64_OpEq8(v)
	case OpEqB:
		return rewriteValueRISCV64_OpEqB(v)
	case OpEqPtr:
		return rewriteValueRISCV64_OpEqPtr(v)
	case OpGeq32F:
		return rewriteValueRISCV64_OpGeq32F(v)
	case OpGeq64F:
		return rewriteValueRISCV64_OpGeq64F(v)
	case OpGetCallerPC:
		v.Op = OpRISCV64LoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpRISCV64LoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpRISCV64LoweredGetClosurePtr
		return true
	case OpGreater32F:
		return rewriteValueRISCV64_OpGreater32F(v)
	case OpGreater64F:
		return rewriteValueRISCV64_OpGreater64F(v)
	case OpHmul32:
		return rewriteValueRISCV64_OpHmul32(v)
	case OpHmul32u:
		return rewriteValueRISCV64_OpHmul32u(v)
	case OpHmul64:
		v.Op = OpRISCV64MULH
		return true
	case OpHmul64u:
		v.Op = OpRISCV64MULHU
		return true
	case OpInterCall:
		v.Op = OpRISCV64CALLinter
		return true
	case OpIsInBounds:
		v.Op = OpLess64U
		return true
	case OpIsNonNil:
		return rewriteValueRISCV64_OpIsNonNil(v)
	case OpIsSliceInBounds:
		v.Op = OpLeq64U
		return true
	case OpLeq16:
		return rewriteValueRISCV64_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueRISCV64_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueRISCV64_OpLeq32(v)
	case OpLeq32F:
		v.Op = OpRISCV64FLES
		return true
	case OpLeq32U:
		return rewriteValueRISCV64_OpLeq32U(v)
	case OpLeq64:
		return rewriteValueRISCV64_OpLeq64(v)
	case OpLeq64F:
		v.Op = OpRISCV64FLED
		return true
	case OpLeq64U:
		return rewriteValueRISCV64_OpLeq64U(v)
	case OpLeq8:
		return rewriteValueRISCV64_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueRISCV64_OpLeq8U(v)
	case OpLess16:
		return rewriteValueRISCV64_OpLess16(v)
	case OpLess16U:
		return rewriteValueRISCV64_OpLess16U(v)
	case OpLess32:
		return rewriteValueRISCV64_OpLess32(v)
	case OpLess32F:
		v.Op = OpRISCV64FLTS
		return true
	case OpLess32U:
		return rewriteValueRISCV64_OpLess32U(v)
	case OpLess64:
		v.Op = OpRISCV64SLT
		return true
	case OpLess64F:
		v.Op = OpRISCV64FLTD
		return true
	case OpLess64U:
		v.Op = OpRISCV64SLTU
		return true
	case OpLess8:
		return rewriteValueRISCV64_OpLess8(v)
	case OpLess8U:
		return rewriteValueRISCV64_OpLess8U(v)
	case OpLoad:
		return rewriteValueRISCV64_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueRISCV64_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueRISCV64_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueRISCV64_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueRISCV64_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueRISCV64_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueRISCV64_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueRISCV64_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueRISCV64_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueRISCV64_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValueRISCV64_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValueRISCV64_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValueRISCV64_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValueRISCV64_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValueRISCV64_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueRISCV64_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueRISCV64_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueRISCV64_OpLsh8x8(v)
	case OpMod16:
		return rewriteValueRISCV64_OpMod16(v)
	case OpMod16u:
		return rewriteValueRISCV64_OpMod16u(v)
	case OpMod32:
		v.Op = OpRISCV64REMW
		return true
	case OpMod32u:
		v.Op = OpRISCV64REMUW
		return true
	case OpMod64:
		v.Op = OpRISCV64REM
		return true
	case OpMod64u:
		v.Op = OpRISCV64REMU
		return true
	case OpMod8:
		return rewriteValueRISCV64_OpMod8(v)
	case OpMod8u:
		return rewriteValueRISCV64_OpMod8u(v)
	case OpMove:
		return rewriteValueRISCV64_OpMove(v)
	case OpMul16:
		return rewriteValueRISCV64_OpMul16(v)
	case OpMul32:
		v.Op = OpRISCV64MULW
		return true
	case OpMul32F:
		v.Op = OpRISCV64FMULS
		return true
	case OpMul64:
		v.Op = OpRISCV64MUL
		return true
	case OpMul64F:
		v.Op = OpRISCV64FMULD
		return true
	case OpMul8:
		return rewriteValueRISCV64_OpMul8(v)
	case OpNeg16:
		v.Op = OpRISCV64NEG
		return true
	case OpNeg32:
		v.Op = OpRISCV64NEG
		return true
	case OpNeg32F:
		v.Op = OpRISCV64FNEGS
		return true
	case OpNeg64:
		v.Op = OpRISCV64NEG
		return true
	case OpNeg64F:
		v.Op = OpRISCV64FNEGD
		return true
	case OpNeg8:
		v.Op = OpRISCV64NEG
		return true
	case OpNeq16:
		return rewriteValueRISCV64_OpNeq16(v)
	case OpNeq32:
		return rewriteValueRISCV64_OpNeq32(v)
	case OpNeq32F:
		v.Op = OpRISCV64FNES
		return true
	case OpNeq64:
		return rewriteValueRISCV64_OpNeq64(v)
	case OpNeq64F:
		v.Op = OpRISCV64FNED
		return true
	case OpNeq8:
		return rewriteValueRISCV64_OpNeq8(v)
	case OpNeqB:
		v.Op = OpRISCV64XOR
		return true
	case OpNeqPtr:
		return rewriteValueRISCV64_OpNeqPtr(v)
	case OpNilCheck:
		v.Op = OpRISCV64LoweredNilCheck
		return true
	case OpNot:
		return rewriteValueRISCV64_OpNot(v)
	case OpOffPtr:
		return rewriteValueRISCV64_OpOffPtr(v)
	case OpOr16:
		v.Op = OpRISCV64OR
		return true
	case OpOr32:
		v.Op = OpRISCV64OR
		return true
	case OpOr64:
		v.Op = OpRISCV64OR
		return true
	case OpOr8:
		v.Op = OpRISCV64OR
		return true
	case OpOrB:
		v.Op = OpRISCV64OR
		return true
	case OpPanicBounds:
		return rewriteValueRISCV64_OpPanicBounds(v)
	case OpRISCV64ADD:
		return rewriteValueRISCV64_OpRISCV64ADD(v)
	case OpRISCV64ADDI:
		return rewriteValueRISCV64_OpRISCV64ADDI(v)
	case OpRISCV64MOVBUload:
		return rewriteValueRISCV64_OpRISCV64MOVBUload(v)
	case OpRISCV64MOVBload:
		return rewriteValueRISCV64_OpRISCV64MOVBload(v)
	case OpRISCV64MOVBstore:
		return rewriteValueRISCV64_OpRISCV64MOVBstore(v)
	case OpRISCV64MOVBstorezero:
		return rewriteValueRISCV64_OpRISCV64MOVBstorezero(v)
	case OpRISCV64MOVDconst:
		return rewriteValueRISCV64_OpRISCV64MOVDconst(v)
	case OpRISCV64MOVDload:
		return rewriteValueRISCV64_OpRISCV64MOVDload(v)
	case OpRISCV64MOVDstore:
		return rewriteValueRISCV64_OpRISCV64MOVDstore(v)
	case OpRISCV64MOVDstorezero:
		return rewriteValueRISCV64_OpRISCV64MOVDstorezero(v)
	case OpRISCV64MOVHUload:
		return rewriteValueRISCV64_OpRISCV64MOVHUload(v)
	case OpRISCV64MOVHload:
		return rewriteValueRISCV64_OpRISCV64MOVHload(v)
	case OpRISCV64MOVHstore:
		return rewriteValueRISCV64_OpRISCV64MOVHstore(v)
	case OpRISCV64MOVHstorezero:
		return rewriteValueRISCV64_OpRISCV64MOVHstorezero(v)
	case OpRISCV64MOVWUload:
		return rewriteValueRISCV64_OpRISCV64MOVWUload(v)
	case OpRISCV64MOVWload:
		return rewriteValueRISCV64_OpRISCV64MOVWload(v)
	case OpRISCV64MOVWstore:
		return rewriteValueRISCV64_OpRISCV64MOVWstore(v)
	case OpRISCV64MOVWstorezero:
		return rewriteValueRISCV64_OpRISCV64MOVWstorezero(v)
	case OpRISCV64SUB:
		return rewriteValueRISCV64_OpRISCV64SUB(v)
	case OpRISCV64SUBW:
		return rewriteValueRISCV64_OpRISCV64SUBW(v)
	case OpRotateLeft16:
		return rewriteValueRISCV64_OpRotateLeft16(v)
	case OpRotateLeft32:
		return rewriteValueRISCV64_OpRotateLeft32(v)
	case OpRotateLeft64:
		return rewriteValueRISCV64_OpRotateLeft64(v)
	case OpRotateLeft8:
		return rewriteValueRISCV64_OpRotateLeft8(v)
	case OpRound32F:
		v.Op = OpCopy
		return true
	case OpRound64F:
		v.Op = OpCopy
		return true
	case OpRsh16Ux16:
		return rewriteValueRISCV64_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueRISCV64_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueRISCV64_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueRISCV64_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueRISCV64_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueRISCV64_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueRISCV64_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueRISCV64_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueRISCV64_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueRISCV64_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueRISCV64_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueRISCV64_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueRISCV64_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueRISCV64_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueRISCV64_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueRISCV64_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValueRISCV64_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValueRISCV64_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValueRISCV64_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValueRISCV64_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValueRISCV64_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValueRISCV64_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValueRISCV64_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValueRISCV64_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValueRISCV64_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueRISCV64_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueRISCV64_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueRISCV64_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueRISCV64_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueRISCV64_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueRISCV64_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueRISCV64_OpRsh8x8(v)
	case OpSignExt16to32:
		return rewriteValueRISCV64_OpSignExt16to32(v)
	case OpSignExt16to64:
		return rewriteValueRISCV64_OpSignExt16to64(v)
	case OpSignExt32to64:
		return rewriteValueRISCV64_OpSignExt32to64(v)
	case OpSignExt8to16:
		return rewriteValueRISCV64_OpSignExt8to16(v)
	case OpSignExt8to32:
		return rewriteValueRISCV64_OpSignExt8to32(v)
	case OpSignExt8to64:
		return rewriteValueRISCV64_OpSignExt8to64(v)
	case OpSlicemask:
		return rewriteValueRISCV64_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpRISCV64FSQRTD
		return true
	case OpStaticCall:
		v.Op = OpRISCV64CALLstatic
		return true
	case OpStore:
		return rewriteValueRISCV64_OpStore(v)
	case OpSub16:
		v.Op = OpRISCV64SUB
		return true
	case OpSub32:
		v.Op = OpRISCV64SUB
		return true
	case OpSub32F:
		v.Op = OpRISCV64FSUBS
		return true
	case OpSub64:
		v.Op = OpRISCV64SUB
		return true
	case OpSub64F:
		v.Op = OpRISCV64FSUBD
		return true
	case OpSub8:
		v.Op = OpRISCV64SUB
		return true
	case OpSubPtr:
		v.Op = OpRISCV64SUB
		return true
	case OpTrunc16to8:
		v.Op = OpCopy
		return true
	case OpTrunc32to16:
		v.Op = OpCopy
		return true
	case OpTrunc32to8:
		v.Op = OpCopy
		return true
	case OpTrunc64to16:
		v.Op = OpCopy
		return true
	case OpTrunc64to32:
		v.Op = OpCopy
		return true
	case OpTrunc64to8:
		v.Op = OpCopy
		return true
	case OpWB:
		v.Op = OpRISCV64LoweredWB
		return true
	case OpXor16:
		v.Op = OpRISCV64XOR
		return true
	case OpXor32:
		v.Op = OpRISCV64XOR
		return true
	case OpXor64:
		v.Op = OpRISCV64XOR
		return true
	case OpXor8:
		v.Op = OpRISCV64XOR
		return true
	case OpZero:
		return rewriteValueRISCV64_OpZero(v)
	case OpZeroExt16to32:
		return rewriteValueRISCV64_OpZeroExt16to32(v)
	case OpZeroExt16to64:
		return rewriteValueRISCV64_OpZeroExt16to64(v)
	case OpZeroExt32to64:
		return rewriteValueRISCV64_OpZeroExt32to64(v)
	case OpZeroExt8to16:
		return rewriteValueRISCV64_OpZeroExt8to16(v)
	case OpZeroExt8to32:
		return rewriteValueRISCV64_OpZeroExt8to32(v)
	case OpZeroExt8to64:
		return rewriteValueRISCV64_OpZeroExt8to64(v)
	}
	return false
}
func rewriteValueRISCV64_OpAvg64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg64u <t> x y)
	// result: (ADD (ADD <t> (SRLI <t> [1] x) (SRLI <t> [1] y)) (ANDI <t> [1] (AND <t> x y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64ADD)
		v0 := b.NewValue0(v.Pos, OpRISCV64ADD, t)
		v1 := b.NewValue0(v.Pos, OpRISCV64SRLI, t)
		v1.AuxInt = 1
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpRISCV64SRLI, t)
		v2.AuxInt = 1
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpRISCV64ANDI, t)
		v3.AuxInt = 1
		v4 := b.NewValue0(v.Pos, OpRISCV64AND, t)
		v4.AddArg2(x, y)
		v3.AddArg(v4)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueRISCV64_OpConst32F(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Const32F [val])
	// result: (FMVSX (MOVWconst [int64(int32(math.Float32bits(float32(math.Float64frombits(uint64(val))))))]))
	for {
		val := v.AuxInt
		v.reset(OpRISCV64FMVSX)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVWconst, typ.UInt32)
		v0.AuxInt = int64(int32(math.Float32bits(float32(math.Float64frombits(uint64(val))))))
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpConst64F(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Const64F [val])
	// result: (FMVDX (MOVDconst [val]))
	for {
		val := v.AuxInt
		v.reset(OpRISCV64FMVDX)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v0.AuxInt = val
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVDconst [0])
	for {
		v.reset(OpRISCV64MOVDconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueRISCV64_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// result: (DIVW (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (DIVUW (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64DIVUW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (DIVW (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (DIVUW (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64DIVUW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (SEQZ (ZeroExt16to64 (SUB <x.Type> x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SEQZ)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpRISCV64SUB, x.Type)
		v1.AddArg2(x, y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32 x y)
	// result: (SEQZ (SUBW <x.Type> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SEQZ)
		v0 := b.NewValue0(v.Pos, OpRISCV64SUBW, x.Type)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpEq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64 x y)
	// result: (SEQZ (SUB <x.Type> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SEQZ)
		v0 := b.NewValue0(v.Pos, OpRISCV64SUB, x.Type)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (SEQZ (ZeroExt8to64 (SUB <x.Type> x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SEQZ)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpRISCV64SUB, x.Type)
		v1.AddArg2(x, y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (XORI [1] (XOR <typ.Bool> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64XORI)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpRISCV64XOR, typ.Bool)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpEqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (EqPtr x y)
	// result: (SEQZ (SUB <x.Type> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SEQZ)
		v0 := b.NewValue0(v.Pos, OpRISCV64SUB, x.Type)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpGeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Geq32F x y)
	// result: (FLES y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64FLES)
		v.AddArg2(y, x)
		return true
	}
}
func rewriteValueRISCV64_OpGeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Geq64F x y)
	// result: (FLED y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64FLED)
		v.AddArg2(y, x)
		return true
	}
}
func rewriteValueRISCV64_OpGreater32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Greater32F x y)
	// result: (FLTS y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64FLTS)
		v.AddArg2(y, x)
		return true
	}
}
func rewriteValueRISCV64_OpGreater64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Greater64F x y)
	// result: (FLTD y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64FLTD)
		v.AddArg2(y, x)
		return true
	}
}
func rewriteValueRISCV64_OpHmul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32 x y)
	// result: (SRAI [32] (MUL (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRAI)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpRISCV64MUL, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpHmul32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32u x y)
	// result: (SRLI [32] (MUL (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRLI)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpRISCV64MUL, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsNonNil p)
	// result: (NeqPtr (MOVDconst) p)
	for {
		p := v_0
		v.reset(OpNeqPtr)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v.AddArg2(v0, p)
		return true
	}
}
func rewriteValueRISCV64_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (Not (Less16 y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess16, typ.Bool)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (Not (Less16U y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess16U, typ.Bool)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32 x y)
	// result: (Not (Less32 y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess32, typ.Bool)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32U x y)
	// result: (Not (Less32U y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess32U, typ.Bool)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpLeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq64 x y)
	// result: (Not (Less64 y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess64, typ.Bool)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpLeq64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq64U x y)
	// result: (Not (Less64U y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess64U, typ.Bool)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (Not (Less8 y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess8, typ.Bool)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (Not (Less8U y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess8U, typ.Bool)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (SLT (SignExt16to64 x) (SignExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SLT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (SLTU (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SLTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32 x y)
	// result: (SLT (SignExt32to64 x) (SignExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SLT)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32U x y)
	// result: (SLTU (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SLTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (SLT (SignExt8to64 x) (SignExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SLT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (SLTU (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SLTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpRISCV64MOVBUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ( is8BitInt(t) && isSigned(t))
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpRISCV64MOVBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ( is8BitInt(t) && !isSigned(t))
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpRISCV64MOVBUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && isSigned(t))
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpRISCV64MOVHload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && !isSigned(t))
	// result: (MOVHUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpRISCV64MOVHUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) && isSigned(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpRISCV64MOVWload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) && !isSigned(t))
	// result: (MOVWUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpRISCV64MOVWUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpRISCV64MOVDload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (FMOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpRISCV64FMOVWload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (FMOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpRISCV64FMOVDload)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpLocalAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LocalAddr {sym} base _)
	// result: (MOVaddr {sym} base)
	for {
		sym := v.Aux
		base := v_0
		v.reset(OpRISCV64MOVaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueRISCV64_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 <t> x y)
	// result: (AND (SLL <t> x y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg16, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x32 <t> x y)
	// result: (AND (SLL <t> x y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg16, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh16x64 <t> x y)
	// result: (AND (SLL <t> x y) (Neg16 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg16, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 <t> x y)
	// result: (AND (SLL <t> x y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg16, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 <t> x y)
	// result: (AND (SLL <t> x y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg32, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x32 <t> x y)
	// result: (AND (SLL <t> x y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg32, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh32x64 <t> x y)
	// result: (AND (SLL <t> x y) (Neg32 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg32, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 <t> x y)
	// result: (AND (SLL <t> x y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg32, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x16 <t> x y)
	// result: (AND (SLL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x32 <t> x y)
	// result: (AND (SLL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh64x64 <t> x y)
	// result: (AND (SLL <t> x y) (Neg64 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x8 <t> x y)
	// result: (AND (SLL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 <t> x y)
	// result: (AND (SLL <t> x y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg8, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x32 <t> x y)
	// result: (AND (SLL <t> x y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg8, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh8x64 <t> x y)
	// result: (AND (SLL <t> x y) (Neg8 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg8, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 <t> x y)
	// result: (AND (SLL <t> x y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg8, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (REMW (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64REMW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (REMUW (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64REMUW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (REMW (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64REMW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (REMUW (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64REMUW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v_2
		v.copyOf(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpRISCV64MOVBstore)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVBload, typ.Int8)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (MOVHstore dst (MOVHload src mem) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpRISCV64MOVHstore)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVHload, typ.Int16)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpRISCV64MOVWstore)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVWload, typ.Int32)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [8] dst src mem)
	// result: (MOVDstore dst (MOVDload src mem) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpRISCV64MOVDstore)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVDload, typ.Int64)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// result: (LoweredMove [t.(*types.Type).Alignment()] dst src (ADDI <src.Type> [s-moveSize(t.(*types.Type).Alignment(), config)] src) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpRISCV64LoweredMove)
		v.AuxInt = t.(*types.Type).Alignment()
		v0 := b.NewValue0(v.Pos, OpRISCV64ADDI, src.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(src)
		v.AddArg4(dst, src, v0, mem)
		return true
	}
}
func rewriteValueRISCV64_OpMul16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mul16 x y)
	// result: (MULW (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64MULW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpMul8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mul8 x y)
	// result: (MULW (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64MULW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (SNEZ (ZeroExt16to64 (SUB <x.Type> x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SNEZ)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpRISCV64SUB, x.Type)
		v1.AddArg2(x, y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32 x y)
	// result: (SNEZ (SUBW <x.Type> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SNEZ)
		v0 := b.NewValue0(v.Pos, OpRISCV64SUBW, x.Type)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpNeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64 x y)
	// result: (SNEZ (SUB <x.Type> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SNEZ)
		v0 := b.NewValue0(v.Pos, OpRISCV64SUB, x.Type)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (SNEZ (ZeroExt8to64 (SUB <x.Type> x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SNEZ)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpRISCV64SUB, x.Type)
		v1.AddArg2(x, y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (NeqPtr x y)
	// result: (SNEZ (SUB <x.Type> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRISCV64SNEZ)
		v0 := b.NewValue0(v.Pos, OpRISCV64SUB, x.Type)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Not x)
	// result: (XORI [1] x)
	for {
		x := v_0
		v.reset(OpRISCV64XORI)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV64_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (OffPtr [off] ptr:(SP))
	// result: (MOVaddr [off] ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpRISCV64MOVaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// cond: is32Bit(off)
	// result: (ADDI [off] ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		if !(is32Bit(off)) {
			break
		}
		v.reset(OpRISCV64ADDI)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// result: (ADD (MOVDconst [off]) ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		v.reset(OpRISCV64ADD)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v0.AuxInt = off
		v.AddArg2(v0, ptr)
		return true
	}
}
func rewriteValueRISCV64_OpPanicBounds(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicBoundsA [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpRISCV64LoweredPanicBoundsA)
		v.AuxInt = kind
		v.AddArg3(x, y, mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicBoundsB [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpRISCV64LoweredPanicBoundsB)
		v.AuxInt = kind
		v.AddArg3(x, y, mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicBoundsC [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpRISCV64LoweredPanicBoundsC)
		v.AuxInt = kind
		v.AddArg3(x, y, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64ADD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADD (MOVDconst [off]) ptr)
	// cond: is32Bit(off)
	// result: (ADDI [off] ptr)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpRISCV64MOVDconst {
				continue
			}
			off := v_0.AuxInt
			ptr := v_1
			if !(is32Bit(off)) {
				continue
			}
			v.reset(OpRISCV64ADDI)
			v.AuxInt = off
			v.AddArg(ptr)
			return true
		}
		break
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64ADDI(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDI [c] (MOVaddr [d] {s} x))
	// cond: is32Bit(c+d)
	// result: (MOVaddr [c+d] {s} x)
	for {
		c := v.AuxInt
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpRISCV64MOVaddr)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}
	// match: (ADDI [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVBUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBUload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(base, mem)
		return true
	}
	// match: (MOVBUload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBUload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(base, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVBload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(base, mem)
		return true
	}
	// match: (MOVBload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(base, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVBstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstore [off1] {sym1} (MOVaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg3(base, val, mem)
		return true
	}
	// match: (MOVBstore [off1] {sym} (ADDI [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg3(base, val, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBconst [0]) mem)
	// result: (MOVBstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpRISCV64MOVBconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpRISCV64MOVBstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVBstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstorezero [off1] {sym1} (MOVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpRISCV64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBstorezero [off1] {sym} (ADDI [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVDconst(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (MOVDconst <t> [c])
	// cond: !is32Bit(c) && int32(c) < 0
	// result: (ADD (SLLI <t> [32] (MOVDconst [c>>32+1])) (MOVDconst [int64(int32(c))]))
	for {
		t := v.Type
		c := v.AuxInt
		if !(!is32Bit(c) && int32(c) < 0) {
			break
		}
		v.reset(OpRISCV64ADD)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v1.AuxInt = c>>32 + 1
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v2.AuxInt = int64(int32(c))
		v.AddArg2(v0, v2)
		return true
	}
	// match: (MOVDconst <t> [c])
	// cond: !is32Bit(c) && int32(c) >= 0
	// result: (ADD (SLLI <t> [32] (MOVDconst [c>>32+0])) (MOVDconst [int64(int32(c))]))
	for {
		t := v.Type
		c := v.AuxInt
		if !(!is32Bit(c) && int32(c) >= 0) {
			break
		}
		v.reset(OpRISCV64ADD)
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v1.AuxInt = c>>32 + 0
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v2.AuxInt = int64(int32(c))
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(base, mem)
		return true
	}
	// match: (MOVDload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(base, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstore [off1] {sym1} (MOVaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg3(base, val, mem)
		return true
	}
	// match: (MOVDstore [off1] {sym} (ADDI [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg3(base, val, mem)
		return true
	}
	// match: (MOVDstore [off] {sym} ptr (MOVDconst [0]) mem)
	// result: (MOVDstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpRISCV64MOVDconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpRISCV64MOVDstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVDstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstorezero [off1] {sym1} (MOVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVDstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpRISCV64MOVDstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDstorezero [off1] {sym} (ADDI [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVDstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVHUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHUload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(base, mem)
		return true
	}
	// match: (MOVHUload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHUload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(base, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVHload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(base, mem)
		return true
	}
	// match: (MOVHload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(base, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVHstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstore [off1] {sym1} (MOVaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg3(base, val, mem)
		return true
	}
	// match: (MOVHstore [off1] {sym} (ADDI [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg3(base, val, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHconst [0]) mem)
	// result: (MOVHstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpRISCV64MOVHconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpRISCV64MOVHstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVHstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstorezero [off1] {sym1} (MOVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpRISCV64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHstorezero [off1] {sym} (ADDI [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVWUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWUload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWUload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(base, mem)
		return true
	}
	// match: (MOVWUload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWUload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(base, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVWload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(base, mem)
		return true
	}
	// match: (MOVWload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(base, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVWstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstore [off1] {sym1} (MOVaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCV64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg3(base, val, mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} (ADDI [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg3(base, val, mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWconst [0]) mem)
	// result: (MOVWstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpRISCV64MOVWconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpRISCV64MOVWstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64MOVWstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstorezero [off1] {sym1} (MOVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpRISCV64MOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpRISCV64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWstorezero [off1] {sym} (ADDI [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpRISCV64ADDI {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCV64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64SUB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUB x (MOVBconst [val]))
	// cond: is32Bit(-val)
	// result: (ADDI [-val] x)
	for {
		x := v_0
		if v_1.Op != OpRISCV64MOVBconst {
			break
		}
		val := v_1.AuxInt
		if !(is32Bit(-val)) {
			break
		}
		v.reset(OpRISCV64ADDI)
		v.AuxInt = -val
		v.AddArg(x)
		return true
	}
	// match: (SUB x (MOVHconst [val]))
	// cond: is32Bit(-val)
	// result: (ADDI [-val] x)
	for {
		x := v_0
		if v_1.Op != OpRISCV64MOVHconst {
			break
		}
		val := v_1.AuxInt
		if !(is32Bit(-val)) {
			break
		}
		v.reset(OpRISCV64ADDI)
		v.AuxInt = -val
		v.AddArg(x)
		return true
	}
	// match: (SUB x (MOVWconst [val]))
	// cond: is32Bit(-val)
	// result: (ADDI [-val] x)
	for {
		x := v_0
		if v_1.Op != OpRISCV64MOVWconst {
			break
		}
		val := v_1.AuxInt
		if !(is32Bit(-val)) {
			break
		}
		v.reset(OpRISCV64ADDI)
		v.AuxInt = -val
		v.AddArg(x)
		return true
	}
	// match: (SUB x (MOVDconst [val]))
	// cond: is32Bit(-val)
	// result: (ADDI [-val] x)
	for {
		x := v_0
		if v_1.Op != OpRISCV64MOVDconst {
			break
		}
		val := v_1.AuxInt
		if !(is32Bit(-val)) {
			break
		}
		v.reset(OpRISCV64ADDI)
		v.AuxInt = -val
		v.AddArg(x)
		return true
	}
	// match: (SUB x (MOVBconst [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpRISCV64MOVBconst || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SUB x (MOVHconst [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpRISCV64MOVHconst || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SUB x (MOVWconst [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpRISCV64MOVWconst || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SUB x (MOVDconst [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpRISCV64MOVDconst || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SUB (MOVBconst [0]) x)
	// result: (NEG x)
	for {
		if v_0.Op != OpRISCV64MOVBconst || v_0.AuxInt != 0 {
			break
		}
		x := v_1
		v.reset(OpRISCV64NEG)
		v.AddArg(x)
		return true
	}
	// match: (SUB (MOVHconst [0]) x)
	// result: (NEG x)
	for {
		if v_0.Op != OpRISCV64MOVHconst || v_0.AuxInt != 0 {
			break
		}
		x := v_1
		v.reset(OpRISCV64NEG)
		v.AddArg(x)
		return true
	}
	// match: (SUB (MOVWconst [0]) x)
	// result: (NEG x)
	for {
		if v_0.Op != OpRISCV64MOVWconst || v_0.AuxInt != 0 {
			break
		}
		x := v_1
		v.reset(OpRISCV64NEG)
		v.AddArg(x)
		return true
	}
	// match: (SUB (MOVDconst [0]) x)
	// result: (NEG x)
	for {
		if v_0.Op != OpRISCV64MOVDconst || v_0.AuxInt != 0 {
			break
		}
		x := v_1
		v.reset(OpRISCV64NEG)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRISCV64SUBW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBW x (MOVWconst [0]))
	// result: (ADDIW [0] x)
	for {
		x := v_0
		if v_1.Op != OpRISCV64MOVWconst || v_1.AuxInt != 0 {
			break
		}
		v.reset(OpRISCV64ADDIW)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
	// match: (SUBW (MOVDconst [0]) x)
	// result: (NEGW x)
	for {
		if v_0.Op != OpRISCV64MOVDconst || v_0.AuxInt != 0 {
			break
		}
		x := v_1
		v.reset(OpRISCV64NEGW)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVHconst [c]))
	// result: (Or16 (Lsh16x64 <t> x (MOVHconst [c&15])) (Rsh16Ux64 <t> x (MOVHconst [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpRISCV64MOVHconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x64, t)
		v1 := b.NewValue0(v.Pos, OpRISCV64MOVHconst, typ.UInt16)
		v1.AuxInt = c & 15
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux64, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64MOVHconst, typ.UInt16)
		v3.AuxInt = -c & 15
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRotateLeft32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft32 <t> x (MOVWconst [c]))
	// result: (Or32 (Lsh32x64 <t> x (MOVWconst [c&31])) (Rsh32Ux64 <t> x (MOVWconst [-c&31])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpRISCV64MOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpLsh32x64, t)
		v1 := b.NewValue0(v.Pos, OpRISCV64MOVWconst, typ.UInt32)
		v1.AuxInt = c & 31
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh32Ux64, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64MOVWconst, typ.UInt32)
		v3.AuxInt = -c & 31
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRotateLeft64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft64 <t> x (MOVDconst [c]))
	// result: (Or64 (Lsh64x64 <t> x (MOVDconst [c&63])) (Rsh64Ux64 <t> x (MOVDconst [-c&63])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpRISCV64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpLsh64x64, t)
		v1 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v1.AuxInt = c & 63
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh64Ux64, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v3.AuxInt = -c & 63
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVBconst [c]))
	// result: (Or8 (Lsh8x64 <t> x (MOVBconst [c&7])) (Rsh8Ux64 <t> x (MOVBconst [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpRISCV64MOVBconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x64, t)
		v1 := b.NewValue0(v.Pos, OpRISCV64MOVBconst, typ.UInt8)
		v1.AuxInt = c & 7
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux64, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64MOVBconst, typ.UInt8)
		v3.AuxInt = -c & 7
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 <t> x y)
	// result: (AND (SRL <t> (ZeroExt16to64 x) y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg16, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 <t> x y)
	// result: (AND (SRL <t> (ZeroExt16to64 x) y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg16, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 <t> x y)
	// result: (AND (SRL <t> (ZeroExt16to64 x) y) (Neg16 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg16, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 <t> x y)
	// result: (AND (SRL <t> (ZeroExt16to64 x) y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg16, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 <t> x y)
	// result: (SRA <t> (SignExt16to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt16to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 <t> x y)
	// result: (SRA <t> (SignExt16to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt32to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 <t> x y)
	// result: (SRA <t> (SignExt16to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 <t> x y)
	// result: (SRA <t> (SignExt16to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt8to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 <t> x y)
	// result: (AND (SRL <t> (ZeroExt32to64 x) y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg32, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 <t> x y)
	// result: (AND (SRL <t> (ZeroExt32to64 x) y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg32, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux64 <t> x y)
	// result: (AND (SRL <t> (ZeroExt32to64 x) y) (Neg32 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg32, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 <t> x y)
	// result: (AND (SRL <t> (ZeroExt32to64 x) y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg32, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 <t> x y)
	// result: (SRA <t> (SignExt32to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt16to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x32 <t> x y)
	// result: (SRA <t> (SignExt32to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt32to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x64 <t> x y)
	// result: (SRA <t> (SignExt32to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 <t> x y)
	// result: (SRA <t> (SignExt32to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt8to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh64Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux16 <t> x y)
	// result: (AND (SRL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh64Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux32 <t> x y)
	// result: (AND (SRL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh64Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh64Ux64 <t> x y)
	// result: (AND (SRL <t> x y) (Neg64 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh64Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux8 <t> x y)
	// result: (AND (SRL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x16 <t> x y)
	// result: (SRA <t> x (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt16to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v1 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v1.AuxInt = -1
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg2(y, v1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueRISCV64_OpRsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x32 <t> x y)
	// result: (SRA <t> x (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt32to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v1 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v1.AuxInt = -1
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg2(y, v1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueRISCV64_OpRsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh64x64 <t> x y)
	// result: (SRA <t> x (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v1 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v1.AuxInt = -1
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg2(y, v1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueRISCV64_OpRsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x8 <t> x y)
	// result: (SRA <t> x (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt8to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v1 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v1.AuxInt = -1
		v2 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg2(y, v1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueRISCV64_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 <t> x y)
	// result: (AND (SRL <t> (ZeroExt8to64 x) y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg8, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 <t> x y)
	// result: (AND (SRL <t> (ZeroExt8to64 x) y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg8, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 <t> x y)
	// result: (AND (SRL <t> (ZeroExt8to64 x) y) (Neg8 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg8, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 <t> x y)
	// result: (AND (SRL <t> (ZeroExt8to64 x) y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64AND)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpNeg8, t)
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueRISCV64_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 <t> x y)
	// result: (SRA <t> (SignExt8to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt16to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 <t> x y)
	// result: (SRA <t> (SignExt8to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt32to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 <t> x y)
	// result: (SRA <t> (SignExt8to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] y))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 <t> x y)
	// result: (SRA <t> (SignExt8to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt8to64 y)))))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpRISCV64SRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpRISCV64OR, y.Type)
		v2 := b.NewValue0(v.Pos, OpRISCV64ADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCV64SLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(y, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueRISCV64_OpSignExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (SignExt16to32 <t> x)
	// result: (SRAI [48] (SLLI <t> [48] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRAI)
		v.AuxInt = 48
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 48
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpSignExt16to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (SignExt16to64 <t> x)
	// result: (SRAI [48] (SLLI <t> [48] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRAI)
		v.AuxInt = 48
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 48
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpSignExt32to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt32to64 <t> x)
	// result: (ADDIW [0] x)
	for {
		x := v_0
		v.reset(OpRISCV64ADDIW)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV64_OpSignExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (SignExt8to16 <t> x)
	// result: (SRAI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRAI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpSignExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (SignExt8to32 <t> x)
	// result: (SRAI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRAI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpSignExt8to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (SignExt8to64 <t> x)
	// result: (SRAI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRAI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (NOT (SRAI <t> [63] (ADDI <t> [-1] x)))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64NOT)
		v0 := b.NewValue0(v.Pos, OpRISCV64SRAI, t)
		v0.AuxInt = 63
		v1 := b.NewValue0(v.Pos, OpRISCV64ADDI, t)
		v1.AuxInt = -1
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpRISCV64MOVBstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpRISCV64MOVHstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)
	// result: (MOVWstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpRISCV64MOVWstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && !is64BitFloat(val.Type)
	// result: (MOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && !is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpRISCV64MOVDstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (FMOVWstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpRISCV64FMOVWstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (FMOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpRISCV64FMOVDstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueRISCV64_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v_1
		v.copyOf(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// result: (MOVBstore ptr (MOVBconst) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpRISCV64MOVBstore)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVBconst, typ.UInt8)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// result: (MOVHstore ptr (MOVHconst) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpRISCV64MOVHstore)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVHconst, typ.UInt16)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [4] ptr mem)
	// result: (MOVWstore ptr (MOVWconst) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpRISCV64MOVWstore)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVWconst, typ.UInt32)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [8] ptr mem)
	// result: (MOVDstore ptr (MOVDconst) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpRISCV64MOVDstore)
		v0 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// result: (LoweredZero [t.(*types.Type).Alignment()] ptr (ADD <ptr.Type> ptr (MOVDconst [s-moveSize(t.(*types.Type).Alignment(), config)])) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		ptr := v_0
		mem := v_1
		v.reset(OpRISCV64LoweredZero)
		v.AuxInt = t.(*types.Type).Alignment()
		v0 := b.NewValue0(v.Pos, OpRISCV64ADD, ptr.Type)
		v1 := b.NewValue0(v.Pos, OpRISCV64MOVDconst, typ.UInt64)
		v1.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg2(ptr, v1)
		v.AddArg3(ptr, v0, mem)
		return true
	}
}
func rewriteValueRISCV64_OpZeroExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (ZeroExt16to32 <t> x)
	// result: (SRLI [48] (SLLI <t> [48] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRLI)
		v.AuxInt = 48
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 48
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpZeroExt16to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (ZeroExt16to64 <t> x)
	// result: (SRLI [48] (SLLI <t> [48] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRLI)
		v.AuxInt = 48
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 48
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpZeroExt32to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (ZeroExt32to64 <t> x)
	// result: (SRLI [32] (SLLI <t> [32] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRLI)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 32
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpZeroExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (ZeroExt8to16 <t> x)
	// result: (SRLI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRLI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpZeroExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (ZeroExt8to32 <t> x)
	// result: (SRLI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRLI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV64_OpZeroExt8to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (ZeroExt8to64 <t> x)
	// result: (SRLI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpRISCV64SRLI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCV64SLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteBlockRISCV64(b *Block) bool {
	switch b.Kind {
	case BlockRISCV64BNE:
		// match: (BNE (SNEZ x) yes no)
		// result: (BNE x yes no)
		for b.Controls[0].Op == OpRISCV64SNEZ {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			b.resetWithControl(BlockRISCV64BNE, x)
			return true
		}
	case BlockIf:
		// match: (If cond yes no)
		// result: (BNE cond yes no)
		for {
			cond := b.Controls[0]
			b.resetWithControl(BlockRISCV64BNE, cond)
			return true
		}
	}
	return false
}
