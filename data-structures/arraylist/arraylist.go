package arraylist

type ArrayList struct {
	size int
	cap int
	data []interface{}
}

func NewArrayList() *ArrayList{
	data := make([]interface{}, 2, 2)
	return &ArrayList{
		size: 0,
		cap: 2,
		data: data,
	}
}

func(a *ArrayList) Append(val interface{}) {
	if a.size == a.cap {
		a.growArray(val)
	}

	a.data[a.size] = val
	a.size++
}

func(a *ArrayList) growArray(val interface{}) {
	a.cap = a.cap * 2
	newArr := make([]interface{}, a.cap, a.cap)
	for i, item := range a.data {
		newArr[i] = item
	}

	newArr[a.size] = val
	a.data = newArr
}

func(a *ArrayList) List() []interface{}{
	return a.data
}
