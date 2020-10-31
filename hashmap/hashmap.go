package hashmap

type HashMap struct {
	data [][]pair
	len int
}

type pair struct {
	key string
	val interface{}
}

func New(len int) *HashMap{
	data := make([][]pair, 20)
	return &HashMap{
		len: 20,
		data: data,
	}
}

func getHashCode(key string) int{
	var total int

	for _, char := range key {
		total = total + int(char)
	}

	return total
}

func(h *HashMap) Set(key string, val interface{}) {
	hashCode := getHashCode(key)
	index := hashCode % h.len
	hashPair := pair{
		key,
		val,
	}

	h.data[index] = append(h.data[index], hashPair)
}

func(h *HashMap) Get(key string) interface{}{
	hashCode := getHashCode(key)
	index := hashCode % h.len

	for _, bucket := range h.data[index] {
		if bucket.key == key {
			return bucket.val
		}
	}

	return nil
}