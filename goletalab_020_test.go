package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv020(t *testing.T){if !isLineEnd(rune(10))||!isLineEnd(rune(13)){t.Fatal("line endings rejected")}}
