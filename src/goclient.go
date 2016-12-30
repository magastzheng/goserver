package main

import (
	"fmt"
	"net"
	//"encoding/binary"
	//"bytes"
	"protocol"
	"bufio"
	"os"
)


func main(){
	//connect to a server
	conn, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	
	fmt.Println("Connected to the server")

	handleConnection(conn)
		
	fmt.Println("Test go client")
}

func handleConnection(conn net.Conn){

	for {
		reader := bufio.NewReader(os.Stdin)
		in, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Input exception: ", err.Error())
			//return
		}
		//input data
		fmt.Println("input data: ", in)
		
		//send data
		data := protocol.Write(in[0:len(in)])
		fmt.Println("input package: ", data)	
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error send data: ", err.Error())
		}
		
		inStr := string(in[0:])
		if inStr == "quit" {
			break
		}
				
		//read server data
		out := make([]byte, 1024)
		c, err := conn.Read(out)
		if err != nil {
			fmt.Println("Read server data exception: ", err)
		}
			
		fmt.Println("Data from server: ", string(out[0:c]))

	}
}
