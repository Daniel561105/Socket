package hello

import (
	"fmt"
)

func Welcome(name string) string {
	message := fmt.Sprintf("Welcome %v", name)
	return message
}
