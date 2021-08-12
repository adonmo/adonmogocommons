package set

type SetString struct {
	set *Set
}

func NewSetString(initial ...string) *SetString {
	values := []interface{}{}
	for _, val := range initial {
		values = append(values, val)
	}

	return &SetString{
		set: NewSet(values...),
	}
}

func (this *SetString) Insert(val string) {
	this.set.Insert(val)
}

func (this *SetString) Remove(val string) {
	this.set.Remove(val)
}

func (this *SetString) Has(val string) bool {
	return this.set.Has(val)
}

func (this *SetString) Len() int {
	return this.set.Len()
}

func (this *SetString) Values() []string {
	values := []string{}

	for _, val := range this.set.Values() {
		values = append(values, val.(string))
	}

	return values
}

func (this *SetString) Map(f func(string) string) {
	fInt := func(val interface{}) interface{} {
		return f(val.(string))
	}

	this.set.Map(fInt)
}
