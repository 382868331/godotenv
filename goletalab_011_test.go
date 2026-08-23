package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv011(t *testing.T){if got:=doubleQuoteEscape("a"+string(rune(10))+"b");got!="a"+string(rune(92))+"n"+"b"{t.Fatalf("got=%q",got)}}
