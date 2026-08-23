package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv014(t *testing.T){m,e:=Unmarshal("A: value");if e!=nil||m["A"]!="value"{t.Fatalf("m=%v err=%v",m,e)}}

func TestGoletaGodotenv014AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m,e:=Unmarshal("B:value2");if e!=nil||m["B"]!="value2"{t.Fatalf("m=%v",m)}
}
