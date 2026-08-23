package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv011(t *testing.T){if got:=doubleQuoteEscape("a"+string(rune(10))+"b");got!="a"+string(rune(92))+"n"+"b"{t.Fatalf("got=%q",got)}}

func TestGoletaGodotenv011AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	got,_:=Marshal(map[string]string{"X":"a"+string(rune(10))+"b"});if !strings.Contains(got,"a"+string(rune(92))+"n"+"b"){t.Fatalf("got=%q",got)}
}
