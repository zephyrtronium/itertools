/*
Package itertools implements a number of operations on "iterators". The
package defines an iterator as a receive-only channel. Functions are
generated from templates for each built-in type; slices, arrays,
structs, and interfaces must use the interface{} functions. For most of
these functions, there is no suffix on the function name; for adapters
and generators, the suffix is Any. Cycle, CycleAndStep, Range, and
Accumulate use ints.

Strings are handled specially by AdaptString, GenerateString, and
CycleRune. GenerateString takes a rune iterator to produce a string;
the other two take strings and produce rune iterators. Use *Int32 tools
on rune iterators.
*/
package itertools
