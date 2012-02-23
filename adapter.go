/*
This is free software. It comes without any warranty, to the extent permitted
by applicable law. You can redistribute it and/or modify it under the terms of
the Do What The Fuck You Want To Public License, Version 2, as published by
Sam Hocevar. See COPYING for more details.
*/

package itertools

// IMPORTANT: adapter.go is generated from adapter.got. Change the latter.

// AdaptString turns a string into a rune iterator.
func AdaptString(s string) <-chan rune {
	ch := make(chan rune, 1)
	go func() {
		for _, r := range s {
			ch <- r
		}
		close(ch)
	}()
	return ch
}

// AdaptStrings returns an iterator over the
// values of s.
func AdaptStrings(s []string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptBytes returns an iterator over the
// values of s.
func AdaptBytes(s []uint8) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptUint16s returns an iterator over the
// values of s.
func AdaptUint16s(s []uint16) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptUint32s returns an iterator over the
// values of s.
func AdaptUint32s(s []uint32) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptUint64s returns an iterator over the
// values of s.
func AdaptUint64s(s []uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptInt8s returns an iterator over the
// values of s.
func AdaptInt8s(s []int8) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptInt16s returns an iterator over the
// values of s.
func AdaptInt16s(s []int16) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptInt32s returns an iterator over the
// values of s.
func AdaptInt32s(s []int32) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptInt64s returns an iterator over the
// values of s.
func AdaptInt64s(s []int64) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptFloat32s returns an iterator over the
// values of s.
func AdaptFloat32s(s []float32) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptFloat64s returns an iterator over the
// values of s.
func AdaptFloat64s(s []float64) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptComplex64s returns an iterator over the
// values of s.
func AdaptComplex64s(s []complex64) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptComplex128s returns an iterator over the
// values of s.
func AdaptComplex128s(s []complex128) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptUints returns an iterator over the
// values of s.
func AdaptUints(s []uint) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptInts returns an iterator over the
// values of s.
func AdaptInts(s []int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptPtrs returns an iterator over the
// values of s.
func AdaptPtrs(s []uintptr) <-chan uintptr {
	ch := make(chan uintptr, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptBools returns an iterator over the
// values of s.
func AdaptBools(s []bool) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// AdaptAny returns an iterator over the
// values of s.
func AdaptAny(s []interface{}) <-chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}


// GenerateString turns a rune iterator into a string.
func GenerateString(ch <-chan rune) string {
	s := make([]rune, 0, 32)
	for r := range ch {
		s = append(s, r)
	}
	return string(s)
}

// GenerateStrings turns an iterator of string
// into a []string.
func GenerateStrings(ch <-chan string) []string {
	s := make([]string, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateBytes turns an iterator of uint8
// into a []uint8.
func GenerateBytes(ch <-chan uint8) []uint8 {
	s := make([]uint8, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateUint16s turns an iterator of uint16
// into a []uint16.
func GenerateUint16s(ch <-chan uint16) []uint16 {
	s := make([]uint16, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateUint32s turns an iterator of uint32
// into a []uint32.
func GenerateUint32s(ch <-chan uint32) []uint32 {
	s := make([]uint32, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateUint64s turns an iterator of uint64
// into a []uint64.
func GenerateUint64s(ch <-chan uint64) []uint64 {
	s := make([]uint64, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateInt8s turns an iterator of int8
// into a []int8.
func GenerateInt8s(ch <-chan int8) []int8 {
	s := make([]int8, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateInt16s turns an iterator of int16
// into a []int16.
func GenerateInt16s(ch <-chan int16) []int16 {
	s := make([]int16, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateInt32s turns an iterator of int32
// into a []int32.
func GenerateInt32s(ch <-chan int32) []int32 {
	s := make([]int32, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateInt64s turns an iterator of int64
// into a []int64.
func GenerateInt64s(ch <-chan int64) []int64 {
	s := make([]int64, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateFloat32s turns an iterator of float32
// into a []float32.
func GenerateFloat32s(ch <-chan float32) []float32 {
	s := make([]float32, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateFloat64s turns an iterator of float64
// into a []float64.
func GenerateFloat64s(ch <-chan float64) []float64 {
	s := make([]float64, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateComplex64s turns an iterator of complex64
// into a []complex64.
func GenerateComplex64s(ch <-chan complex64) []complex64 {
	s := make([]complex64, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateComplex128s turns an iterator of complex128
// into a []complex128.
func GenerateComplex128s(ch <-chan complex128) []complex128 {
	s := make([]complex128, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateUints turns an iterator of uint
// into a []uint.
func GenerateUints(ch <-chan uint) []uint {
	s := make([]uint, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateInts turns an iterator of int
// into a []int.
func GenerateInts(ch <-chan int) []int {
	s := make([]int, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GeneratePtrs turns an iterator of uintptr
// into a []uintptr.
func GeneratePtrs(ch <-chan uintptr) []uintptr {
	s := make([]uintptr, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateBools turns an iterator of bool
// into a []bool.
func GenerateBools(ch <-chan bool) []bool {
	s := make([]bool, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

// GenerateAny turns an iterator of interface{}
// into a []interface{}.
func GenerateAny(ch <-chan interface{}) []interface{} {
	s := make([]interface{}, 0, 32)
	for v := range ch {
		s = append(s, v)
	}
	return s
}

