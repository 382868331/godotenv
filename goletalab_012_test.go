package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv012(t *testing.T){m,e:=Unmarshal("A=1\n");if e!=nil||m["A"]!="1"{t.Fatalf("m=%v err=%v",m,e)}}

func TestGoletaGodotenv012AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m,e:=Unmarshal("# comment\nB=2\n");if e!=nil||m["B"]!="2"{t.Fatalf("m=%v err=%v",m,e)}
}
