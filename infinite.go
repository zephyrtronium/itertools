/*
This is free software. It comes without any warranty, to the extent permitted
by applicable law. You can redistribute it and/or modify it under the terms of
the Do What The Fuck You Want To Public License, Version 2, as published by
Sam Hocevar. See COPYING for more details.
*/

package itertools

// IMPORTANT: infinite.go is generated from infinite.got. Change the latter.

// Count creates an iterator of type int that starts at
// 0 and increments by 1.
func Count() <-chan int {
	ch := make(chan int, 1)
	go func() {
		var x int
		for ;; x++ {
			ch <- x
		}
	}()
	return ch
}

// CountStartStep creates an iterator of type int that starts at
// start and increments by step.
func CountStartStep(start, step int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		
		for ;; start += step {
			ch <- start
		}
	}()
	return ch
}

// Count64 creates an iterator of type uint64 that starts at
// 0 and increments by 1.
func Count64() <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		var x uint64
		for ;; x++ {
			ch <- x
		}
	}()
	return ch
}

// CountStartStep64 creates an iterator of type uint64 that starts at
// start and increments by step.
func CountStartStep64(start, step uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		
		for ;; start += step {
			ch <- start
		}
	}()
	return ch
}

// Range creates an iterator of type int that starts at start,
// increments by step, and stops when the value becomes greater than or equal
// to stop. If step is 0, return nil.
func Range(start, step, stop int) <-chan int {
	ch := make(chan int, 1)
	if step > 0 {
		go func() {
			for ; start < stop; start += step {
				ch <- start
			}
			close(ch)
		}()
	} else if step < 0 {
		go func() {
			for ; start > stop; start += step {
				ch <- start
			}
			close(ch)
		}()
	} else {
		return nil
	}
	return ch
}

// Range64 creates an iterator of type uint64 that starts at start,
// increments by step, and stops when the value becomes greater than or equal
// to stop. If step is 0, return nil.
func Range64(start, step, stop uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	if step > 0 {
		go func() {
			for ; start < stop; start += step {
				ch <- start
			}
			close(ch)
		}()
	} else if step < 0 {
		go func() {
			for ; start > stop; start += step {
				ch <- start
			}
			close(ch)
		}()
	} else {
		return nil
	}
	return ch
}

// CycleRune creates an infinitely cycling iterator over a string yielding its
// runes.
func CycleRune(s string) <-chan rune {
	ch := make(chan rune, 1)
	r := []rune(s)
	go func() {
		i := 0
		for {
			ch <- r[i]
			if i == len(r) - 1 {
				i = 0
			} else {
				i++
			}
		}
	}()
	return ch
}

