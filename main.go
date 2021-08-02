package main

import (
	mnt "com.linxianda.test/Util"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func JSONToMap(str string) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		panic(err)
	}

	return tempMap
}




func main()  {
	/*str := "2274400358142s7jkc5ohh88vnb9pluhf133fa4900d1eefffcd3982fb8db66ea42756ea"
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str)*/

	var filePath = "xinxinUser.txt"
	//data, err :=  mnt.ReadLine(filePath); if err != nil {
	//	errors.New(fmt.Sprintf("read file error, %s", err))
	//}
	//
	//for index, value := range data {
	//	fmt.Printf("phone = %s, index = %d \n", index, value)
	//}
	//fmt.Println(len(data))
	data, err := mnt.ReadLineBySlice(filePath);if err != nil {
		_ = errors.New(fmt.Sprintf("read file error, %s", err))
	}
	for _, phone := range data{

		i, e := strconv.Atoi(phone); if e != nil {
			fmt.Println(" i = ", i)
		}
		//mnt.PhoneAol()
	}
}