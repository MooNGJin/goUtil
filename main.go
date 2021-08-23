package main

import (
	mnt "com.linxianda.test/Util"
	"encoding/json"
	_ "errors"
	"fmt"
	_ "strconv"
)

func JSONToMap(str string) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		panic(err)
	}

	return tempMap
}

func sortMap(v ...interface{}) {

	for _, arg := range v {
		switch arg.(type) { //获取参数的类型
		case int:
			fmt.Println(arg, " is an int value")
		case string:
			fmt.Println(arg, "is an string value")
		case int64:
			fmt.Println(arg, " is an int64 value")
		default:
			fmt.Println(arg, "is an unknown type.")

		}
	}

	var m = map[int]string{
		0: "java",
		1: "php",
		2: "go",
		3: "py",
	}

	var smp mnt.SortMapTool
	smp = mnt.MapSort{}
	d := smp.SortInt(m)

	for i, k := range d {
		fmt.Printf("%d == %s \n", i, k)

		fmt.Printf("%s == %T \n", k, k)
	}
}

func main() {
	/*str := "2274400358142s7jkc5ohh88vnb9pluhf133fa4900d1eefffcd3982fb8db66ea42756ea"
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str)*/

	//var filePath = "xinxinUser.txt"
	//data, err :=  mnt.ReadLine(filePath); if err != nil {
	//	errors.New(fmt.Sprintf("read file error, %s", err))
	//}
	//
	//for index, value := range data {
	//	fmt.Printf("phone = %s, index = %d \n", index, value)
	//}
	//fmt.Println(len(data))
	//data, err := mnt.ReadLineBySlice(filePath);if err != nil {
	//	_ = errors.New(fmt.Sprintf("read file error, %s", err))
	//}
	//for _, phone := range data{
	//
	//	r := mnt.IsPhone(phone)
	//	fmt.Printf("isPhone = %T \n", r)
	//
	//	i, e := strconv.Atoi(phone); if e != nil {
	//		fmt.Println("i = ", i)
	//	}
	//	//mnt.PhoneAol()
	//}

	sortMap()
}
