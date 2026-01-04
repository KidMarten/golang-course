package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Error code: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Unkown erro type")
	}
}

// generics
func sum[T int | float32 | float64](a, b T) T {
	return a + b
}
