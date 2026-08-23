package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv002(t *testing.T){m,e:=Unmarshal("KEY=value");if e!=nil||m["KEY"]!="value"{t.Fatalf("m=%v err=%v",m,e)}}

func TestGoletaGodotenv002AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m,e:=Unmarshal("A=1\nB=2");if e!=nil||len(m)!=2{t.Fatalf("m=%v",m)}
}
