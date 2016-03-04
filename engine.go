package main

type key uint

const (
	KeyEscape key = iota
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
)

type keys map[key]struct{}

var Keys = make(keys)

func (k keys) Down(ky key) {
	k[ky] = struct{}{}
}

func (k keys) Up(ky key) {
	delete(k[ky])
}

func (k keys) Pressed(ky key) bool {
	_, ok := k[ky]
	return ok
}
