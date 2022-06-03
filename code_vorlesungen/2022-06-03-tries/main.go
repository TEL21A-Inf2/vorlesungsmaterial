package main

import (
	"fmt"
)

type Element struct {
	key, value string
	children   map[string]*Element
}

func NewElement() *Element {
	return &Element{"", "", make(map[string]*Element)}
}

func Add(element *Element, key, value string) {
	if key == "" {
		element.value = value
		return
	}
	head, tail := key[:1], key[1:]
	element.key = head
	if tail == "" {
		return
	}
	element.children[tail[:1]] = NewElement()
	Add(element.children[tail[:1]], tail, value)
}

func GetValues(element *Element) []string {
	result := []string{}
	if element.value != "" {
		result = append(result, element.value)
	}
	if element.key == "" {
		return result
	}
	if len(element.children) == 0 {
		return result
	}
	for _, child := range element.children {
		childValues := GetValues(child)
		result = append(result, childValues...)
	}
	return result
}

func main() {

	e1 := NewElement()
	Add(e1, "Hallo", "Hallo")
	fmt.Println(GetValues(e1))

	fmt.Println(e1.key)
	fmt.Println(e1.children["a"].key)
	fmt.Println(e1.children["a"].children["l"].key)
}
