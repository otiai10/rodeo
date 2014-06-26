package rodeo

// Tame provide active model of sorted sets specified by key name.
func (v *Vaquero) Tame(key string, representative interface{}) (gr *Group, e error) {
	// TODO: delegate connection of Vaquero to Group
	facade, e := connect(v.Conf.Host, v.Conf.Port)
	if e != nil {
		return
	}
	gr = &Group{
		key:            key,
		representative: representative,
		facade:         facade,
	}
	return
}

// Group is active model of sorted sets.
type Group struct {
	key            string
	representative interface{}
	elements       []Element
	facade         pFacade
}

// Add adds value to Group.
func (gr *Group) Add(score int64, v interface{}) (e error) {
	// TODO: validate type of v to equal representative
	return gr.facade.ZAdd(gr.key, score, v)
}

// Find finds scored values in Group.
func (gr *Group) Find(args ...int) (elements []*Element) {
	stuff := gr.representative
	scoredValues := gr.facade.ZRange(
		gr.key,
		args,
		stuff,
	)
	for _, scored := range scoredValues {
		el := &Element{
			scored.Value,
			scored.Score,
		}
		elements = append(elements, el)
	}
	return
}

// Count counts the values of Group.
func (gr *Group) Count() (int, error) {
	return gr.facade.ZCount(gr.key)
}

// Element is a model of Group.
type Element struct {
	v     interface{}
	score int64
}

// Score provides the score of Element.
func (el *Element) Score() int64 {
	return el.score
}

// Retrieve provides the value of Element in interface definition.
// so it requires type assertion in application layor.
func (el *Element) Retrieve() interface{} {
	return el.v
}
