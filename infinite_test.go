package itertools

import (
	"math/rand"
	"testing"
)

func TestCount(t *testing.T) {
	count := Count()
	for i := 0; i < 1000000; i++ {
		x, ok := <-count
		if !ok {
			t.Fatal("count stopped at", i)
		}
		if i != x {
			t.Fatal("count incorrect: got", x, "expected", i)
		}
	}
}

func TestCountStartStep(t *testing.T) {
	start := rand.Intn(200000) - 100000
	step := rand.Intn(100) + 1
	count := CountStartStep(start, step)
	for i := start; i < start+step*1000; i += step {
		if x, ok := <-count; !ok {
			t.Fatal("count stopped at", i)
		} else if i != x {
			t.Fatal("count incorrect: got", x, "expected", i)
		}
	}
}

func TestRange(t *testing.T) {
	start := rand.Intn(200000) - 100000
	step := rand.Intn(100) + 1
	stop := start + step*(rand.Intn(50)+1)
	rang := Range(start, step, stop)
	for i := start; i < stop; i += step {
		if x, ok := <-rang; !ok {
			t.Fatal("range stopped at", i)
		} else if i != x {
			t.Fatal("range incorrect: got", x, "expected", i)
		}
	}
	select {
	case _, bad := <-rang:
		if bad {
			t.Fatal("range too large")
		}
	default:
	}
}

func TestRangeNegativeStep(t *testing.T) {
	start := rand.Intn(200000) - 100000
	step := -rand.Intn(100) - 1
	stop := start + step*(rand.Intn(50)+1)
	rang := Range(start, step, stop)
	for i := start; i < stop; i += step {
		if x, ok := <-rang; !ok {
			t.Fatal("range stopped at", i)
		} else if i != x {
			t.Fatal("range incorrect: got", x, "expected", i)
		}
	}
	select {
	case _, bad := <-rang:
		if bad {
			t.Fatal("range too large")
		}
	default:
	}
}

func TestRangeZeroStep(t *testing.T) {
	if rang := Range(-5, 0, 5); rang != nil {
		t.Fatal("range with zero step not nil")
	}
}

func TestCycleRune(t *testing.T) {
	s := "I don't much enjoy writing tests. :c"
	runes := []rune(s)
	cycle := CycleRune(s)
	for i := 0; i < 100000; i++ {
		if r, ok := <-cycle; !ok {
			t.Fatal("cycle stopped at", i)
		} else if c := runes[i%len(runes)]; r != c {
			t.Fatal("cycle incorrect: got", r, "expected", c)
		}
	}
}

func TestCycle(t *testing.T) {
	s := make([]interface{}, 64)
	for i := range s {
		s[i] = i
	}
	cycle := Cycle(s)
	for i := 0; i < 100000; i++ {
		if r, ok := <-cycle; !ok {
			t.Fatal("cycle stopped at", i)
		} else if c := s[i%len(s)]; r.(int) != c {
			t.Fatal("cycle incorrect: got", r, "expected", c)
		}
	}
}

func TestRepeat(t *testing.T) {
	v := rand.Int()
	repeat := Repeat(v)
	for i := 0; i < 100000; i++ {
		if r, ok := <-repeat; !ok {
			t.Fatal("repeat stopped at", i)
		} else if r.(int) != v {
			t.Fatal("repeat incorrect: got", r, "expected", v)
		}
	}
}

func TestRepeatN(t *testing.T) {
	v := rand.Int()
	n := rand.Intn(100000) + 10
	repeat := RepeatN(v, n)
	for i := 0; i < n; i++ {
		if r, ok := <-repeat; !ok {
			t.Fatal("repeat stopped at", i)
		} else if r.(int) != v {
			t.Fatal("repeat incorrect: got", r, "expected", v)
		}
	}
	select {
	case _, bad := <-repeat:
		if bad {
			t.Fatal("repeated too many times")
		}
	default:
	}
}
