package enum

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	var a StatusType = Normal
	s := a.String()
	fmt.Println(s)
}
