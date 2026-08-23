package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv018(t *testing.T){input:="A="+string(rune(34))+"line\\nnext"+string(rune(34));m,e:=Unmarshal(input);if e!=nil||m["A"]!="line"+string(rune(10))+"next"{t.Fatalf("m=%q err=%v",m,e)}}

func TestGoletaGodotenv018AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m,e:=Unmarshal("A='line\nnext'");if e!=nil||m["A"]!="line\nnext"{t.Fatalf("m=%q",m)}
}
