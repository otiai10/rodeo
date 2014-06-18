package rodeo

import "fmt"

func (vaq *Vaquero) Tame(dataName string, representative interface{}) (gr *Group, e error) {
	// TODO: delegate connection of Vaquero to Group
	gr = &Group{
		representative: representative,
	}
	return
}

type Group struct {
	representative interface{}
	elements       []Element
}

func (gr *Group) Add(score int64, v interface{}) (e error) {
	// TODO: validate type of v to equal representative
	gr.elements = append(gr.elements, Element{v, score})
	return
}
func (gr *Group) Find(i int) (el Element, e error) {
	if len(gr.elements) <= i {
		// TODO: common method to create error with `[rodeo]` prefix
		e = fmt.Errorf("Element for index `%v` not found in this group", i)
		return
	}
	el = gr.elements[i]
	return
}
func (gr *Group) Count() int {
	return len(gr.elements)
}

type Element struct {
	v     interface{}
	score int64
}

func (el *Element) Score() int64 {
	return el.score
}
func (el *Element) Retrieve() interface{} {
	return el.v
}
