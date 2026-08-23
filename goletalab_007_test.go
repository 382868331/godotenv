package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv007(t *testing.T){got,_:=Marshal(map[string]string{"N":"12","S":"abc"});want:="S="+string(rune(34))+"abc"+string(rune(34));if !strings.Contains(got,"N=12")||!strings.Contains(got,want){t.Fatalf("got=%q",got)}}
