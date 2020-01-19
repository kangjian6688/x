package main

import (
	"fmt"
	// "github.com/astaxie/beego"
	// "github.com/bitly/go-simplejson"
	"context"
	"reflect"
	"runtime"
	"time"
)

type T struct {
	Name  string
	Value int
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello world!")
}
func Info(o interface{}) {
	t := reflect.TypeOf(o)         //反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
	fmt.Println("Type:", t.Name()) //调用t.Name方法来获取这个类型的名称

	v := reflect.ValueOf(o) //打印出所包含的字段
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ { //通过索引来取得它的所有字段，这里通过t.NumField来获取它多拥有的字段数量，同时来决定循环的次数
		f := t.Field(i)               //通过这个i作为它的索引，从0开始来取得它的字段
		val := v.Field(i).Interface() //通过interface方法来取出这个字段所对应的值
		fmt.Printf("%6s:%v =%v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ { //这里同样通过t.NumMethod来获取它拥有的方法的数量，来决定循环的次数
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)

	}
}

type Kind uint

func rotine(cancel func()) {
	fmt.Println("start rotine")
	time.Sleep(5 * time.Second)
	fmt.Println("will end rotine")
	cancel()
	fmt.Println("end rotine")
}

func tc() {
	ctx, cancel := context.WithCancel(context.Background())

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	fmt.Println(time.Now())
	// defer cancel()
	go rotine(cancel)

	select {
	case <-time.After(6 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("dddddd", ctx.Err())
	}
	fmt.Println(time.Now())
}

type Chan struct {
	done  chan int
	value string
}

func change(ch *Chan) {
	ch.value = "test"
	time.Sleep(3 * time.Second)
	p := 1
	ch.done <- p
}
func callChange() {
	c := Chan{done: make(chan int), value: ""}
	fmt.Println("start change", time.Now())
	go change(&c)
	select {
	case <-c.done:
		fmt.Println("change done")
	}
	fmt.Println("done change", time.Now(), c)
}

func main() {
	maxCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(maxCPU)
	u := User{1, "Jack", 23}
	Info(u)

	// tc()
	callChange()
}

/*
func main() {
	maxCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(maxCPU)
	// beego.Run()

	t := T{"test", 1}
	reflectType := reflect.TypeOf(t)
	reflectValue := reflect.ValueOf(t)

	fmt.Println("TypeOf", reflectType)
	fmt.Println("TypeOfName", reflectType.Name())
	fmt.Println("TypeOfString", reflectType.String())
	fmt.Println("TypeOfPath", reflectType.PkgPath())
	fmt.Println("TypeOfNumField", reflectType.NumField())
	p := []int{1}
	fmt.Println("TypeOfElem", reflect.TypeOf(p).Elem())
	fmt.Println("TypeOfKind", reflectType.Kind())
	fmt.Println("ValueOf", reflectValue)
	fmt.Println("Value Kind", reflectValue.Kind())
	fmt.Println("Value NumField", reflectValue.NumField())
	// reflectValue.Set
	fmt.Println("Value", t)
}

*/
