package main
import 
(
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	counter:=200  
	// wait.Add(1)
	wait.Add(counter)
	for i:=0; i<counter; i++{
		go hello(&wait)

	}


	defer wait.Wait()
	
}
func hello(wait *sync.WaitGroup) {
	fmt.Print(1)
	wait.Done()
}

