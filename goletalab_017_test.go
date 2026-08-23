package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv017(t *testing.T){m,e:=Unmarshal("A=\nB=2");if e!=nil||m["A"]!=""||m["B"]!="2"{t.Fatalf("m=%v err=%v",m,e)}}
