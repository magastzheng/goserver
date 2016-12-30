package main

import (
	"fmt"
	"net"
	//"bufio"
	"time"
	//"encoding/binary"
	//"bytes"
	"protocol"
)


func main(){
	
	//tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:90001")
	//if err != nil{
	//	fmt.Println("Fail to parse the IP address")
	//}
	
	//conn, err := net.DialTCP("tcp", nil, tcpAddr)
	//if err != nil{
	//	fmt.Println("Fail to DialTCP")
	//	panic(err)
	//}
	
	//listen, err := net.ListenTCP("tcp", tcpAddr)
	//if err != nil{
	//	panic(err)
	//}
	//conn, err = listen.Accept()
	//
	//if err != ni{
	//	panic(err)
	//}
	//defer conn.Close()
	
	//conn, err := net.Dial("tcp", "127.0.0.1:90001")
	//if err != nil {
	//	panic(err)
	//}
	
	listener, err := net.Listen("tcp", ":9001")
	if err != nil{
		panic(err)
	}
	fmt.Println("Server is started...")
	for{
		conn, err := listener.Accept()
		if err != nil{
			//fmt.Println("Cannot accept the connection")
		}
		fmt.Println("A client connected: ", conn.RemoteAddr())
		go handleConnection(conn)
	}	
	fmt.Println("Test go server")
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	
	fmt.Println("handle connection")
	//reader := bufio.NewReader(conn)

	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		//message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			return
		}
		
		pkgbuf,_ := protocol.Read(buf)
		
		fmt.Println("Accept data from client: ", string(pkgbuf[0:]))
		//fmt.Println("Accept data from client: ", message)
	
		//send reply
		replyMsg := time.Now().String() + "\n"
		b := []byte(replyMsg)
		_, err = conn.Write(b)
		if err != nil {
			fmt.Println("Error send reply: ", err.Error())
			return
		}
		
	}
}
