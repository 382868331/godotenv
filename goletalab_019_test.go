package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv019(t *testing.T){if got:=expandEscapes("a\\nb");got!="a"+string(rune(10))+"b"{t.Fatalf("got=%q",got)}}
