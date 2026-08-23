package godotenv
import("os";"path/filepath";"strings";"testing")
var _=os.Getenv;var _=filepath.Join;var _=strings.Contains
func TestGoletaGodotenv010(t *testing.T){k:="GOLETA_ENV_018_A";os.Unsetenv(k);f:=filepath.Join(t.TempDir(),"a.env");os.WriteFile(f,[]byte(k+"=value\n"),0600);if e:=loadFile(f,false);e!=nil||os.Getenv(k)!="value"{t.Fatalf("value=%q err=%v",os.Getenv(k),e)}}

func TestGoletaGodotenv010AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	k:="GOLETA_ENV_018_B";os.Setenv(k,"old");f:=filepath.Join(t.TempDir(),"b.env");os.WriteFile(f,[]byte(k+"=new\n"),0600);_ = loadFile(f,true);if os.Getenv(k)!="new"{t.Fatalf("value=%q",os.Getenv(k))}
}
