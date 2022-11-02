package event

type Storer interface {
	Store(event Event) error
}
