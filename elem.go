package bst

import "sort"

// Element represents a value that is stores in the set
type Element interface {
	// Must return true if element is less than the other
	Less(other Element) bool
	// Must return true if element is the equivalent of the other
	Equals(other Element) bool
}

// --------------------------------------------------------------------

// Int element type
type Int int

func (e Int) Less(other Element) bool   { return e < other.(Int) }
func (e Int) Equals(other Element) bool { return e == other.(Int) }

// Float64 element type
type Float64 float64

func (e Float64) Less(other Element) bool   { return e < other.(Float64) }
func (e Float64) Equals(other Element) bool { return e == other.(Float64) }

// String element type
type String string

func (e String) Less(other Element) bool   { return e < other.(String) }
func (e String) Equals(other Element) bool { return e == other.(String) }

// --------------------------------------------------------------------

type elements []Element

// Len returns the size
func (p elements) Len() int { return len(p) }

// Exists checks for the existence of the element
func (p elements) Exists(v Element) bool {
	pos := p.search(v)
	return pos < len(p) && p[pos].Equals(v)
}

func (p elements) search(v Element) int {
	return sort.Search(len(p), func(i int) bool { return !p[i].Less(v) })
}

func (p elements) searchFrom(v Element, offset int) (int, bool) {
	sub := p[offset:]
	pos := sort.Search(len(sub), func(i int) bool { return !sub[i].Less(v) }) + offset
	return pos, pos < len(p) && p[pos].Equals(v)
}

func (p elements) put(v Element, replace bool) (elements, bool) {
	if pos := p.search(v); pos < len(p) {
		if p[pos].Equals(v) {
			if replace {
				p[pos] = v
			}
			return p, false
		}

		p = append(p, v)
		copy(p[pos+1:], p[pos:])
		p[pos] = v
	} else {
		p = append(p, v)
	}
	return p, true
}

func (p elements) delete(v Element) (elements, bool) {
	if pos := p.search(v); pos < len(p) && p[pos].Equals(v) {
		p = p[:pos+copy(p[pos:], p[pos+1:])]
		return p, true
	}
	return p, false
}

// --------------------------------------------------------------------

type sortable []Element

func (p sortable) Len() int           { return len(p) }
func (p sortable) Less(i, j int) bool { return p[i].Less(p[j]) }
func (p sortable) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
