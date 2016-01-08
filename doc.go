// Package bst is a fast and generic implementation of sets/maps using
// binary-search-trees in Go.
// Performance is optimised towards reads. Both, inserts and lookups are
// binary searches and therefore O(log n). Others tructures, such as b-trees
// and skiplists may be better suited if you are looking for a more balanced
// performance.
//
// Please see examples for more details on usage.
package bst
