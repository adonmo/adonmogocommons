package set

type nothing struct{}
type Set struct {
	hash map[interface{}]nothing
}

func NewSet(initial []interface{}) *Set {
	set := &Set{
		hash: map[interface{}]nothing{},
	}

	for _, val := range initial {
		set.Insert(val)
	}

	return set
}

func (this *Set) Insert(val interface{}) {
	this.hash[val] = nothing{}
}

func (this *Set) Remove(val interface{}) {
	delete(this.hash, val)
}

func (this *Set) Has(val interface{}) bool {
	_, exists := this.hash[val]
	return exists
}

func (this *Set) Len() int {
	return len(this.hash)
}

func (this *Set) Values() []interface{} {
	values := []interface{}{}

	for k := range this.hash {
		values = append(values, k)
	}

	return values
}

func (this *Set) Map(f func(interface{}) interface{}) {

	hash := make(map[interface{}]nothing)

	for k := range this.hash {
		hash[f(k)] = nothing{}
	}

	this.hash = hash
}
