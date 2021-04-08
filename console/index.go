package console

import (
	"fmt"
	"os"
)

// SetTitle ...
func SetTitle(message string) {
	os.Stdout.WriteString(fmt.Sprintf("\033]0;%s\007", message))
}
