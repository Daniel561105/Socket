package hello

import (
	"fmt"
)

func Welcome(name string) string {
	message := fmt.Sprintf("Welcome %v for the first module use", name)
	return message
}
