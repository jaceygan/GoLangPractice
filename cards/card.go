package main

import (
	"fmt"
	"strconv"
)

type card struct {
	value string
	suit  string
}

func (c card) print() {
	fmt.Println(c.value + c.suit)
}
func printCard(c card) string {
	return (c.value + c.suit)
}

func (c1 card) isLarger(c2 card) bool {
	c1v := c1.numValue()
	c2v := c2.numValue()
	if c1v != c2v {
		return (c1v > c2v)
	} else { //both cards same value
		c1s := c1.suitValue()
		c2s := c2.suitValue()
		return (c1s > c2s)
	}
}

func (c card) numValue() int {
	switch v := c.value; v {
	case "A":
		return 14
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	default:
		i, _ := strconv.Atoi(c.value)
		return i
	}
}

func (c card) suitValue() int {
	switch s := c.suit; s {
	case "S":
		return 4
	case "H":
		return 3
	case "C":
		return 2
	case "D":
		return 1
	default:
		return 0
	}
}

type testmap map[string]string

func testsomething() {

	tm := testmap{"abc": "123",
		"asd": "hah",
	}
	fmt.Println(tm)
}
