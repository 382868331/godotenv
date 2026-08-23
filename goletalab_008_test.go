package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv008(t *testing.T){got,_:=Marshal(map[string]string{"A":"1","B":"2"});if got!="A=1\nB=2"{t.Fatalf("got=%q",got)}}

func TestGoletaGodotenv008AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	got,_:=Marshal(map[string]string{"B":"2","A":"1","C":"3"});if strings.Count(got,"\n")!=2||strings.Contains(got,"\r"){t.Fatalf("got=%q",got)}
}
