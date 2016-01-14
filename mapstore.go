package mapstore

type Entry struct {
	Key   string
	Value interface{}
}

type Store interface {
	Set(string, interface{})
	Get(string, interface{}) (interface{}, bool)

	Load(chan Entry)
	Save(chan<- Entry)

	//for single map
	Len() int

	//for sharded map
	ShardStats() []int
}

func New(shardsCount int) Store {
	if shardsCount > 1 {
		return newStoreShard(shardsCount)
	}
	return newStoreSingle()
}
