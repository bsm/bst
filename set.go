package bst

// Set stores elements as a sorted slice. It is NOT thread-safe, please
// implement your own locking when reading and writing it from multiple
// goroutines.
type Set struct{ elements }

// NewSet creates a set with a given capacity
func NewSet(capacity int) *Set {
	return &Set{elements: make(elements, 0, capacity)}
}

// Add adds a value to the set
func (s *Set) Add(v Element) (ok bool) {
	s.elements, ok = s.elements.put(v, false)
	return
}

// Delete removes an item from the set, returns true if successful
func (s *Set) Delete(v Element) (ok bool) {
	s.elements, ok = s.elements.delete(v)
	return
}

// Iterator creates a new iterator
func (s *Set) Iterator() *SetIterator {
	return &SetIterator{baseIterator{pos: -1, elems: s.elements}}
}

// Intersects checks if the set is intersectable with other
func (s *Set) Intersects(other *Set) bool {
	ls, lo := len(s.elements), len(other.elements)
	if lo < ls {
		ls, lo = lo, ls
		s, other = other, s
	}
	if ls == 0 ||
		other.elements[lo-1].Less(s.elements[0]) ||
		other.elements[lo-1].Equals(s.elements[0]) ||
		s.elements[ls-1].Less(other.elements[0]) ||
		s.elements[ls-1].Equals(other.elements[0]) {
		return false
	}

	offset := 0
	for _, v := range s.elements {
		pos, ok := other.searchFrom(v, offset)
		if ok {
			return true
		} else if pos >= lo {
			return false
		}
		offset = pos
	}
	return false
}

// --------------------------------------------------------------------

// SetIterator is used to iterate over a set
type SetIterator struct{ baseIterator }

// Value returns the value at the current cursor position or nil if cursor is invalid
func (i *SetIterator) Value() Element {
	if i.valid(i.pos) {
		return i.elems[i.pos]
	}
	return nil
}
