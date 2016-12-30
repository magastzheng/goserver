package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	ReadLine()
}

func ReadLine() {
	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data[0:])
		fmt.Println(command)
		if command == "stop"{
			running = false
		}
	}
}
