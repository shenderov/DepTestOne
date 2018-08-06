package deptestone // import "github.com/shenderov/DepTestOne"

import (
	"fmt"
	"github.com/shenderov/HelloKostya"
)

func Hello() {
	fmt.Println("This is dep two")
	deptestone.Hello()
}
