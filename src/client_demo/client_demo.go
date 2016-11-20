package main 

import (
	//"protocol"
	"fmt"
	"net"
	"os"
)


func main() {
	server := "localhost:6060"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	
	fmt.Println("connect success")
	
	
	//缓冲区，存储被截断的数据
// 	tmpBuffer := make([]byte, 0)
	
	//接收解包
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel, conn)
	buffer := make([]byte, 1024)
	
	var input string
	fmt.Scanln(&input)
	for input != "exit" {
		conn.Write(([]byte(input)))
		
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "connection error:", err)
			return
		}
// 		tmpBuffer = protocol.Depack(append(tmpBuffer, buffer[:n]...), readerChannel)
		
		//处理完解包后才允许再次输入
		fmt.Scanln(&input)
	}
	conn.Write(([]byte(input)))
}

func reader(readerChannel chan []byte, conn net.Conn) {
	for {
		select {
			case data := <-readerChannel:
				fmt.Println(string(data))
			//case <-readerChannel:
		}
	}
}
