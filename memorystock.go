package sx

type MemoryStock struct {
	words []string
}

func NewMemoryStock(words ...string) (*MemoryStock, error) {
	var s = &MemoryStock{}
	s.words = words
	return s, nil
}

func (ms *MemoryStock) ReadAll() []string {
	return ms.words
}
