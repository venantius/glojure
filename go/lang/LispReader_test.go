package lang

import (
	"fmt"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	r := strings.NewReader("[1]")
	fmt.Println(r.ReadRune())


}
