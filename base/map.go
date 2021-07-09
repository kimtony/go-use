package base


import (
	"fmt"
	"encoding/json"

)

// map 的正确声明
func NewMap() {
	m := make(map[string]int)  // map 的正确声明，分配了实际的内存
    m["one"] = 1  
	fmt.Println("--------",m["one"])
    // var m map[string]int    // 错误声明,会报错 
	// m["one"] = 1			   // error: panic: assignment to entry in nil map	
} 


//map 判断key存不存在
func MapKeyIsExist()  {
	x := map[string]string{"one": "1", "two": "2", "three": "3"}
	if _, ok := x["two"]; ok {
		// 存在
		fmt.Println("存在key",x["two"])
	}
	if _, ok := x["two"]; !ok {
		// 不存在
		fmt.Println("不存在key!")
	}
}

// jsonstr 转 map
func JsonStrToMap(){
	str := ``
	var data map[string]interface{}
	json.Unmarshal([]byte(str), &data);
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		fmt.Println("解析json字符数据报错:%v",err)
	}
}
