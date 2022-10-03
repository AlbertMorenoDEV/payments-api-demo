package money

const SGD = "SGD"

type Currency string

func (c Currency) String() string {
	return string(c)
}
