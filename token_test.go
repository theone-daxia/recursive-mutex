package recursive_mutex

import (
	"testing"
)

func TestTokenRecursiveMutex(t *testing.T)  {
	var mux TokenRecursiveMutex

	mux.Lock(1)
	if mux.token != 1 {
		t.Errorf("token error! expect 1, actual %d\n", mux.token)
	}
	if mux.recursion != 1 {
		t.Errorf("recursion error! expect 1, actual %d\n", mux.recursion)
	}

	mux.Lock(1)
	if mux.recursion != 2 {
		t.Errorf("recursion error! expect 2, actual %d\n", mux.recursion)
	}

	t.Log("123")
	mux.Unlock(1)
	if mux.recursion != 1 {
		t.Errorf("recursion error! expect 1, actual %d\n", mux.recursion)
	}

	mux.Unlock(1)
	if mux.recursion != 0 {
		t.Errorf("recursion error! expect 0, actual %d\n", mux.recursion)
	}
	if mux.token != -1 {
		t.Errorf("token error! expect -1, actual %d\n", mux.token)
	}
}