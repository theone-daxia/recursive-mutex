package recursive_mutex

import (
	"github.com/petermattis/goid"
	"testing"
)

func TestRecursiveMutex(t *testing.T)  {
	var mux RecursiveMutex
	gid := goid.Get()

	mux.Lock()
	if gid != mux.owner {
		t.Errorf("owner error! expect %d, actual %d\n", gid, mux.owner)
	}
	if mux.recursion != 1 {
		t.Errorf("recursion error! expect 1, actual %d\n", mux.recursion)
	}

	mux.Lock()
	if mux.recursion != 2 {
		t.Errorf("recursion error! expect 2, actual %d\n", mux.recursion)
	}

	t.Log("123")
	mux.Unlock()
	if mux.recursion != 1 {
		t.Errorf("recursion error! expect 1, actual %d\n", mux.recursion)
	}

	mux.Unlock()
	if mux.recursion != 0 {
		t.Errorf("recursion error! expect 0, actual %d\n", mux.recursion)
	}
	if mux.owner != -1 {
		t.Errorf("owner error! expect -1, actual %d\n", mux.owner)
	}
}