// CycleString creates an infinitely cycling iterator over a slice of
// strings. See Cycle for more information.
func CycleString(s []string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleByte creates an infinitely cycling iterator over a slice of
// uint8s. See Cycle for more information.
func CycleByte(s []uint8) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleUint16 creates an infinitely cycling iterator over a slice of
// uint16s. See Cycle for more information.
func CycleUint16(s []uint16) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleUint32 creates an infinitely cycling iterator over a slice of
// uint32s. See Cycle for more information.
func CycleUint32(s []uint32) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleUint64 creates an infinitely cycling iterator over a slice of
// uint64s. See Cycle for more information.
func CycleUint64(s []uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleInt8 creates an infinitely cycling iterator over a slice of
// int8s. See Cycle for more information.
func CycleInt8(s []int8) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleInt16 creates an infinitely cycling iterator over a slice of
// int16s. See Cycle for more information.
func CycleInt16(s []int16) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleInt32 creates an infinitely cycling iterator over a slice of
// int32s. See Cycle for more information.
func CycleInt32(s []int32) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleInt64 creates an infinitely cycling iterator over a slice of
// int64s. See Cycle for more information.
func CycleInt64(s []int64) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleFloat32 creates an infinitely cycling iterator over a slice of
// float32s. See Cycle for more information.
func CycleFloat32(s []float32) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleFloat64 creates an infinitely cycling iterator over a slice of
// float64s. See Cycle for more information.
func CycleFloat64(s []float64) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleComplex64 creates an infinitely cycling iterator over a slice of
// complex64s. See Cycle for more information.
func CycleComplex64(s []complex64) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleComplex128 creates an infinitely cycling iterator over a slice of
// complex128s. See Cycle for more information.
func CycleComplex128(s []complex128) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleUint creates an infinitely cycling iterator over a slice of
// uints. See Cycle for more information.
func CycleUint(s []uint) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleInt creates an infinitely cycling iterator over a slice of
// ints. See Cycle for more information.
func CycleInt(s []int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CyclePtr creates an infinitely cycling iterator over a slice of
// uintptrs. See Cycle for more information.
func CyclePtr(s []uintptr) <-chan uintptr {
	ch := make(chan uintptr, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// CycleBool creates an infinitely cycling iterator over a slice of
// bools. See Cycle for more information.
func CycleBool(s []bool) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// Cycle creates an infinitely cycling iterator over a slice of
// interface{}s. The argument is
// not copied, so it never will be garbage-collected, and modifications to it
// will be apparent in the iterator.
func Cycle(s []interface{}) <-chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		for i := 0;; i++ {
			ch <- s[i%len(s)]
		}
	}()
	return ch
}

// RepeatString returns an iterator that yields an infinitude of its
// argument.
func RepeatString(v string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatByte returns an iterator that yields an infinitude of its
// argument.
func RepeatByte(v uint8) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatUint16 returns an iterator that yields an infinitude of its
// argument.
func RepeatUint16(v uint16) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatUint32 returns an iterator that yields an infinitude of its
// argument.
func RepeatUint32(v uint32) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatUint64 returns an iterator that yields an infinitude of its
// argument.
func RepeatUint64(v uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatInt8 returns an iterator that yields an infinitude of its
// argument.
func RepeatInt8(v int8) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatInt16 returns an iterator that yields an infinitude of its
// argument.
func RepeatInt16(v int16) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatInt32 returns an iterator that yields an infinitude of its
// argument.
func RepeatInt32(v int32) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatInt64 returns an iterator that yields an infinitude of its
// argument.
func RepeatInt64(v int64) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatFloat32 returns an iterator that yields an infinitude of its
// argument.
func RepeatFloat32(v float32) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatFloat64 returns an iterator that yields an infinitude of its
// argument.
func RepeatFloat64(v float64) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatComplex64 returns an iterator that yields an infinitude of its
// argument.
func RepeatComplex64(v complex64) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatComplex128 returns an iterator that yields an infinitude of its
// argument.
func RepeatComplex128(v complex128) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatUint returns an iterator that yields an infinitude of its
// argument.
func RepeatUint(v uint) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatInt returns an iterator that yields an infinitude of its
// argument.
func RepeatInt(v int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatPtr returns an iterator that yields an infinitude of its
// argument.
func RepeatPtr(v uintptr) <-chan uintptr {
	ch := make(chan uintptr, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatBool returns an iterator that yields an infinitude of its
// argument.
func RepeatBool(v bool) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// Repeat returns an iterator that yields an infinitude of its
// argument.
func Repeat(v interface{}) <-chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		for {
			ch <- v
		}
	}()
	return ch
}

// RepeatStringN returns an iterator that yields its argument n times.
func RepeatStringN(v string, n int) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatByteN returns an iterator that yields its argument n times.
func RepeatByteN(v uint8, n int) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatUint16N returns an iterator that yields its argument n times.
func RepeatUint16N(v uint16, n int) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatUint32N returns an iterator that yields its argument n times.
func RepeatUint32N(v uint32, n int) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatUint64N returns an iterator that yields its argument n times.
func RepeatUint64N(v uint64, n int) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatInt8N returns an iterator that yields its argument n times.
func RepeatInt8N(v int8, n int) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatInt16N returns an iterator that yields its argument n times.
func RepeatInt16N(v int16, n int) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatInt32N returns an iterator that yields its argument n times.
func RepeatInt32N(v int32, n int) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatInt64N returns an iterator that yields its argument n times.
func RepeatInt64N(v int64, n int) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatFloat32N returns an iterator that yields its argument n times.
func RepeatFloat32N(v float32, n int) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatFloat64N returns an iterator that yields its argument n times.
func RepeatFloat64N(v float64, n int) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatComplex64N returns an iterator that yields its argument n times.
func RepeatComplex64N(v complex64, n int) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatComplex128N returns an iterator that yields its argument n times.
func RepeatComplex128N(v complex128, n int) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatUintN returns an iterator that yields its argument n times.
func RepeatUintN(v uint, n int) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatIntN returns an iterator that yields its argument n times.
func RepeatIntN(v int, n int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatPtrN returns an iterator that yields its argument n times.
func RepeatPtrN(v uintptr, n int) <-chan uintptr {
	ch := make(chan uintptr, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatBoolN returns an iterator that yields its argument n times.
func RepeatBoolN(v bool, n int) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

// RepeatN returns an iterator that yields its argument n times.
func RepeatN(v interface{}, n int) <-chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()
	return ch
}

