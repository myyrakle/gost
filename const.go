package gost

import "math"

const U8_MIN = U8(0)
const U8_MAX = U8(math.MaxUint8)

const U16_MIN = U16(0)
const U16_MAX = U16(math.MaxUint16)

const U32_MIN = U32(0)
const U32_MAX = U32(math.MaxUint32)

const U64_MIN = U64(0)
const U64_MAX = U64(math.MaxUint64)

var U128_MIN = U128{
	high: U64_MIN,
	low:  U64_MIN,
}
var U128_MAX = U128{
	high: U64_MAX,
	low:  U64_MAX,
}

const USize_MIN = USize(0)
const USize_MAX = USize(math.MaxUint)

const I8_MIN = I8(math.MinInt8)
const I8_MAX = I8(math.MaxInt8)

const I16_MIN = I16(math.MinInt16)
const I16_MAX = I16(math.MaxInt16)

const I32_MIN = I32(math.MinInt32)
const I32_MAX = I32(math.MaxInt32)

const I64_MIN = I64(math.MinInt64)
const I64_MAX = I64(math.MaxInt64)

var I128_MIN = I128{
	high: I64_MIN,
	low:  U64_MIN,
}
var I128_MAX = I128{
	high: I64_MAX,
	low:  U64_MAX,
}

const ISize_MIN = ISize(math.MinInt)
const ISize_MAX = ISize(math.MaxInt)
