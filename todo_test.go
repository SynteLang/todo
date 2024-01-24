// these are very minimal tests, and not in a good way.
// to better ensure correct function of 'todo':
// add failing tests
// add more test cases
// moar coverage
// try fuzzing

package main

import (
	"fmt"
	"testing"
)

type testJig struct {
	in string
	td todoList
	exp todoList
}

var testList = map[string]testJig{
	"pushString()": { "pushtest", todoList{"item"}, todoList{"pushtest", "item"}},
	"done()": { "", todoList{"donetest", "item"}, todoList{"item"}},
	"swap()": { "", todoList{"item one", "item two", "swaptest"}, todoList{"item two", "item one", "swaptest"}},
}

func init() {
	flag = flagDummy
}

func TestPushString(t *testing.T) {
	name := "pushString()"
	tst := testList[name]
	tst.td.pushString(tst.in)
	fmt.Println(name)
	if !passing(tst) {
		t.Errorf(`pushString(%q) => %v, expected %v`, tst.in, tst.td, tst.exp)
	}
}

func passing(tst testJig) bool {
	for i,v := range tst.td {
		if v != tst.exp[i] {
			return false
		}
	}
	if len(tst.exp) != len(tst.td) {
		return false
	}
	return true
}

func flagDummy(s string) bool {
/*	if len(os.Args) > 2 {
		return false
	}
	if len(os.Args) == 2 && os.Args[1] == s {
		return true
	}
	return false*/
	return true
}

func TestDone(t *testing.T) {
	name := "done()"
	tst:= testList[name]
	fmt.Println(name)
	tst.td.done()
	if !passing(tst) {
		t.Errorf(`%s => %v, expected %v`, name, tst.td, tst.exp)
	}
}

func TestSwap(t *testing.T) {
	name := "swap()"
	tst:= testList[name]
	fmt.Println(name)
	tst.td.swap(0)
	if !passing(tst) {
		t.Errorf(`%s => %v, expected %v`, name, tst.td, tst.exp)
	}
}
