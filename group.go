package rodeo

func (vaq *Vaquero) Tame(key string, representative interface{}) (gr *Group, e error) {
	// TODO: delegate connection of Vaquero to Group
	facade, e := connect(vaq.Conf.Host, vaq.Conf.Port)
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

type Group struct {
	key            string
	representative interface{}
	elements       []Element
	facade         pFacade
}

func (gr *Group) Add(score int64, v interface{}) (e error) {
	// TODO: validate type of v to equal representative
	return gr.facade.ZAdd(gr.key, score, v)
}
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
func (gr *Group) Count() (int, error) {
	return gr.facade.ZCount(gr.key)
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
