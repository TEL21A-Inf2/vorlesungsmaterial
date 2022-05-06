package main

import (
	"fmt"
)

type LinkedListElement struct {
	data int
	next *LinkedListElement
}

// Liefert true, falls das Element leer (d.h. ein Dummy) ist.
// Ein Element ist leer, falls es keinen Nachfolger hat.
func (element LinkedListElement) IsEmpty() bool {
	return element.next == nil
}

// Setzt den Nachfolger des gegebenen Elements.
func (element *LinkedListElement) SetNext(next *LinkedListElement) {
	element.next = next
}

// Liefert den Dummy am Ende der Liste, die bei element beginnt.
func (element *LinkedListElement) GetEnd() *LinkedListElement {
	if element.IsEmpty() {
		return element
	}
	return element.next.GetEnd()
}

// Hängt ein Element mit data am Ende an.
func (element *LinkedListElement) Append(data int) {
	end := element.GetEnd()
	end.data = data
	end.next = MakeDummy()
}

// Liefert den Wert, der im Element an Stelle pos steht.
func (element *LinkedListElement) GetValue(pos int) int {

	if element.IsEmpty() {
		return -1 // TODO: Funktioniert nur, wenn -1 nicht in der Liste vorkommt.
	}

	if pos == 0 {
		return element.data
	}
	return element.next.GetValue(pos - 1)

	/* Alternative:
	current := element
	for i := 0; i<pos; i++ {
	  current = current.next
	}

	return current.data
	*/
}

// Fügt ein Element mit dem gegebenen Wert an Stelle pos ein.
// Genauer: Zwischen Stelle pos-1 und pos.
func (element *LinkedListElement) Insert(pos, value int) *LinkedListElement {
	// Wenn pos == 0, dann erzeugen wir hier intern einen Dummy
	// und hängen dort element an.
	if pos == 0 {
		d := MakeDummy()
		d.data = value
		d.next = element
		return d
	}

	// Wenn pos == 1, dann wollen wir zwischen element und element.next
	// ein neues Element einfügen.
	if pos == 1 {
		n := MakeDummy()
		n.data = value
		n.next = element.next
		element.next = n
		return element
	}
	element.next.Insert(pos-1, value)
	return element
}

// Vertauscht die Elemente an den Stellen pos1 und pos2.
func (element *LinkedListElement) Swap(pos1, pos2 int) {
	// TODO
}

// Liefert die Länge der Liste.
func (element LinkedListElement) GetLength() int {
	if element.IsEmpty() {
		return 0
	}
	return element.next.GetLength() + 1
}

// Konstruktor für eine Liste: Erzeugt ein neues Dummy-Element
// und liefert einen Pointer darauf.
func MakeDummy() *LinkedListElement {
	return &LinkedListElement{0, nil}
}

type LinkedList struct {
	head *LinkedListElement
}

func (list *LinkedList) Append(value int) {
	list.head.Append(value)
}

func (list *LinkedList) Insert(pos, value int) {
	list.head = list.head.Insert(pos, value)
}

func MakeLinkedList() LinkedList {
	result := LinkedList{MakeDummy()}
	return result
}

func main() {

	e1 := MakeDummy()
	e1.Append(42)
	e1.Append(77)
	e1.Append(38)
	e1.Append(25)
	e1.Insert(1, 105)
	e1.Insert(3, 10000)
	e1 = e1.Insert(0, 9999999)

	// In einer Schleife alles ausgeben.
	for current := e1; !current.IsEmpty(); current = current.next {
		fmt.Println(current.data)
	}

	fmt.Println(e1.GetValue(6))
	fmt.Println(e1.GetValue(-1))

	l1 := MakeLinkedList()
	l1.Append(42)
	l1.Insert(0, 10000)

}
