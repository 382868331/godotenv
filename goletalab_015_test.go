package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv015(t *testing.T){m,e:=Unmarshal("MY_KEY=value");if e!=nil||m["MY_KEY"]!="value"{t.Fatalf("m=%v err=%v",m,e)}}
