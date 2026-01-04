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
	// intValue, ok := a.(int) - не сработает
	// switch d := a.(type) {case ...} - не сработает
	// return "1" не сработает, return должен попадать под все типы
	// нельзя передавать интерфейс (например, error)
	return a + b
}

// Generic struct
type List[T any] struct {
	elements []T
}
