package set

type Set []interface{}

func (set *Set) Cap () int {
	return cap(*set)
}

func (set *Set) Len () int {
	return len(*set)
}

func (set *Set) Insert (x interface{}, f func(interface{},interface{}) bool) {
	for _, e := range (*set) {
		if f(x,e) {
			return
		}
	}
	(*set) = append((*set), x)
}

func (set *Set) Remove (x interface{}, f func(interface{}, interface{}) bool) {
	var filtered Set
	for i, e := range (*set) {
		if f(x,e) == false {
			filtered = append(filtered, (*set)[i])
		}
	}
	*set = filtered
}

func (set *Set) Contains (x interface{}, f func(interface{}, interface{}) bool) bool {
	for _, e := range (*set) {
		if f(x,e) {
			return true
		}
	}
	return false
}

func (set *Set) Union (other *Set, f func(interface{}, interface{}) bool) {
	for _, e := range (*other) {
		set.Insert(e, f)
	}
}

func (set *Set) Filter (f func(interface{}) bool) *Set {
	var filtered Set
	for _, e := range (*set) {
		if f(e) {
			filtered = append(filtered, e)
		}
	}
	return &filtered
}

func (set *Set) Map (f func(interface{})) {
	for _, e := range (*set) {
		f(e)
	}
}

func (set *Set) MapWith (x interface{}, f func(interface{}, interface{})) {
	for _, e := range (*set) {
		f(x,e)
	}
}