package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv013(t *testing.T){m,e:=Unmarshal("export A=1");if e!=nil||m["A"]!="1"{t.Fatalf("m=%v err=%v",m,e)}}

func TestGoletaGodotenv013AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m,e:=Unmarshal("export LONG=value");if e!=nil||m["LONG"]!="value"{t.Fatalf("m=%v",m)}
}
