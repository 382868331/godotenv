package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv006(t *testing.T){if !isInt("19"){t.Fatal("digit 9 rejected")}}

func TestGoletaGodotenv006AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if !isInt("999"){t.Fatal("digit 9 rejected")}
}
