package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv005(t *testing.T){if isInt("")||isInt("-"){t.Fatal("empty value treated as integer")}}
