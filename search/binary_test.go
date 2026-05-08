func TestBinaryWithDuplicates(t *testing.T) {
	arr := []int{1, 2, 4, 4, 4, 6, 7, 8}
	target := 4

	result, err := Binary(arr, target, 0, len(arr)-1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result < 0 || result >= len(arr) {
		t.Fatalf("returned invalid index: %d", result)
	}

	if arr[result] != target {
		t.Fatalf(
			"expected target %d at returned index, got %d",
			target,
			arr[result],
		)
	}
}

func TestBinaryIterativeWithDuplicates(t *testing.T) {
	arr := []int{1, 2, 4, 4, 4, 6, 7, 8}
	target := 4

	result, err := BinaryIterative(arr, target)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result < 0 || result >= len(arr) {
		t.Fatalf("returned invalid index: %d", result)
	}

	if arr[result] != target {
		t.Fatalf(
			"expected target %d at returned index, got %d",
			target,
			arr[result],
		)
	}
}
