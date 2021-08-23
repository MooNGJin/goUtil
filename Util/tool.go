package Util

import (
	"regexp"
	"sort"
)

func IsPhone(phone string) bool {
	pre := `^(1[3|4|5|8][0-9]\d{4,8})$`
	r, err := regexp.MatchString(pre, phone); if err != nil {
		//_ = errors.New(fmt.Sprintf("read file error, %s", err))
		panic(err)
	}
	return r
}

// 定义接口，必须实现的方法
type SortMapTool interface {
	SortInt(v map[int]string) map[int]interface{}
	SortString (v map[string]string) map[string]interface{}
}

type MapSort struct {
	sortInt []int
	sortString []string
}

func (ms MapSort) SortInt(v map[int]string) map[int]interface{} {
	for k := range v {
		ms.sortInt = append(ms.sortInt, k)
	}
	sort.Ints(ms.sortInt)

	var mp = make(map[int]interface{})
	for _, k := range ms.sortInt{
		mp[k] = v[k]
	}

	return mp
}

func (ms MapSort) SortString(v map[string]string) map[string]interface{} {
	for k := range v {
		ms.sortString = append(ms.sortString, k)
	}
	sort.Strings(ms.sortString)

	var mp = make(map[string]interface{})
	for _, k := range ms.sortString{
		mp[k] = v[k]
	}

	return mp
}