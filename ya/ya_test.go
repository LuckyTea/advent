package ya

import "testing"

func TestCase1(t *testing.T) {
	var (
		in   = []int{1, 4, 5, 2, 3, 9, 8, 11, 0}
		want = "0-5,8-9,11"
	)

	out := transform(in)

	if out != want {
		t.Errorf("fail: %v", out)
	}
}
func TestCase2(t *testing.T) {
	var (
		in   = []int{1, 4, 3, 2}
		want = "1-4"
	)

	out := transform(in)

	if out != want {
		t.Errorf("fail: %v", out)
	}
}

func TestCase3(t *testing.T) {
	var (
		in   = []int{1, 4}
		want = "1,4"
	)

	out := transform(in)

	if out != want {
		t.Errorf("fail: %v", out)
	}
}

func TestCase4(t *testing.T) {
	var (
		in   = []int{1, 2, 4}
		want = "1-2,4"
	)

	out := transform(in)

	if out != want {
		t.Errorf("fail: %v", out)
	}
}
