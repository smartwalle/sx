package sx

type MemoryStock struct {
	words []string
}

func NewMemoryStock(words ...string) (*MemoryStock, error) {
	var m = &MemoryStock{}
	m.words = words
	return m, nil
}

func (m *MemoryStock) ReadAll() []string {
	return m.words
}
