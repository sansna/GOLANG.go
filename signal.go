package main
import "os"
import "os/signal"
import "syscall"
import "fmt"
import "time"
func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		time.Sleep(10*time.Second)
	}()
	fmt.Println("Press <Ctrl-c> to terminate task gracefully.")
	<-c
	fmt.Println("\nok!")
	return
}
