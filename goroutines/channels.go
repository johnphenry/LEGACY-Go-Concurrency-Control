/*
New strategy to implement at some point:
I'm thinking every routine gets it's own channel instead.
*/
package main
import(
"fmt"
"time"
)
var(
	sync chan int = make(chan int)
	response chan int = make(chan int)
)

type Object struct{
	x,y float32
}
var(
	objects []Object
)

func Frame(){
	for range(objects){
		<- response
	}
	for range(objects){
		sync <- 0
	}	
}
func WaitFrame(){
	response <- 0
	<- sync
}
func update(this *Object){
	{
		fmt.Printf("First\n")
		WaitFrame()
		fmt.Printf("Second\n")
		WaitFrame()
	}
}
func main(){
	for i:=0;i<2;i++{
		var ob Object
		objects = append( objects, ob )
		go update( &objects[ len(objects) - 1 ] )
	}
	for{
		time.Sleep(1000 * time.Millisecond)
		Frame()
	}
}
