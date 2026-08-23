package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv013(t *testing.T){m,e:=Unmarshal("export A=1");if e!=nil||m["A"]!="1"{t.Fatalf("m=%v err=%v",m,e)}}
