package main

// Player ...
type Player struct {
	letter string
}

// NewPlayer ...
func NewPlayer(letter string) *Player {
	return &Player{letter: letter}
}

// String ...
func (p *Player) String() string {
	return p.letter
}
