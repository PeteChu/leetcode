package lrucache_test

import (
	lrucache "_/Users/phuengton-webdev/Workspace/leetcode/go/146_LRU_Cache"
	"testing"
)

func TestLRUCache_Initiate(t *testing.T) {
	capacity := 4
	lru := lrucache.Constructor(4)

	if lru.Capacity() != capacity {
		t.Errorf("want %d, but got %d", capacity, lru.Capacity())
	}
}

func TestLRUCache_ZeroCapacity(t *testing.T) {
	lru := lrucache.Constructor(0)
	got := lru.Capacity()
	want := 0

	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}

func TestLRUCache_Get_NoValue(t *testing.T) {
	lru := lrucache.Constructor(5)

	got := lru.Get(6)
	want := -1

	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}

func TestLRUCache_Get_MovesAccessedToHead(t *testing.T) {
	lru := lrucache.Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)

	got := lru.Get(2) // access key 2 to move it to the head
	want := lru.Head().Value()

	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}

func TestLRUCache_Get_HasValue(t *testing.T) {
	lru := lrucache.Constructor(5)

	lru.Put(1, 1)
	got := lru.Get(1)
	want := 1

	if got != want {
		t.Errorf("want %d, but got %d", want, got)
	}
}

func TestLRUCache_Put_EvictionOrderAndTailUpdate(t *testing.T) {
	lru := lrucache.Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	lru.Put(5, 5)

	capacity := lru.Capacity()
	size := lru.Size()

	if size != capacity {
		t.Errorf("want size %d, but got %d", capacity, size)
	}

	shouldBeEvicted := lru.Get(2)
	want := -1

	if shouldBeEvicted != want {
		t.Errorf("want %d, but got %d", want, shouldBeEvicted)
	}
}

func TestLRUCache_Put_UpdateExistingKey(t *testing.T) {
	lru := lrucache.Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(2, 3) // update value of key 2

	size := lru.Size()
	wantSize := 2

	if size != wantSize {
		t.Errorf("want size %d, but got %d", wantSize, size)
	}

	value := lru.Get(2)
	wantValue := 3

	if value != wantValue {
		t.Errorf("want value %d, but got %d", wantValue, value)
	}
}

func TestLRUCache_SequentialOperations(t *testing.T) {
	lru := lrucache.Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)

	got := lru.Get(1)
	want := 1
	if got != want {
		t.Errorf("Get(1) want %d, but got %d", want, got)
	}

	lru.Put(3, 3)
	got = lru.Get(2)
	want = -1
	if got != want {
		t.Errorf("Get(2) want %d, but got %d", want, got)
	}

	lru.Put(4, 4)

	got = lru.Get(1)
	want = -1
	if got != want {
		t.Errorf("Get(1) want %d, but got %d", want, got)
	}

	got = lru.Get(3)
	want = 3
	if got != want {
		t.Errorf("Get(3) want %d, but got %d", want, got)
	}

	got = lru.Get(4)
	want = 4
	if got != want {
		t.Errorf("Get(4) want %d, but got %d", want, got)
	}
}

func TestLRUCache_ComplexSequence(t *testing.T) {
	lru := lrucache.Constructor(2)

	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)

	got := lru.Get(1)
	want := -1
	if got != want {
		t.Errorf("Get(1) want %d, but got %d", want, got)
	}

	got = lru.Get(2)
	want = 3
	if got != want {
		t.Errorf("Get(2) want %d, but got %d", want, got)
	}
}
