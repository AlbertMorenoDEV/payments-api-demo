package money

type Money struct {
	amount   Amount
	currency Currency
}

func New(amount Amount, currency Currency) Money {
	return Money{amount: amount, currency: currency}
}

func NewFromPrimitives(amount int64, currency string) Money {
	return Money{amount: Amount(amount), currency: Currency(currency)}
}

func NewSGD(amount Amount) Money {
	return Money{amount: amount, currency: SGD}
}

func (m Money) Amount() Amount {
	return m.amount
}

func (m Money) Currency() Currency {
	return m.currency
}

func (m Money) Add(amount Money) Money {
	return Money{
		amount:   m.amount + amount.Amount(),
		currency: m.Currency(),
	}
}

func (m Money) Subtract(amount Money) Money {
	return Money{
		amount:   m.amount - amount.Amount(),
		currency: m.Currency(),
	}
}

func (m Money) IsLessThan(amount Money) bool {
	// ToDo check currency

	return m.amount < amount.Amount()
}
