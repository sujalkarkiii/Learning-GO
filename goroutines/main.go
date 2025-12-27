package main

import (
	"fmt"
	"time"
)
func printMessage(text string, channel chan string) {
	for i := 0; i < 2; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
	channel <- "Done" // send message to main goroutine
		close(channel)   
}

func main(){ //main goroutine


// here we have to wait for first function pass
// to complete along some time to complete
//  the task and same for second

//  go lang routines=light weight thread
// thread=different lines of execution running 
// at same time
// we have to use go at prefix of any function call
// we have maingoroutine printgoroutine initially maingoroutine runs till } 
// if all are gorotuine than since all are running in own rotuine
//  main dont have anything to do that it does nothing


	channel := make(chan string)
	go printMessage("go is great",channel)
	response := <-channel
	fmt.Println(response)

}

