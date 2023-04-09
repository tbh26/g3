package main

import (
	"fmt"
	"testing"
)

// run tests with; ... % cd hello;  % go test . -v

func Test_hello(t *testing.T) {
	expect := "hello world"
	if hello(nil) != expect {
		t.Errorf("failed, expect: %s", expect)
	}

	empty := ""
	if hello(&empty) != expect {
		t.Errorf("failed, expect: %s", expect)
	}

	kitty := "kitty"
	expect = fmt.Sprintf("hello %s", kitty)
	if hello(&kitty) != expect {
		t.Errorf("failed, expect: %s", expect)
	}

	eleven := 11
	expect = fmt.Sprintf("hello number %d", eleven)
	if helloInt(eleven) != expect {
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
		{name: "first test", values: []string{"Bob", ""}, expectedResults: []string{"hello Bob", "hello world"}, msg: "first hello test(s)"},
		{name: "next test", values: []string{"there"}, expectedResults: []string{"hello there"}, msg: "next hello test"},
	}

	for _, e := range helloEntries {
		for i, a := range e.values {
			r := hello(&a)
			if r != e.expectedResults[i] {
				t.Errorf("failed, %s; result: %s != expected: %s", e.msg, r, e.expectedResults[i])
			}
		}
	}
}
