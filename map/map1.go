package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {

	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	//创建了一个初始存储能力 为100的map
	myMap := make(map[string]PersonInfo, 100)
	myMap = map[string]PersonInfo{
		"666": {"2", "Jeen", "beijing"},
		"777": {"3", "Micro", "shenzhen"},
	}

	//往这个map里插入几条数据
	personDB["12345"] = PersonInfo{"12345", "zoolye", "hangzhou"}
	personDB["1"] = PersonInfo{"1", "Jack", "shanghai"}

	//从这个map查找键为"12345"的信息
	//ok是一个返回的bool型，返回true表示找到了对应的数据
	person, ok := personDB["12345"]

	if ok {
		fmt.Println(person)
	} else {
		fmt.Println("没有找到相应的数据")
	}

	info, myMapOk := myMap["666"]

	if myMapOk {
		fmt.Println(myMap)
		delete(myMap, "666")
		fmt.Println(myMap)
	} else {
		fmt.Println(info)
	}

}
