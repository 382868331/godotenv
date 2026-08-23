package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv015(t *testing.T){m,e:=Unmarshal("MY_KEY=value");if e!=nil||m["MY_KEY"]!="value"{t.Fatalf("m=%v err=%v",m,e)}}

func TestGoletaGodotenv015AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m,e:=Unmarshal("_LEADING=ok");if e!=nil||m["_LEADING"]!="ok"{t.Fatalf("m=%v",m)}
}
