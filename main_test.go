package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	negativeSizeData := generateRandomElements(-1)
	if len(negativeSizeData) != 0 {
		t.Errorf("size -1: length = %d, want 0", len(negativeSizeData))
	}

	data0 := generateRandomElements(0)
	if len(data0) != 0 {
		t.Errorf("size 0: length = %d, want 0", len(data0))
	}

	data1 := generateRandomElements(1)
	if len(data1) != 1 {
		t.Errorf("size 1: length = %d, want 1", len(data1))
	}
	for _, value := range data1 {
		if value <= 0 {
			t.Errorf("size 1: nonpositive value %d", value)
		}
	}

	data8 := generateRandomElements(8)
	if len(data8) != 8 {
		t.Errorf("size 8: length = %d, want 8", len(data8))
	}
	for _, value := range data8 {
		if value <= 0 {
			t.Errorf("size 8: nonpositive value %d", value)
		}
	}

	data1000 := generateRandomElements(1000)
	if len(data1000) != 1000 {
		t.Errorf("size 1000: length = %d, want 1000", len(data1000))
	}
	for _, value := range data1000 {
		if value <= 0 {
			t.Errorf("size 1000: nonpositive value %d", value)
		}
	}
}

func TestMaximum(t *testing.T) {
	if got := maximum(nil); got != 0 {
		t.Errorf("nil: got %d, want 0", got)
	}

	if got := maximum([]int{}); got != 0 {
		t.Errorf("empty: got %d, want 0", got)
	}

	if got := maximum([]int{7}); got != 7 {
		t.Errorf("single: got %d, want 7", got)
	}

	if got := maximum([]int{9, 2, 1}); got != 9 {
		t.Errorf("first: got %d, want 9", got)
	}

	if got := maximum([]int{1, 2, 9}); got != 9 {
		t.Errorf("last: got %d, want 9", got)
	}

	if got := maximum([]int{3, 3, 3}); got != 3 {
		t.Errorf("equal: got %d, want 3", got)
	}

	if got := maximum([]int{0, 0}); got != 0 {
		t.Errorf("zero: got %d, want 0", got)
	}

	if got := maximum([]int{-8, -2, -5}); got != -2 {
		t.Errorf("negative: got %d, want -2", got)
	}

	if got := maximum([]int{1, 2, 3, 10, 5, 6, 7, 8}); got != 10 {
		t.Errorf("eight: got %d, want 10", got)
	}

	if got := maximum([]int{1, 2, 3, 4, 5, 6, 7, 8, 99}); got != 99 {
		t.Errorf("remainder: got %d, want 99", got)
	}
}

func TestMaxChunks(t *testing.T) {
	if got := maxChunks(nil); got != 0 {
		t.Errorf("nil: got %d, want 0", got)
	}

	if got := maxChunks([]int{}); got != 0 {
		t.Errorf("empty: got %d, want 0", got)
	}

	if got := maxChunks([]int{7}); got != 7 {
		t.Errorf("single: got %d, want 7", got)
	}

	if got := maxChunks([]int{9, 2, 1}); got != 9 {
		t.Errorf("first: got %d, want 9", got)
	}

	if got := maxChunks([]int{1, 2, 9}); got != 9 {
		t.Errorf("last: got %d, want 9", got)
	}

	if got := maxChunks([]int{3, 3, 3}); got != 3 {
		t.Errorf("equal: got %d, want 3", got)
	}

	if got := maxChunks([]int{0, 0}); got != 0 {
		t.Errorf("zero: got %d, want 0", got)
	}

	if got := maxChunks([]int{-8, -2, -5}); got != -2 {
		t.Errorf("negative: got %d, want -2", got)
	}

	if got := maxChunks([]int{1, 2, 3, 10, 5, 6, 7, 8}); got != 10 {
		t.Errorf("eight: got %d, want 10", got)
	}

	if got := maxChunks([]int{1, 2, 3, 4, 5, 6, 7, 8, 99}); got != 99 {
		t.Errorf("remainder: got %d, want 99", got)
	}
}
