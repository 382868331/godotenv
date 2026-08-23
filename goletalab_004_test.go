package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv004(t *testing.T){if !isInt("-12")||isInt("+12"){t.Fatal("sign handling wrong")}}

func TestGoletaGodotenv004AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if !isInt("0")||!isInt("42"){t.Fatal("positive integer rejected")}
}
