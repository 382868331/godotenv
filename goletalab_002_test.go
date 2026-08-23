package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv002(t *testing.T){m,e:=Unmarshal("KEY=value");if e!=nil||m["KEY"]!="value"{t.Fatalf("m=%v err=%v",m,e)}}
