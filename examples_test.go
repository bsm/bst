package bst_test

import (
	"fmt"

	"github.com/bsm/bst"
)

func ExampleSet() {
	set := bst.NewSet(5)
	set.Add(bst.Int(3))
	set.Add(bst.Int(5))
	set.Add(bst.Int(1))
	set.Add(bst.Int(3))

	fmt.Println(set.Len())
	// Output: 3
}

func ExampleSet_Add() {
	set := bst.NewSet(5)

	fmt.Println(set.Add(bst.Int(1)))
	fmt.Println(set.Add(bst.Int(2)))
	fmt.Println(set.Add(bst.Int(1)))
	// Output:
	// true
	// true
	// false
}

func ExampleSetIterator() {
	set := bst.NewSet(5)
	set.Add(bst.Int(3))
	set.Add(bst.Int(5))
	set.Add(bst.Int(1))
	set.Add(bst.Int(3))

	for iter := set.Iterator(); iter.Next(); {
		fmt.Println(iter.Value())
	}
	// Output:
	// 1
	// 3
	// 5
}

func ExampleMap() {
	set := bst.NewMap(5)
	set.Set(bst.Int(5), "bar")
	set.Set(bst.Int(3), "foo")

	value, ok := set.Get(bst.Int(3))
	fmt.Println(value, ok)
	value, ok = set.Get(bst.Int(4))
	fmt.Println(value, ok)
	// Output:
	// foo true
	// <nil> false
}

func ExampleMapIterator() {
	set := bst.NewMap(5)
	set.Set(bst.Int(5), "bar")
	set.Set(bst.Int(3), "foo")

	for iter := set.Iterator(); iter.Next(); {
		fmt.Println(iter.Key(), iter.Value())
	}
	// Output:
	// 3 foo
	// 5 bar
}
