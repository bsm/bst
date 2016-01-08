package bst

type baseIterator struct {
	elems elements
	pos   int
}

func (i *baseIterator) valid(pos int) bool { return pos > -1 && pos < len(i.elems) }

// Next advances the cursor to the next position
func (i *baseIterator) Next() bool {
	if !i.valid(i.pos + 1) {
		return false
	}
	i.pos++
	return true
}

// Previous advances the cursor to the previous position
func (i *baseIterator) Previous() bool {
	if !i.valid(i.pos - 1) {
		return false
	}
	i.pos--
	return true
}

// Seek places the cursor to the value greater or equal to pivot
func (i *baseIterator) Seek(pivot Element) bool {
	pos := i.elems.search(pivot)
	if !i.valid(pos) {
		return false
	}
	i.pos = pos
	return true
}
