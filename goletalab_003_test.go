package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv003(t *testing.T){m,e:=UnmarshalBytes([]byte("A=1"));if e!=nil||m["A"]!="1"{t.Fatalf("m=%v err=%v",m,e)}}

func TestGoletaGodotenv003AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m,e:=UnmarshalBytes([]byte{});if e!=nil||m==nil{t.Fatalf("m=%v",m)}
}
