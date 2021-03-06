# bst set/map

[![Build Status](https://travis-ci.org/bsm/bst.png?branch=master)](https://travis-ci.org/bsm/bst)
[![GoDoc](https://godoc.org/github.com/bsm/bst?status.png)](http://godoc.org/github.com/bsm/bst)
[![Go Report Card](https://goreportcard.com/badge/github.com/bsm/bst)](https://goreportcard.com/report/github.com/bsm/bst)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Fast and generic Set and Map implementations using binary-search-trees.

### Documentation

Full documentation is available on [GoDoc](http://godoc.org/github.com/bsm/bst)

### Examples

As a Set:

```go
{
	set := bst.NewSet(5)
	set.Add(bst.Int(3))
	set.Add(bst.Int(5))
	set.Add(bst.Int(1))
	set.Add(bst.Int(3))

	for iter := set.Iterator(); iter.Next(); {
		fmt.Println(iter.Value())
	}
}
```

As a Map:

```go
{
	set := bst.NewMap(5)
	set.Set(bst.Int(5), "bar")
	set.Set(bst.Int(3), "foo")

	for iter := set.Iterator(); iter.Next(); {
		fmt.Println(iter.Key(), iter.Value())
	}
}
```
