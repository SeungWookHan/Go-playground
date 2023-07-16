package structure

import "fmt"

type TestStruct struct {
	TestStructValue int
}

func (t TestStruct) String() string {
	return fmt.Sprintf("TestStructValue is %d", t.TestStructValue)
}

func (t *TestStruct) Increment() {
	t.TestStructValue = t.TestStructValue + t.one()
}

func (t *TestStruct) one() int {
	return 1
}
