package handler

import (
	"fmt"
	"strings"
	"testing"
)

func TestUpsertEduProduct(t *testing.T) {
	join := strings.Join([]string{"123", "5555", " ", "7777"}, ";")
	fmt.Println(join)
}
