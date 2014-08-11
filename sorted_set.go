package rodeo

// Tame provide active model of sorted sets specified by key name.
func (v *Vaquero) Tame(key string, representative interface{}) (ss *SortedSet, e error) {
	// TODO: delegate connection of Vaquero to SortedSet
	facade, e := connect(v.Conf.Host, v.Conf.Port)
	if e != nil {
		return
	}
	ss = &SortedSet{
		key:            key,
		representative: representative,
		facade:         facade,
	}
	return
}

// SortedSet is active model of sorted sets.
type SortedSet struct {
	key            string
	representative interface{}
	elements       []Element
	facade         pFacade
}

// Add adds value to SortedSet.
func (ss *SortedSet) Add(score int64, v interface{}) (e error) {
	// TODO: validate type of v to equal representative
	return ss.facade.ZAdd(ss.key, score, v)
}

// Range finds scored values in SortedSet by rank.
func (ss *SortedSet) Range(startStop ...int) (elements []*Element) {
	stuff := ss.representative
	scoredValues := ss.facade.ZRange(
		ss.key,
		startStop,
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

// Find finds scored values in SortedSet by score.
func (ss *SortedSet) Find(min int64, max int64) (elements []*Element) {
	stuff := ss.representative
	scoredValues := ss.facade.ZRangeByScore(
		ss.key,
		min,
		max,
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

// Count counts the values of SortedSet.
func (ss *SortedSet) Count() (int, error) {
	return ss.facade.ZCount(ss.key)
}

// Element is a model of SortedSet.
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
