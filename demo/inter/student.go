package inter

import (
	"fmt"
	"inter/data"
)

type Student struct{}

var _ Output = (*Student)(nil)

func (s *Student) Printf(stu data.Person) {
	if stu.Job == "student" {
		fmt.Printf("Student %v\n", stu)
	}
}
