package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s student) Study() string {
	msg := "good good study, day day up!"
	fmt.Println(msg)
	return msg
}
func (s student) Sleep() string {
	msg := "good good sleep, quick quick up!"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
func reflectValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a)
	reflectValue(b)

	c := reflect.ValueOf(a)
	fmt.Printf("type c :%T\n", c)

	//想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。
	//reflectValue1(b)

	reflectValue2(&b)
	fmt.Println(b)

	//*int类型指针
	var d *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(d).IsNil())
	//nil 值
	fmt.Println("nil IsValid:", reflect.ValueOf(d).IsValid())
	//实例化一个匿名结构体
	e := struct{}{}
	//尝试从结构体中查找abc 字段
	fmt.Println("不存在的结构体成员: ", reflect.ValueOf(e).FieldByName("abc").IsValid())
	//尝试从结构体中查找abc方法
	fmt.Println("不存在的结构体方法: ", reflect.ValueOf(e).MethodByName("abc").IsValid())
	//map
	f := map[string]int{}
	//尝试从map中查找一个不存在的键
	fmt.Println("map不存在的键:", reflect.ValueOf(f).MapIndex(reflect.ValueOf("娜扎")).IsValid())
	fmt.Println("=============================结构体字段信息============================================")
	stu1 := student{
		Name:  "小王子",
		Score: 19,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v  json tag:%v\n", field.Name, field.Index, field.Type, field.Tag)
	}

	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag)
	}
	fmt.Println("=============================结构体方法信息============================================")
	printMethod(stu1)
}
