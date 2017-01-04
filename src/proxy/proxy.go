package proxy

import (
	"log"
	"net"
	"net/url"
	"bytes"
	"strings"
	"io"
	"fmt"
)

func StartProxy(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":8081")
	if err != nil{
		log.Panic(err)
	}
	
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn){
	if client == nil {
		return
	}

	defer client.Close()
	remoteAddr := client.RemoteAddr()
	fmt.Println(remoteAddr)	

	//var b [1024]byte
	b := make([]byte, 1024)
	n, err := client.Read(b)
	if err != nil{
		log.Println(err)
		return
	}

	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	hostPortUrl, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(hostPortUrl)	

	if hostPortUrl.Opaque == "443" {//https访问
		address = hostPortUrl.Scheme + ":443"
	}else{//http访问
		if strings.Index(hostPortUrl.Host, ":") == -1{
			//不带端口号，默认为80
			address = hostPortUrl.Host + ":80"
		}else{
			address = hostPortUrl.Host
		}
	}

	//获取了请求的host和port，可以开始拨号了
	server, err := net.Dial("tcp", address)
	if err != nil{
		log.Println(err)
		return
	}

	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n")
	}else{
		server.Write(b[:n])
	}
	
	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}
