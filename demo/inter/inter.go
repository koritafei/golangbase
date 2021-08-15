package inter

import (
	"inter/data"
)

type Output interface {
	Printf(stu data.Person)
}
