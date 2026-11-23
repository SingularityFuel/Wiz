package u

import "fmt"

// Returns type as string
func Type(x any) string {
	return fmt.Sprintf("%T", x)
}
