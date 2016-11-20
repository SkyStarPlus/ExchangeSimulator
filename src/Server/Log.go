package Server

import (
	//"bufio"
	//"fmt"
	//"io/ioutil"
	"io"
	"os"
	"time"
)

//LOG的使用方法举例
//要使用Log时，先通过CreateLog(日志文件名)方法生成一个Log对象，然后调用Log（日志内容）记录日志，Log方法会自动记录记日志时间
//改了个\r\n，每次调用Log时，后面加空格
type Log interface {
	Log(string)
}

type LogImpl struct {
	_logFileName string
	_file *os.File
}

func (l *LogImpl) Init(fileName string) {
	l._logFileName = fileName
	var err error
	if l.checkFilesExist(l._logFileName) {
		l._file, err = os.OpenFile(l._logFileName, os.O_APPEND, 0666)
	} else {
		l._file, err = os.Create(l._logFileName)
	}
	l.check(err)
}

func (l *LogImpl) checkFilesExist(filename string) bool {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func (l *LogImpl) check(e error) {
	if e != nil {
		panic(e)	
	}
}

func (l *LogImpl) Log(log string) {
	currentTime := time.Now().String()
	sentence := currentTime + ">>" + log + "\r\n"	//改了个加回车
	io.WriteString(l._file, sentence)
}

func CreateLog(fileName string) Log {
	log := new(LogImpl)
	log.Init(fileName);
	return log
}