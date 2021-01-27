package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type laser int

type martian struct{}

func (l laser) talk() string {
	return strings.Repeat("Pew ", int(l))
}

func (m martian) talk() string {
	return "nack nack"
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martian{})
	shout(laser(2))
	// 由于嵌入了laser所以starship能够转发laser的talk方法
	type starship struct {
		laser
	}
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)

	// test
	r := rover("whir whir")
	shout(r)
}

// test
type rover string

func (r rover) talk() string {
	return string(r)
}
