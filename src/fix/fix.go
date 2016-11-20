package fix

import (
    "fmt"
    "strings"
    "strconv"
)

// formal struct, not complete yet
type Message struct{
    username, password, order_type, order_id, price, quantity, action, stock_id, trade_type string
}

// Struct for mid term 
type MessageTemp struct{
    username, password, order_type, trade_type string
    quantity, price, stock_id int
}

// Decode message from client, return a Message struct.
func Decode(msg string) MessageTemp {
    var r = make([]string,15)
    r = strings.Split(msg, ";")
    var result = MessageTemp{}
    for _,v := range r {
        key_value := strings.Split(v,"=")
        switch key_value[0] {
            case "35":
                if key_value[1]=="8"{ // 35=8; 股票交易相关信息
                    result.order_type = "trade"
                } else if key_value[1]=="9"{ // 35=9; 用户相关信息
                    result.order_type = "user"
                } else if key_value[1]=="10"{ //35=10; 查询order book
                    result.order_type = "view"
                }
            case "44":
                result.price, _ = strconv.Atoi(key_value[1])
            case "11":
                result.stock_id, _ = strconv.Atoi(key_value[1])
            case "38":
                result.quantity, _ = strconv.Atoi(key_value[1])
            case "54":
                if key_value[1]=="1"{
                    result.trade_type = "buy"
                } else if key_value[1]=="2"{
                    result.trade_type = "sell"
                }
            case "5001":
                result.username = key_value[1]
            case "5002":
                result.password = key_value[1]
         }
    }
    //fmt.Println(result)
    return result
}

// Encode message from server, return a string of message.
func Encode(msg MessageTemp) string{
    var r = make([]string,15)
    i := 0
    r[i] = "35="+msg.order_type
    i++
    if msg.order_type=="trade"{
        r[i] = fmt.Sprintf("44=%d", msg.price)
        i++
        r[i] = fmt.Sprintf("11=%d", msg.stock_id)
        i++
        r[i] = fmt.Sprintf("38=%d", msg.quantity)
        i++
        if msg.trade_type == "buy"{
            r[i] = "54=1"
            i++
        } else if msg.trade_type == "sell"{
            r[i] = "54=2"
            i++
        }
    } else if msg.order_type=="user"{
        r[i] = "5001="+msg.username
        i++
        r[i] = "5002="+msg.password
        i++
    }

    var result string
    for ;i>=0;i--{
        result += r[i] + ";"
    }
    //fmt.Println(result)
    return result
}

// for test
// func main(){
//     var MessageFromServer = MessageTemp{
//         order_type : "trade",
//         trade_type : "sell",
//         quantity : 250,
//         price : 1024,
//         stock_id : 110,
//     }
//     fmt.Println(Encode(MessageFromServer))
//     fmt.Println(Decode("11=10086;38=250;54=2;5001=sixgodoo;5002=lc3018121;1=2;4=7;5=0;1=2;4=7;35=8;44=122"))
// }