# Slices are Like References To Arrays

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

> **Source:** https://tour.golang.org/moretypes/8
