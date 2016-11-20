package Server

import (
	"net"
	"fmt"
	"os"
	"fix"
	//"protocol"
)


func HandleConnection(conn net.Conn, operationLog Log) {
	//缓冲区，存储被截断的数据
	//tmpBuffer := make([]byte, 0)
	
	//接收解包
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel, conn, operationLog)
	
	buffer := make([]byte, 1024)
	for {
		//buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "connection error:", err)
			operationLog.Log(conn.RemoteAddr().String()+" connection error: "+err.Error())
			return
		}
		
		fmt.Println(string(buffer[:n]))
		operationLog.Log(string(buffer[:n]))
		fmt.Println(fix.Decode(string(buffer[:n])))
		conn.Write([]byte("sb\n"))
		//tmpBuffer = protocol.Depack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
	defer conn.Close()
}

func reader(readerChannel chan []byte, conn net.Conn, operationLog Log) {
	for {
		select {
			//数据在这，当缓冲区中有数据进来时，就会执行
			case data := <-readerChannel:
				fmt.Println(string(data))
				operationLog.Log(string(data))
				//每收到一条数据就返回一条信息
				conn.Write([]byte("sb\n"))
		}
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}