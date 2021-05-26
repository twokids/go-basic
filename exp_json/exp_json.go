package exp_json

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	dateFormat = "2006-01-02"
	timeFormat = "2006-01-02 15:04:05"
)

func Json_main() {
	//struct_serialize()
	//map_serialize()
	struct_unserialize()
	map_unserialize()
}

type Demo_Person struct {
	Name     string `json:"na"`
	Age      int
	Birthday time.Time
}

//结构体转json
func struct_serialize() {
	lisa := new(Demo_Person)
	lisa.Name = "lisa"
	lisa.Age = 28
	lisa.Birthday = time.Date(2019, 9, 20, 5, 28, 0, 0, time.Local)

	fmt.Println("日期格式 birthday,", lisa.Birthday) //日期格式
	str_lisa_birthday := lisa.Birthday.Format(dateFormat)

	fmt.Println("日期转字符串格式 birthday,", str_lisa_birthday)  //日期转字符串格式
	lisa.Birthday, _ = time.Parse(dateFormat, str_lisa_birthday) //字符串再转回日期

	b, err := json.Marshal(lisa)
	if err != nil {
		fmt.Println("转json异常，", err)
	}
	fmt.Println("转json后结果：", string(b)) // {"na":"lisa","Age":28,"Birthday":"2019-09-20T00:00:00Z"}

}

//map转json
func map_serialize()  {
	lisa:=make(map[string]interface{})
	lisa["Name"]="lisa"
	lisa["Age"]="22"
	lisa["Birthday"]="2021-05-22"

	b,err:=json.Marshal(lisa)
	if err != nil {
		fmt.Println("转json异常1，", err)
	}
	fmt.Println("转json后结果1：", string(b)) //{"Age":"22","Birthday":"2021-05-22","Name":"lisa"}

}


//反序列化成struct
func struct_unserialize() {
	str_lisa := `{"na":"lisa","Age":28,"Birthday":"2019-09-20T00:00:00Z"}`
	lisa:=new(Demo_Person)

	err:= json.Unmarshal([]byte(str_lisa),lisa)//lisa前可加&，也可以不加

	if err != nil {
		fmt.Println("反序列化成struct异常，", err)
	}
	fmt.Println("反序列化成struct结果：", lisa) // &{lisa 28 2019-09-20 00:00:00 +0000 UTC}

}

//反序列化成struct
func map_unserialize() {
	str_lisa := `{"Name":"lisa","Age":28,"Birthday":"2019-09-20T00:00:00Z"}`
	map_lisa:=make(map[string]interface{})

	err:= json.Unmarshal([]byte(str_lisa),&map_lisa)//map_lisa前要加&
	if err != nil {
		fmt.Println("反序列化成map异常，", err)
	}
	fmt.Println("反序列化成map结果：", map_lisa) // map[Age:28 Birthday:2019-09-20T00:00:00Z Name:lisa]

}