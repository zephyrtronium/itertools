package itertools

import (
	"math/rand"
	"testing"
)

func TestAdaptString(t *testing.T) {
	s := make([]rune, 64)
	for i := range s {
		s[i] = rune(rand.Intn(256))
	}
	s2 := make([]rune, 64)
	copy(s2, s)
	ch := AdaptString(string(s))
	for _, c := range s2 {
		if c2, ok := <-ch; !ok {
			t.Fatal("adapter short")
		} else if c != c2 {
			t.Fatal("adapter miscommunicating: got", c2, "expected", c)
		}
	}
	select {
	case _, bad := <-ch:
		if bad {
			t.Fatal("adapter long")
		}
	default:
	}
}

type AdaptTester struct {
	A int
	B int
}

func TestAdaptAny(t *testing.T) {
	s := make([]interface{}, 64)
	for i := range s {
		s[i] = AdaptTester{rand.Int(), rand.Int()}
	}
	s2 := make([]interface{}, 64)
	copy(s2, s)
	ch := AdaptAny(s)
	for _, v := range s2 {
		if v2, ok := <-ch; !ok {
			t.Fatal("adapter short")
		} else if v != v2 {
			t.Fatal("adapter miscommunicating: got", v2, "expected", v)
		}
	}
	select {
	case _, bad := <-ch:
		if bad {
			t.Fatal("adapter long")
		}
	default:
	}
}

func TestGenerateString(t *testing.T) {
	cl := make([]rune, 64)
	ch := make(chan rune, 64)
	for i := range cl {
		cl[i] = rune(rand.Intn(256))
		ch <- cl[i]
	}
	close(ch)
	s := GenerateString(ch)
	if len(s) != len(string(cl)) {
		t.Fatal("generator has bad length: got", len(s), "expected", len(cl))
	}
	if s != string(cl) {
		t.Fatal("generator failed")
	}
}

func TestGenerateAny(t *testing.T) {
	cl := make([]AdaptTester, 64)
	ch := make(chan interface{}, 64)
	for i := range cl {
		cl[i] = AdaptTester{rand.Int(), rand.Int()}
		ch <- cl[i]
	}
	close(ch)
	s := GenerateAny(ch)
	if len(s) != len(cl) {
		t.Fatal("generator has bad length: got", len(s), "expected", len(cl))
	}
	for i := range s {
		if s[i] != cl[i] {
			t.Fatal("generator failed")
		}
	}
}
