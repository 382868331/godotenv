package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv020(t *testing.T){if !isLineEnd(rune(10))||!isLineEnd(rune(13)){t.Fatal("line endings rejected")}}

func TestGoletaGodotenv020AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m,e:=Unmarshal("A=1\nB=2");if e!=nil||len(m)!=2{t.Fatalf("m=%v",m)}
}
