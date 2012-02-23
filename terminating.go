/*
This is free software. It comes without any warranty, to the extent permitted
by applicable law. You can redistribute it and/or modify it under the terms of
the Do What The Fuck You Want To Public License, Version 2, as published by
Sam Hocevar. See COPYING for more details.
*/

package itertools

// IMPORTANT: terminating.go is generated from terminating.got. Change the latter.


type StringPredicate func(v string) bool

type BytePredicate func(v uint8) bool

type Uint16Predicate func(v uint16) bool

type Uint32Predicate func(v uint32) bool

type Uint64Predicate func(v uint64) bool

type Int8Predicate func(v int8) bool

type Int16Predicate func(v int16) bool

type Int32Predicate func(v int32) bool

type Int64Predicate func(v int64) bool

type Float32Predicate func(v float32) bool

type Float64Predicate func(v float64) bool

type Complex64Predicate func(v complex64) bool

type Complex128Predicate func(v complex128) bool

type UintPredicate func(v uint) bool

type IntPredicate func(v int) bool

type PtrPredicate func(v uintptr) bool

type BoolPredicate func(v bool) bool

type Predicate func(v interface{}) bool



// AccumulateUint8 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateUint8(it <-chan uint8) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateUint16 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateUint16(it <-chan uint16) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateUint32 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateUint32(it <-chan uint32) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateUint64 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateUint64(it <-chan uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateInt8 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateInt8(it <-chan int8) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateInt16 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateInt16(it <-chan int16) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateInt32 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateInt32(it <-chan int32) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateInt64 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateInt64(it <-chan int64) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateFloat32 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateFloat32(it <-chan float32) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateFloat64 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateFloat64(it <-chan float64) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateComplex64 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateComplex64(it <-chan complex64) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateComplex128 returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateComplex128(it <-chan complex128) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// AccumulateUint returns an iterator that yields accumulated sums of its
// input iterator's values.
func AccumulateUint(it <-chan uint) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// Accumulate returns an iterator that yields accumulated sums of its
// input iterator's values.
func Accumulate(it <-chan int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		sum := <-it
		ch <- sum
		for i := range it {
			sum += i
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

// ChainString returns an iterator that yields values from its arguments in
// order.
func ChainString(p ...<-chan string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainByte returns an iterator that yields values from its arguments in
// order.
func ChainByte(p ...<-chan uint8) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainUint16 returns an iterator that yields values from its arguments in
// order.
func ChainUint16(p ...<-chan uint16) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainUint32 returns an iterator that yields values from its arguments in
// order.
func ChainUint32(p ...<-chan uint32) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainUint64 returns an iterator that yields values from its arguments in
// order.
func ChainUint64(p ...<-chan uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainInt8 returns an iterator that yields values from its arguments in
// order.
func ChainInt8(p ...<-chan int8) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainInt16 returns an iterator that yields values from its arguments in
// order.
func ChainInt16(p ...<-chan int16) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainInt32 returns an iterator that yields values from its arguments in
// order.
func ChainInt32(p ...<-chan int32) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainInt64 returns an iterator that yields values from its arguments in
// order.
func ChainInt64(p ...<-chan int64) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainFloat32 returns an iterator that yields values from its arguments in
// order.
func ChainFloat32(p ...<-chan float32) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainFloat64 returns an iterator that yields values from its arguments in
// order.
func ChainFloat64(p ...<-chan float64) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainComplex64 returns an iterator that yields values from its arguments in
// order.
func ChainComplex64(p ...<-chan complex64) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainComplex128 returns an iterator that yields values from its arguments in
// order.
func ChainComplex128(p ...<-chan complex128) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainUint returns an iterator that yields values from its arguments in
// order.
func ChainUint(p ...<-chan uint) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainInt returns an iterator that yields values from its arguments in
// order.
func ChainInt(p ...<-chan int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainPtr returns an iterator that yields values from its arguments in
// order.
func ChainPtr(p ...<-chan uintptr) <-chan uintptr {
	ch := make(chan uintptr, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// ChainBool returns an iterator that yields values from its arguments in
// order.
func ChainBool(p ...<-chan bool) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// Chain returns an iterator that yields values from its arguments in
// order.
func Chain(p ...<-chan interface{}) <-chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		for _, s := range p {
			for i := range s {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// CompressString returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressString(p <-chan string, s <-chan bool) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressByte returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressByte(p <-chan uint8, s <-chan bool) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressUint16 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressUint16(p <-chan uint16, s <-chan bool) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressUint32 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressUint32(p <-chan uint32, s <-chan bool) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressUint64 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressUint64(p <-chan uint64, s <-chan bool) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressInt8 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressInt8(p <-chan int8, s <-chan bool) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressInt16 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressInt16(p <-chan int16, s <-chan bool) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressInt32 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressInt32(p <-chan int32, s <-chan bool) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressInt64 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressInt64(p <-chan int64, s <-chan bool) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressFloat32 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressFloat32(p <-chan float32, s <-chan bool) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressFloat64 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressFloat64(p <-chan float64, s <-chan bool) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressComplex64 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressComplex64(p <-chan complex64, s <-chan bool) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressComplex128 returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressComplex128(p <-chan complex128, s <-chan bool) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressUint returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressUint(p <-chan uint, s <-chan bool) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressInt returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressInt(p <-chan int, s <-chan bool) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressPtr returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressPtr(p <-chan uintptr, s <-chan bool) <-chan uintptr {
	ch := make(chan uintptr, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// CompressBool returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func CompressBool(p <-chan bool, s <-chan bool) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// Compress returns an iterator that yields each item from its first
// argument for which the corresponding item in the second argument is true.
func Compress(p <-chan interface{}, s <-chan bool) <-chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		for v := range p {
			b, ok := <-s
			if !ok {
				break
			}
			if b {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// DropString consumes values from s while pred(<-s) returns true.
func DropString(pred StringPredicate, s <-chan string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropByte consumes values from s while pred(<-s) returns true.
func DropByte(pred BytePredicate, s <-chan uint8) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropUint16 consumes values from s while pred(<-s) returns true.
func DropUint16(pred Uint16Predicate, s <-chan uint16) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropUint32 consumes values from s while pred(<-s) returns true.
func DropUint32(pred Uint32Predicate, s <-chan uint32) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropUint64 consumes values from s while pred(<-s) returns true.
func DropUint64(pred Uint64Predicate, s <-chan uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropInt8 consumes values from s while pred(<-s) returns true.
func DropInt8(pred Int8Predicate, s <-chan int8) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropInt16 consumes values from s while pred(<-s) returns true.
func DropInt16(pred Int16Predicate, s <-chan int16) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropInt32 consumes values from s while pred(<-s) returns true.
func DropInt32(pred Int32Predicate, s <-chan int32) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropInt64 consumes values from s while pred(<-s) returns true.
func DropInt64(pred Int64Predicate, s <-chan int64) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropFloat32 consumes values from s while pred(<-s) returns true.
func DropFloat32(pred Float32Predicate, s <-chan float32) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropFloat64 consumes values from s while pred(<-s) returns true.
func DropFloat64(pred Float64Predicate, s <-chan float64) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropComplex64 consumes values from s while pred(<-s) returns true.
func DropComplex64(pred Complex64Predicate, s <-chan complex64) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropComplex128 consumes values from s while pred(<-s) returns true.
func DropComplex128(pred Complex128Predicate, s <-chan complex128) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropUint consumes values from s while pred(<-s) returns true.
func DropUint(pred UintPredicate, s <-chan uint) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropInt consumes values from s while pred(<-s) returns true.
func DropInt(pred IntPredicate, s <-chan int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropPtr consumes values from s while pred(<-s) returns true.
func DropPtr(pred PtrPredicate, s <-chan uintptr) <-chan uintptr {
	ch := make(chan uintptr, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// DropBool consumes values from s while pred(<-s) returns true.
func DropBool(pred BoolPredicate, s <-chan bool) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// Drop consumes values from s while pred(<-s) returns true.
func Drop(pred Predicate, s <-chan interface{}) <-chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		for {
			if v, ok := <-s; !ok {
				break
			} else if !pred(v) {
				ch <- v
				for i := range s {
					ch <- i
				}
				break
			}
		}
		close(ch)
	}()
	return ch
}

// FilterString returns an iterator over values from s for which pred(v)
// returns true.
func FilterString(pred StringPredicate, s <-chan string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterByte returns an iterator over values from s for which pred(v)
// returns true.
func FilterByte(pred BytePredicate, s <-chan uint8) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterUint16 returns an iterator over values from s for which pred(v)
// returns true.
func FilterUint16(pred Uint16Predicate, s <-chan uint16) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterUint32 returns an iterator over values from s for which pred(v)
// returns true.
func FilterUint32(pred Uint32Predicate, s <-chan uint32) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterUint64 returns an iterator over values from s for which pred(v)
// returns true.
func FilterUint64(pred Uint64Predicate, s <-chan uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterInt8 returns an iterator over values from s for which pred(v)
// returns true.
func FilterInt8(pred Int8Predicate, s <-chan int8) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterInt16 returns an iterator over values from s for which pred(v)
// returns true.
func FilterInt16(pred Int16Predicate, s <-chan int16) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterInt32 returns an iterator over values from s for which pred(v)
// returns true.
func FilterInt32(pred Int32Predicate, s <-chan int32) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterInt64 returns an iterator over values from s for which pred(v)
// returns true.
func FilterInt64(pred Int64Predicate, s <-chan int64) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterFloat32 returns an iterator over values from s for which pred(v)
// returns true.
func FilterFloat32(pred Float32Predicate, s <-chan float32) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterFloat64 returns an iterator over values from s for which pred(v)
// returns true.
func FilterFloat64(pred Float64Predicate, s <-chan float64) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterComplex64 returns an iterator over values from s for which pred(v)
// returns true.
func FilterComplex64(pred Complex64Predicate, s <-chan complex64) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterComplex128 returns an iterator over values from s for which pred(v)
// returns true.
func FilterComplex128(pred Complex128Predicate, s <-chan complex128) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterUint returns an iterator over values from s for which pred(v)
// returns true.
func FilterUint(pred UintPredicate, s <-chan uint) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterInt returns an iterator over values from s for which pred(v)
// returns true.
func FilterInt(pred IntPredicate, s <-chan int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterPtr returns an iterator over values from s for which pred(v)
// returns true.
func FilterPtr(pred PtrPredicate, s <-chan uintptr) <-chan uintptr {
	ch := make(chan uintptr, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// FilterBool returns an iterator over values from s for which pred(v)
// returns true.
func FilterBool(pred BoolPredicate, s <-chan bool) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// Filter returns an iterator over values from s for which pred(v)
// returns true.
func Filter(pred Predicate, s <-chan interface{}) <-chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		for v := range s {
			if pred(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// MapString returns an iterator over the values of f(v) for each v in s.
func MapString(f func(v string) string, s <-chan string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapByte returns an iterator over the values of f(v) for each v in s.
func MapByte(f func(v uint8) uint8, s <-chan uint8) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapUint16 returns an iterator over the values of f(v) for each v in s.
func MapUint16(f func(v uint16) uint16, s <-chan uint16) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapUint32 returns an iterator over the values of f(v) for each v in s.
func MapUint32(f func(v uint32) uint32, s <-chan uint32) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapUint64 returns an iterator over the values of f(v) for each v in s.
func MapUint64(f func(v uint64) uint64, s <-chan uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapInt8 returns an iterator over the values of f(v) for each v in s.
func MapInt8(f func(v int8) int8, s <-chan int8) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapInt16 returns an iterator over the values of f(v) for each v in s.
func MapInt16(f func(v int16) int16, s <-chan int16) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapInt32 returns an iterator over the values of f(v) for each v in s.
func MapInt32(f func(v int32) int32, s <-chan int32) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapInt64 returns an iterator over the values of f(v) for each v in s.
func MapInt64(f func(v int64) int64, s <-chan int64) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapFloat32 returns an iterator over the values of f(v) for each v in s.
func MapFloat32(f func(v float32) float32, s <-chan float32) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapFloat64 returns an iterator over the values of f(v) for each v in s.
func MapFloat64(f func(v float64) float64, s <-chan float64) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapComplex64 returns an iterator over the values of f(v) for each v in s.
func MapComplex64(f func(v complex64) complex64, s <-chan complex64) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapComplex128 returns an iterator over the values of f(v) for each v in s.
func MapComplex128(f func(v complex128) complex128, s <-chan complex128) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapUint returns an iterator over the values of f(v) for each v in s.
func MapUint(f func(v uint) uint, s <-chan uint) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapInt returns an iterator over the values of f(v) for each v in s.
func MapInt(f func(v int) int, s <-chan int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapPtr returns an iterator over the values of f(v) for each v in s.
func MapPtr(f func(v uintptr) uintptr, s <-chan uintptr) <-chan uintptr {
	ch := make(chan uintptr, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// MapBool returns an iterator over the values of f(v) for each v in s.
func MapBool(f func(v bool) bool, s <-chan bool) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// Map returns an iterator over the values of f(v) for each v in s.
func Map(f func(v interface{}) interface{}, s <-chan interface{}) <-chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		for v := range s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// TakeString returns an iterator over the values of s until pred(s) returns
// false.
func TakeString(pred StringPredicate, s <-chan string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeByte returns an iterator over the values of s until pred(s) returns
// false.
func TakeByte(pred BytePredicate, s <-chan uint8) <-chan uint8 {
	ch := make(chan uint8, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeUint16 returns an iterator over the values of s until pred(s) returns
// false.
func TakeUint16(pred Uint16Predicate, s <-chan uint16) <-chan uint16 {
	ch := make(chan uint16, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeUint32 returns an iterator over the values of s until pred(s) returns
// false.
func TakeUint32(pred Uint32Predicate, s <-chan uint32) <-chan uint32 {
	ch := make(chan uint32, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeUint64 returns an iterator over the values of s until pred(s) returns
// false.
func TakeUint64(pred Uint64Predicate, s <-chan uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeInt8 returns an iterator over the values of s until pred(s) returns
// false.
func TakeInt8(pred Int8Predicate, s <-chan int8) <-chan int8 {
	ch := make(chan int8, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeInt16 returns an iterator over the values of s until pred(s) returns
// false.
func TakeInt16(pred Int16Predicate, s <-chan int16) <-chan int16 {
	ch := make(chan int16, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeInt32 returns an iterator over the values of s until pred(s) returns
// false.
func TakeInt32(pred Int32Predicate, s <-chan int32) <-chan int32 {
	ch := make(chan int32, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeInt64 returns an iterator over the values of s until pred(s) returns
// false.
func TakeInt64(pred Int64Predicate, s <-chan int64) <-chan int64 {
	ch := make(chan int64, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeFloat32 returns an iterator over the values of s until pred(s) returns
// false.
func TakeFloat32(pred Float32Predicate, s <-chan float32) <-chan float32 {
	ch := make(chan float32, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeFloat64 returns an iterator over the values of s until pred(s) returns
// false.
func TakeFloat64(pred Float64Predicate, s <-chan float64) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeComplex64 returns an iterator over the values of s until pred(s) returns
// false.
func TakeComplex64(pred Complex64Predicate, s <-chan complex64) <-chan complex64 {
	ch := make(chan complex64, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeComplex128 returns an iterator over the values of s until pred(s) returns
// false.
func TakeComplex128(pred Complex128Predicate, s <-chan complex128) <-chan complex128 {
	ch := make(chan complex128, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeUint returns an iterator over the values of s until pred(s) returns
// false.
func TakeUint(pred UintPredicate, s <-chan uint) <-chan uint {
	ch := make(chan uint, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeInt returns an iterator over the values of s until pred(s) returns
// false.
func TakeInt(pred IntPredicate, s <-chan int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakePtr returns an iterator over the values of s until pred(s) returns
// false.
func TakePtr(pred PtrPredicate, s <-chan uintptr) <-chan uintptr {
	ch := make(chan uintptr, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TakeBool returns an iterator over the values of s until pred(s) returns
// false.
func TakeBool(pred BoolPredicate, s <-chan bool) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// Take returns an iterator over the values of s until pred(s) returns
// false.
func Take(pred Predicate, s <-chan interface{}) <-chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		for v := range s {
			if !pred(v) {
				close(ch)
				return
			}
			ch <- v
		}
	}()
	return ch
}

// TeeString returns n iterators over the values of s, each of which receives
// each item of s.
func TeeString(s <-chan string, n int) []<-chan string {
	chs := make([]chan string, n)
	for i := range chs {
		chs[i] = make(chan string, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan string, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeByte returns n iterators over the values of s, each of which receives
// each item of s.
func TeeByte(s <-chan uint8, n int) []<-chan uint8 {
	chs := make([]chan uint8, n)
	for i := range chs {
		chs[i] = make(chan uint8, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan uint8, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeUint16 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeUint16(s <-chan uint16, n int) []<-chan uint16 {
	chs := make([]chan uint16, n)
	for i := range chs {
		chs[i] = make(chan uint16, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan uint16, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeUint32 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeUint32(s <-chan uint32, n int) []<-chan uint32 {
	chs := make([]chan uint32, n)
	for i := range chs {
		chs[i] = make(chan uint32, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan uint32, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeUint64 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeUint64(s <-chan uint64, n int) []<-chan uint64 {
	chs := make([]chan uint64, n)
	for i := range chs {
		chs[i] = make(chan uint64, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan uint64, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeInt8 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeInt8(s <-chan int8, n int) []<-chan int8 {
	chs := make([]chan int8, n)
	for i := range chs {
		chs[i] = make(chan int8, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan int8, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeInt16 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeInt16(s <-chan int16, n int) []<-chan int16 {
	chs := make([]chan int16, n)
	for i := range chs {
		chs[i] = make(chan int16, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan int16, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeInt32 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeInt32(s <-chan int32, n int) []<-chan int32 {
	chs := make([]chan int32, n)
	for i := range chs {
		chs[i] = make(chan int32, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan int32, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeInt64 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeInt64(s <-chan int64, n int) []<-chan int64 {
	chs := make([]chan int64, n)
	for i := range chs {
		chs[i] = make(chan int64, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan int64, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeFloat32 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeFloat32(s <-chan float32, n int) []<-chan float32 {
	chs := make([]chan float32, n)
	for i := range chs {
		chs[i] = make(chan float32, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan float32, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeFloat64 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeFloat64(s <-chan float64, n int) []<-chan float64 {
	chs := make([]chan float64, n)
	for i := range chs {
		chs[i] = make(chan float64, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan float64, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeComplex64 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeComplex64(s <-chan complex64, n int) []<-chan complex64 {
	chs := make([]chan complex64, n)
	for i := range chs {
		chs[i] = make(chan complex64, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan complex64, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeComplex128 returns n iterators over the values of s, each of which receives
// each item of s.
func TeeComplex128(s <-chan complex128, n int) []<-chan complex128 {
	chs := make([]chan complex128, n)
	for i := range chs {
		chs[i] = make(chan complex128, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan complex128, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeUint returns n iterators over the values of s, each of which receives
// each item of s.
func TeeUint(s <-chan uint, n int) []<-chan uint {
	chs := make([]chan uint, n)
	for i := range chs {
		chs[i] = make(chan uint, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan uint, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeInt returns n iterators over the values of s, each of which receives
// each item of s.
func TeeInt(s <-chan int, n int) []<-chan int {
	chs := make([]chan int, n)
	for i := range chs {
		chs[i] = make(chan int, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan int, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeePtr returns n iterators over the values of s, each of which receives
// each item of s.
func TeePtr(s <-chan uintptr, n int) []<-chan uintptr {
	chs := make([]chan uintptr, n)
	for i := range chs {
		chs[i] = make(chan uintptr, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan uintptr, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// TeeBool returns n iterators over the values of s, each of which receives
// each item of s.
func TeeBool(s <-chan bool, n int) []<-chan bool {
	chs := make([]chan bool, n)
	for i := range chs {
		chs[i] = make(chan bool, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan bool, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// Tee returns n iterators over the values of s, each of which receives
// each item of s.
func Tee(s <-chan interface{}, n int) []<-chan interface{} {
	chs := make([]chan interface{}, n)
	for i := range chs {
		chs[i] = make(chan interface{}, 8)
	}
	go func() {
		for v := range s {
			for _, ch := range chs {
				ch <- v
			}
		}
		for _, ch := range chs {
			close(ch)
		}
	}()
	ret := make([]<-chan interface{}, n)
	for i := range ret {
		ret[i] = chs[i]
	}
	return ret
}

// ZipString returns an iterator over each nth value of each of its
// arguments.
func ZipString(p ...<-chan string) <-chan []string {
	ch := make(chan []string, 1)
	go func() {
		var ok bool
		for {
			m := make([]string, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipByte returns an iterator over each nth value of each of its
// arguments.
func ZipByte(p ...<-chan uint8) <-chan []uint8 {
	ch := make(chan []uint8, 1)
	go func() {
		var ok bool
		for {
			m := make([]uint8, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipUint16 returns an iterator over each nth value of each of its
// arguments.
func ZipUint16(p ...<-chan uint16) <-chan []uint16 {
	ch := make(chan []uint16, 1)
	go func() {
		var ok bool
		for {
			m := make([]uint16, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipUint32 returns an iterator over each nth value of each of its
// arguments.
func ZipUint32(p ...<-chan uint32) <-chan []uint32 {
	ch := make(chan []uint32, 1)
	go func() {
		var ok bool
		for {
			m := make([]uint32, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipUint64 returns an iterator over each nth value of each of its
// arguments.
func ZipUint64(p ...<-chan uint64) <-chan []uint64 {
	ch := make(chan []uint64, 1)
	go func() {
		var ok bool
		for {
			m := make([]uint64, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipInt8 returns an iterator over each nth value of each of its
// arguments.
func ZipInt8(p ...<-chan int8) <-chan []int8 {
	ch := make(chan []int8, 1)
	go func() {
		var ok bool
		for {
			m := make([]int8, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipInt16 returns an iterator over each nth value of each of its
// arguments.
func ZipInt16(p ...<-chan int16) <-chan []int16 {
	ch := make(chan []int16, 1)
	go func() {
		var ok bool
		for {
			m := make([]int16, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipInt32 returns an iterator over each nth value of each of its
// arguments.
func ZipInt32(p ...<-chan int32) <-chan []int32 {
	ch := make(chan []int32, 1)
	go func() {
		var ok bool
		for {
			m := make([]int32, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipInt64 returns an iterator over each nth value of each of its
// arguments.
func ZipInt64(p ...<-chan int64) <-chan []int64 {
	ch := make(chan []int64, 1)
	go func() {
		var ok bool
		for {
			m := make([]int64, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipFloat32 returns an iterator over each nth value of each of its
// arguments.
func ZipFloat32(p ...<-chan float32) <-chan []float32 {
	ch := make(chan []float32, 1)
	go func() {
		var ok bool
		for {
			m := make([]float32, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipFloat64 returns an iterator over each nth value of each of its
// arguments.
func ZipFloat64(p ...<-chan float64) <-chan []float64 {
	ch := make(chan []float64, 1)
	go func() {
		var ok bool
		for {
			m := make([]float64, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipComplex64 returns an iterator over each nth value of each of its
// arguments.
func ZipComplex64(p ...<-chan complex64) <-chan []complex64 {
	ch := make(chan []complex64, 1)
	go func() {
		var ok bool
		for {
			m := make([]complex64, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipComplex128 returns an iterator over each nth value of each of its
// arguments.
func ZipComplex128(p ...<-chan complex128) <-chan []complex128 {
	ch := make(chan []complex128, 1)
	go func() {
		var ok bool
		for {
			m := make([]complex128, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipUint returns an iterator over each nth value of each of its
// arguments.
func ZipUint(p ...<-chan uint) <-chan []uint {
	ch := make(chan []uint, 1)
	go func() {
		var ok bool
		for {
			m := make([]uint, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipInt returns an iterator over each nth value of each of its
// arguments.
func ZipInt(p ...<-chan int) <-chan []int {
	ch := make(chan []int, 1)
	go func() {
		var ok bool
		for {
			m := make([]int, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipPtr returns an iterator over each nth value of each of its
// arguments.
func ZipPtr(p ...<-chan uintptr) <-chan []uintptr {
	ch := make(chan []uintptr, 1)
	go func() {
		var ok bool
		for {
			m := make([]uintptr, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// ZipBool returns an iterator over each nth value of each of its
// arguments.
func ZipBool(p ...<-chan bool) <-chan []bool {
	ch := make(chan []bool, 1)
	go func() {
		var ok bool
		for {
			m := make([]bool, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

// Zip returns an iterator over each nth value of each of its
// arguments.
func Zip(p ...<-chan interface{}) <-chan []interface{} {
	ch := make(chan []interface{}, 1)
	go func() {
		var ok bool
		for {
			m := make([]interface{}, len(p))
			for i, s := range p {
				m[i], ok = <-s
				if !ok {
					close(ch)
					return
				}
			}
			ch <- m
		}
	}()
	return ch
}

