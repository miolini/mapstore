package mapstore

type Store struct {
  m sync.RWMutex
  s map[string]interface{}
}

func New() *Store {
  return &Store{
    s: make(map[string]interface{},
  }
}

func (s *Store) Get(key string, defaultValue interface{}) (interface{}, bool) {
  s.m.RLock()
  value, ok := s.s[key]
  s.m.RUnlock()
  if !ok {
    return defaultValue, false
  }
  return value, true
}

func (s *Store) Set(key string, value interface{}) {
  s.m.Lock()
  s.s[key] = value
  s.m.Unlock()
}

func (s *Store) Len() int {
  s.m.RLock()
  defer s.m.RUnlock()
  return len(s.s)
}
