package set

type SetInt struct {
	set *Set
}

func NewSetInt(initial []int) *SetInt {
	values := []interface{}{}
	for _, val := range initial {
		values = append(values, val)
	}

	return &SetInt{
		set: NewSet(values),
	}
}

func (this *SetInt) Insert(val int) {
	this.set.Insert(val)
}

func (this *SetInt) Remove(val int) {
	this.set.Remove(val)
}

func (this *SetInt) Has(val int) bool {
	return this.set.Has(val)
}

func (this *SetInt) Len() int {
	return this.set.Len()
}

func (this *SetInt) Values() []int {
	values := []int{}

	for _, val := range this.set.Values() {
		values = append(values, val.(int))
	}

	return values
}

func (this *SetInt) Map(f func(int) int) {
	fInt := func(val interface{}) interface{} {
		return f(val.(int))
	}

	this.set.Map(fInt)
}
