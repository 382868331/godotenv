package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv001(t *testing.T){m,e:=Parse(strings.NewReader("KEY=value\n"));if e!=nil||m["KEY"]!="value"{t.Fatalf("m=%v err=%v",m,e)}}
