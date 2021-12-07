package robcommon

import "testing"

func TestReadCSIntList(t *testing.T) {
	ExpectedInputD6 := []uint64{3, 4, 3, 1, 2}
	ExpectedInputD7 := []uint64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	Input, err := ReadCSIntList("testInputD6")
	if err != nil {
		t.Fatalf("ReadCSIntList error'd: %v", err)
	}

	if !compareInputArray(Input, ExpectedInputD6) {
		t.Errorf("Input array(D6) did not match\nExpected: %v\n     Got: %v", Input, ExpectedInputD6)
	}

	Input, err = ReadCSIntList("testInputD7")
	if err != nil {
		t.Fatalf("ReadCSIntList error'd: %v", err)
	}

	if !compareInputArray(Input, ExpectedInputD7) {
		t.Errorf("Input array(D7) did not match\nExpected: %v\n     Got: %v", Input, ExpectedInputD7)
	}
}

func compareInputArray(one, two []uint64) bool {
	if one == nil || two == nil {
		return false
	}
	if len(one) != len(two) {
		return false
	}

	for i := 0; i < len(one); i++ {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
