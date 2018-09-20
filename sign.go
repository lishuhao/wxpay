package wxpay

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func signReq(req interface{},signKey string) (sign string,err error) {
	t:=reflect.TypeOf(req)
	v:=reflect.ValueOf(req)
	keys:=make([]string,0)
	reqMap:=make(map[string]string,0)


	for i := 0; i < v.NumField(); i++ {
		key :=t.Field(i).Tag.Get("xml")
		if key=="xml"{
			continue
		}

		if strings.Contains(key,","){
			key = strings.Split(key,",")[0]
		}
		val:=v.Field(i).String()
		if val == ""{
			continue
		}
		keys = append(keys,key)
		reqMap[key] = val
	}

	sort.Strings(keys)

	reqSlice:=make([]string,0)
	for _,k:=range keys{
		reqSlice = append(reqSlice,k+"="+reqMap[k])
	}
	originStr := strings.Join(reqSlice,"&")+"&key="+signKey
	fmt.Println(originStr)

	md5Byte:=md5.Sum([]byte(originStr))
	md5Str:=hex.EncodeToString(md5Byte[:])
	return strings.ToUpper(md5Str),nil
}
