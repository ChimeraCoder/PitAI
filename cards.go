package main


const (
    BARLEY = iota
    CORN = iota
    WHEAT = iota
    ORANGES = iota
    SOYBEANS = iota
    SUGAR = iota
)

type Card struct {
    Good int
}


func (c Card) Matches(other Card) bool {
    c.Good == other.Good
}

