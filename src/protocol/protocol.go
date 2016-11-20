//通信协议处理
package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader = "Headers"
	ConstHeaderLength = 7		//即Hearders的字节长度
	ConstMLength = 4			//int转换为byte的长度，4字节
)

//封包
func Enpack(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

//解包
func Depack(buffer []byte, readerChannel chan []byte) []byte {
	var i int
// 	length := len(buffer)
// 	
// 	//这边的i的变化情况估计得用断点调试才能看清楚了
//  	for i=0; i<length; i=i+1 {
// 		//如果缓冲区的大小比当前数据量加上头和消息长度字节还小，则break
//  		if length < i+ConstHeaderLength+ConstMLength {
//  			break
//  		}
// 		//如果当前的数据正好是接收到头
// 		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
// 			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstMLength])
// 			//如果缓冲区大小比当前数据量加上头和消息长度字节和消息长度还小，则break
//  			if length < i+ConstHeaderLength+ConstMLength+messageLength {
//  				break
//  			}
// 			data := buffer[i+ConstHeaderLength+ConstMLength : i+ConstHeaderLength+ConstMLength+messageLength]
 			readerChannel <- buffer
//  			readerChannel <- data
// 			//fmt.Println(i)
// 			//？这边不用把i改变大小吗？
// 		}
//  	}
// 	
// 	//这边是用来干嘛的？
// 	if i == length {
// 		return make([]byte, 0)
// 	}
	return buffer[i:]
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}