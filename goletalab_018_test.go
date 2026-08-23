package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv018(t *testing.T){input:="A="+string(rune(34))+"line\\nnext"+string(rune(34));m,e:=Unmarshal(input);if e!=nil||m["A"]!="line"+string(rune(10))+"next"{t.Fatalf("m=%q err=%v",m,e)}}
