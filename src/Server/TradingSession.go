package Server

import (
	"errors"
)

const (
	ContinousMarket = iota
	CallMarket
)

type TradingSession interface {
	Init() 
	
	AddOrder(order Order) (bool, error)
	CancelOrder(order Order) (bool, error)
	QueryOrderBook(productId int) OrderBook
	
	SessionId() string
	SessionType() int
	
	SetExecutionSystemType(int)
	ExecutionSystemType() int
	
	Start()
	Suspend()
	End()//期间要检查有没有过期的Order,干掉
}

type ContinousTradingSession struct {
	_executionSystemType int
	_executionSystem ExecutionSystem
	_started bool 
	_log Log
}

func (s *ContinousTradingSession) Init() {
	s._executionSystem = CreateExecutionSystem(s._executionSystemType)
	s._started = false
	s._log = CreateLog("Server.txt")
}

func (s *ContinousTradingSession) AddOrder(order Order) (bool, error) {
	//TODO 这里可以考虑直接跟客户通信
	s._log.Log("AddOrder:" + order.OrderId())//日志使用举例
	if (s._started) {
		return s._executionSystem.AddOrder(order)
	} else {
		return false, errors.New("Maket is not running")
	}
}

func (s *ContinousTradingSession) CancelOrder(order Order)(bool, error) {
	if (s._started) {
		//TODO 这里可以直接根据结果跟客户通信
		return s._executionSystem.CancelOrder(order)
	} else {
		return false, errors.New("Market is not running")
	}
} 

func (s *ContinousTradingSession) QueryOrderBook(productId int) OrderBook {
	//TODO这里直接考虑跟客户端通信
	return s._executionSystem.QueryOrderBook(productId)
}

func (s *ContinousTradingSession) SessionId() string {
	return ""//TODO
}

func (s *ContinousTradingSession) SessionType() int {
	return ContinousMarket;
}

func (s *ContinousTradingSession) SetExecutionSystemType(executionSystemType int) {
	s._executionSystemType = executionSystemType
} 

func (s *ContinousTradingSession) ExecutionSystemType() int {
	return s._executionSystemType
}

func (s *ContinousTradingSession) Start() {
	s._started = true
}

func (s *ContinousTradingSession) Suspend() {
	s._started = false
}

func (s *ContinousTradingSession) End() {
	s._started =  false
	//TODO 将当天过期的order全部清理掉
}

type CallTradingSession struct {
	_executionSystemType int
	_executionSystem ExecutionSystem
	_started bool
}

func (s *CallTradingSession) Init() {
	//TODO
	s._executionSystem = CreateExecutionSystem(s._executionSystemType)
	s._started = false;
}

func (s *CallTradingSession) AddOrder(order Order) (bool, error){
	//TODO 这里可以考虑直接跟客户通信
	if (s._started) {
		return s._executionSystem.AddOrder(order)
	} else {
		return false, errors.New("Maket is not running")
	}
}

func (s *CallTradingSession) CancelOrder(order Order) (bool, error){
	if (s._started) {
		//TODO 这里可以直接根据结果跟客户通信
		return s._executionSystem.CancelOrder(order)
	} else {
		return false, errors.New("Market is not running")
	}
} 

func (s *CallTradingSession) QueryOrderBook(productId int) OrderBook {
	return s._executionSystem.QueryOrderBook(productId)
}

func (s *CallTradingSession) SessionId() string {
	return ""//TODO
}

func (s *CallTradingSession) SessionType() int {
	return CallMarket;
}

func (s *CallTradingSession) SetExecutionSystemType(executionSystemType int) {
	s._executionSystemType = executionSystemType
} 

func (s *CallTradingSession) ExecutionSystemType() int {
	return s._executionSystemType
}

func (s *CallTradingSession) Start() {
	s._started = true
}

func (s *CallTradingSession) Suspend() {
	s._started = false
}

func (s *CallTradingSession) End() {
	s._started = false
	//TODO 清理当天过期的order数据
}


func CreateTradingSession(sessionType int, executionSystemType int) TradingSession {
	if sessionType == ContinousMarket {
		exchange := new(ContinousTradingSession)//TODO
		exchange.SetExecutionSystemType(executionSystemType)
		exchange.Init()
		return exchange
	} else {
		exchange := new(CallTradingSession)//TODO
		exchange.SetExecutionSystemType(executionSystemType)
		exchange.Init()
		return exchange
	}
}
