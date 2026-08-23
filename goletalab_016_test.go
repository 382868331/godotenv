package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv016(t *testing.T){m,e:=Unmarshal("A=value");if e!=nil||m["A"]!="value"{t.Fatalf("m=%v err=%v",m,e)}}
