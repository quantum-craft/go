package constant

// MaxUint is MaxUint
const MaxUint = ^uint(0)

// MinUint is MinUint
const MinUint = uint(0)

// MaxInt is MaxInt
const MaxInt = int(MaxUint >> 1)

// MinInt is MinInt
const MinInt = -MaxInt - 1
