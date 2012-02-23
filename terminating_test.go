package itertools

import (
	"testing"
)

func TestAccumulate(t *testing.T) {
	accum := Accumulate(Range(0, 1, 10))
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		if x, ok := <-accum; !ok {
			t.Fatal("accumulate ended early")
		} else if sum != x {
			t.Fatal("accumulate incorrect: got", x, "expected", sum)
		}
	}
	select {
	case _, bad := <-accum:
		if bad {
			t.Fatal("accumulate too large")
		}
	default:
	}
}

func TestChain(t *testing.T) {
	answer := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	chain := ChainInt32(AdaptString("abc"), AdaptString("def"))
	for _, r := range answer {
		if x, ok := <-chain; !ok {
			t.Fatal("chain too short")
		} else if x != r {
			t.Fatal("chain incorrect")
		}
	}
	select {
	case _, bad := <-chain:
		if bad {
			t.Fatal("chain too long")
		}
	default:
	}
}

func TestCompress(t *testing.T) {
	bools := CycleBool([]bool{true, false, false, true, false})
	comp := CompressInt(Range(0, 1, 100), bools)
	for i := range comp {
		if m := i % 5; m != 0 && m != 3 {
			t.Fatal("compress incorrect")
		}
	}
}

func TestDrop(t *testing.T) {
	answer := []int{6, 5, 2}
	drop := DropInt(func(x int) bool { return x < 5 }, AdaptInts([]int{1, 4, 6, 5, 2}))
	for _, v := range answer {
		if x, ok := <-drop; !ok {
			t.Fatal("drop short")
		} else if v != x {
			t.Fatal("drop incorrect: got", x, "expected", v)
		}
	}
}

func TestFilter(t *testing.T) {
	filter := FilterInt(func(x int) bool { return x&1 == 0 }, Range(0, 1, 1000))
	for i := range filter {
		if i&1 != 0 {
			t.Fatal("filter incorrect")
		}
	}
}

func TestMap(t *testing.T) {
	mapper := MapInt(func(x int) int { return x << 1 }, Range(0, 1, 1000))
	for i := 0; i < 1000; i++ {
		if x, ok := <-mapper; !ok {
			t.Fatal("map ended early")
		} else if i != x>>1 {
			t.Fatal("map incorrect")
		}
	}
	select {
	case _, bad := <-mapper:
		if bad {
			t.Fatal("map too long")
		}
	default:
	}
}

func TestTake(t *testing.T) {
	answer := []int{1, 4}
	take := TakeInt(func(x int) bool { return x < 5 }, AdaptInts([]int{1, 4, 6, 4, 1}))
	for _, v := range answer {
		if v != <-take {
			t.Fatal("take incorrect")
		}
	}
}

func TestTee(t *testing.T) {
	tees := TeeInt(Range(0, 1, 1000), 5) // such a *shades* tees YEEAAAAH
	for i := 0; i < 1000; i++ {
		for _, ch := range tees {
			if x, ok := <-ch; !ok {
				t.Fatal("tee ended early")
			} else if x != i {
				t.Fatal("tee incorrect")
			}
		}
	}
	// I could test that they're not too long, or I could be lazy.
}

func TestZip(t *testing.T) {
	zip := ZipInt(Range(0, 1, 100), Range(1, 1, 101))
	for i := range zip {
		if len(i) != 2 {
			t.Fatal("incorrect number of zipped values")
		}
		if i[0] != i[1]-1 {
			t.Fatal("zip incorrect")
		}
	}
}
