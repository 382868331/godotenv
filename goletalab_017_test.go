package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv017(t *testing.T){m,e:=Unmarshal("A=\nB=2");if e!=nil||m["A"]!=""||m["B"]!="2"{t.Fatalf("m=%v err=%v",m,e)}}

func TestGoletaGodotenv017AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m,e:=Unmarshal("A=x\n");if e!=nil||m["A"]!="x"{t.Fatalf("m=%v",m)}
}
