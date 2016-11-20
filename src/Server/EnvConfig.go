package Server

import (
	"gopkg.in/yaml.v2"
	"fmt"
	"io/ioutil"
)

type EnvConfig interface {
	LoadGlobalSetting()
	TradingSessionType() int 
	ExecutionSystemType() int
}

type EnvConfigImpl struct {
	//TODO
	_tradingSessionType int
	_executionSystemType int
}

func (config *EnvConfigImpl) LoadGlobalSetting() {
	//TODO
	//从数据库或者配置文件加载配置数据
	config._tradingSessionType = ContinousMarket
	config._executionSystemType = OrderDriven
}

func (config *EnvConfigImpl) TradingSessionType() int {
	//TODO 暂时只用一个session，所以只返回一个类型
	return config._tradingSessionType
}

func (config *EnvConfigImpl) ExecutionSystemType() int {
	//TODO 暂时只用一个session，所以只返回一个类型
	return config._executionSystemType
}

func CreateEnvConfig() EnvConfig {
	envConfig := new(EnvConfigImpl)
	envConfig.LoadGlobalSetting()
	return envConfig
}


//解析文件，取出所有参数
func GetYamlConfig() map[interface{}]interface{}{
	//将参数从config.yaml中读取出来
	data, err := ioutil.ReadFile("config.yaml")
	//将解析出的参数转化为map的形式
	m := make(map[interface{}]interface{})
	if err != nil {
		//LogErr("error: %v", err)
		fmt.Println("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(data), &m)
	
	return m
}

//根据需求取出对应值  string是要读取属性的字符串 themap是读取config.yaml生成的map
func GetElement(key string, themap map[interface{}]interface{})string {
	if value, ok := themap[key]; ok {
		return value.(string)
	}
	
	//LogErr("Can't find the *.yaml")
	fmt.Println("Can't find the *.yaml")
	return ""
}