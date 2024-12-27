// pkg/utils/utils.go
package utils

import (
	"strings"
)

func IndentCode(code string, spaces int) string {
	indent := strings.Repeat(" ", spaces)
	lines := strings.Split(code, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = indent + line
		}
	}
	return strings.Join(lines, "\n")
}
