package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv001(t *testing.T){m,e:=Parse(strings.NewReader("KEY=value\n"));if e!=nil||m["KEY"]!="value"{t.Fatalf("m=%v err=%v",m,e)}}

func TestGoletaGodotenv001AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m,e:=Parse(strings.NewReader("A=1\nB=2\n"));if e!=nil||m["A"]!="1"||m["B"]!="2"{t.Fatalf("m=%v",m)}
}
