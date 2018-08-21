package deptestone // import "github.com/shenderov/DepTestOne"

import (
	"fmt"
	"github.com/shenderov/HelloGo"
)

func Hello() {
	fmt.Println("This is dep two")
	hellogo.Hello()
}
