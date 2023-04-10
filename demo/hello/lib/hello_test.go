package lib

import (
	"fmt"
	"testing"
)

// run tests with; ... % cd lib;  % go test . -v

func Test_hello(t *testing.T) {
	expect := "lib world"
	if Hello(nil) != expect {
		t.Errorf("failed, expect: %s", expect)
	}

	empty := ""
	if Hello(&empty) != expect {
		t.Errorf("failed, expect: %s", expect)
	}

	kitty := "kitty"
	expect = fmt.Sprintf("lib %s", kitty)
	if Hello(&kitty) != expect {
		t.Errorf("failed, expect: %s", expect)
	}

	eleven := 11
	expect = fmt.Sprintf("lib number %d", eleven)
	if HelloInt(eleven) != expect {
		t.Errorf("failed, expect: %s", expect)
	}

}

func Test_hello_with_table(t *testing.T) {
	helloEntries := []struct {
		name            string
		values          []string
		expectedResults []string
		msg             string
	}{
		{name: "first test", values: []string{"Bob", ""}, expectedResults: []string{"lib Bob", "lib world"}, msg: "first lib test(s)"},
		{name: "next test", values: []string{"there"}, expectedResults: []string{"lib there"}, msg: "next lib test"},
	}

	for _, e := range helloEntries {
		for i, a := range e.values {
			r := Hello(&a)
			if r != e.expectedResults[i] {
				t.Errorf("failed, %s; result: %s != expected: %s", e.msg, r, e.expectedResults[i])
			}
		}
	}
}
