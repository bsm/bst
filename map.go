package bst

type entry struct {
	key Element
	val interface{}
}

func (e entry) Less(other Element) bool   { return e.key.Less(other.(entry).key) }
func (e entry) Equals(other Element) bool { return e.key.Equals(other.(entry).key) }

// --------------------------------------------------------------------

// Map is an extension of Set and allows to store additional information in the Element.
// It is NOT thread-safe, please implement your own locking when reading and writing it
// from multiple goroutines.
type Map struct{ elements }

// New creates a map with a given capacity
func NewMap(capacity int) *Map { return &Map{elements: make(elements, 0, capacity)} }

// Set adds a value on key, returns true if the key was already present and the value was replaced.
func (m *Map) Set(key Element, value interface{}) (replaced bool) {
	m.elements, replaced = m.put(entry{key, value}, true)
	replaced = !replaced
	return
}

// Add acts like Set, but only adds the key if it does not exist, returns true if successful.
func (m *Map) Add(key Element, value interface{}) (ok bool) {
	m.elements, ok = m.put(entry{key, value}, false)
	return
}

// Get retrieves a value from the map, returns false as a second argument to indicate that key was not found.
func (m *Map) Get(key Element) (value interface{}, found bool) {
	if pos := m.search(entry{key, nil}); pos < len(m.elements) {
		if stored := m.elements[pos].(entry); stored.key.Equals(key) {
			value = stored.val
			found = true
		}
	}
	return
}

// Delete removes an item from the map, returns true if successful
func (m *Map) Delete(key Element) (ok bool) {
	m.elements, ok = m.elements.delete(entry{key, nil})
	return
}

// Iterator creates a new iterator
func (m *Map) Iterator() *MapIterator {
	return &MapIterator{baseIterator{pos: -1, elems: m.elements}}
}

// Exists checks the existence
func (m *Map) Exists(key Element) bool { return m.elements.Exists(entry{key, nil}) }

// --------------------------------------------------------------------

// MapIterator is used to iterate over a map
type MapIterator struct{ baseIterator }

// Seek places the cursor to the value greater or equal to pivot
func (i *MapIterator) Seek(pivot Element) bool {
	return i.baseIterator.Seek(entry{pivot, nil})
}

// Key returns the current key or nil if invalid
func (i *MapIterator) Key() Element {
	if i.valid(i.pos) {
		return i.elems[i.pos].(entry).key
	}
	return nil
}

// Value returns the current key or nil if invalid
func (i *MapIterator) Value() interface{} {
	if i.valid(i.pos) {
		return i.elems[i.pos].(entry).val
	}
	return nil
}
