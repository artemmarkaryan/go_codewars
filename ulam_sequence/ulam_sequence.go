// https://www.codewars.com/kata/5ac94db76bde60383d000038/go

package ulam_sequence

import (
	"log"
)


type Deque struct {
	slice []int
}

func NewDeque(slice []int) Deque {
	return Deque{slice: slice}
}
func (d *Deque) Get() []int { return d.slice }

func (d *Deque) Len() int { return len(d.slice) }

func (d *Deque) PushBack(el int) {
	d.slice = append(d.slice, el)
}

func (d *Deque) PopFront() (el int) {
	if len(d.slice) == 0 {
		return
	}
	el = d.slice[0]
	d.slice = d.slice[1:]
	return
}

func (d *Deque) RemoveDuplicates() {
	unique := map[int]int{}

	for _, v := range d.slice {
		unique[v] += 1
	}

	//goland:noinspection ALL
	result := []int{}
	for k, v := range unique {
		if v == 1 {
			result = append(result, k)
		}
	}

	d.slice = result
}

func UlamSequence(u0, u1, n int) []int {
	var elements []int
	var newElements Deque

	elements = append(elements, u0, u1)
	//newElements.PushBack(u0 + u1)

	for len(elements) < n {
		log.Print("elements: ", elements)
		lastElement := elements[len(elements)-1]
		log.Print("lastElement: ", lastElement)

		for _, element := range elements[:len(elements)-1] {
			newElements.PushBack(element + lastElement)
		}

		log.Print("newElements: ", newElements)
		newElements.RemoveDuplicates()

		log.Print("newElements-duplicates: ", newElements)
		elements = append(elements, newElements.PopFront())
	}

	return elements
}
