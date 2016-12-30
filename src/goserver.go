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

//每次读入部分字节,如果本次读取的数据没有结束,再读入更多字符。直到读完本次数据包为止
//此时会出现多读入的情况,在读下一个包的时候,需要回填过来
//如果设计开始标记和结束标记,则需要对数据做编码解码[发送前编码,接收后先解码]
//是否需要添加校验标识符
func handleConnection(conn net.Conn){
	defer conn.Close()
	
	fmt.Println("handle connection")
	//reader := bufio.NewReader(conn)

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			return
		}
		
		fmt.Println("accept data: ", buf[0:n])	
		pkgbuf,_ := protocol.Read(buf[0:n])
		
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
