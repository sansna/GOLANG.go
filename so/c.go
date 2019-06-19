package main
import (
	"log"
	"os"
	"os/exec"
	"fmt"
)

func main() {
	cmd := exec.Command("./b")
	cmd.Env = append(os.Environ(),
		"LD_LIBRARY_PATH=.",
	)
	// Simply runs it, ignores output.
	//if err := cmd.Run(); err != nil {
	//	log.Fatal(err)
	//}

	// This returns stdout as out.
	var out []byte
	var err error
	if out, err = cmd.Output(); err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}
