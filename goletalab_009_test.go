package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv009(t *testing.T){got:=filenamesOrDefault(nil);if len(got)!=1||got[0]!=".env"{t.Fatalf("got=%v",got)}}

func TestGoletaGodotenv009AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	got:=filenamesOrDefault([]string{});if got[0]!=".env"{t.Fatalf("got=%v",got)}
}
