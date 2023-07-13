package question

import (
	"fmt"
	"strings"
)

func Test() {
	fmt.Println(strings.Count("abcdedf", "ad"))
	fmt.Println(strings.Count("abcdedf", "ab"))
}
