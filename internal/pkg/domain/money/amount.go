package money

import "strconv"

type Amount int64

func (a Amount) Float64() float64 {
	return float64(a) / 100
}

func (a Amount) String() string {
	return strconv.Itoa(int(a))
}

func (a Amount) Int64() int64 {
	return int64(a)
}

func NewZeroAmount() Amount {
	return Amount(0)
}
