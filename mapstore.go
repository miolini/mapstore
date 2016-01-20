package mapstore

const StoreDefaultSize = 1024

type Entry struct {
	Key   string
	Value interface{}
}

type Store interface {
	Set(string, interface{})
	Get(string, interface{}) (interface{}, bool)
	GetOrSet(string, interface{}) (interface{}, bool)
	Delete(string) bool

	Load(chan Entry)
	Save(chan<- Entry)

	//for single map
	Len() int

	//for sharded map
	ShardStats() []int
}

func NewWithSize(shardsCount int) Store {
	if shardsCount > 1 {
		return newStoreShard(shardsCount)
	}
	return newStoreSingle()
}

func New() Store {
	return NewWithSize(StoreDefaultSize)
}
