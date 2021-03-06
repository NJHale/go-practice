# Slice Length and Capacity

A slice has both a `length` and a `capacity`.

The `length` of a slice is the number of elements it contains.
README>mdREADME>md
The `capacity` of a slice is the number of elements in the underlying array,
counting from the first element in the slice.

The length and capacity of a slice `s` can be obtained using the expressions
`len(s)` and `cap(s)`.

You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

> **Note:** Once elements at lower indicies of the slice are dropped, the overall capacity of
the slice is permanently reduced.

> **Source:** https://tour.golang.org/moretypes/11
