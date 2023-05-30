package main

import (
	"fmt"
	"time"
)

func main() {
	_exec_parallel()
}

func varibale_int(x int) {
	println(fmt.Sprintf("%d", x))
}

func println(word ...interface{}) {
	fmt.Println(word...)
}

func array() {
	arr := [3]string{"test", "test2", "test3"}
	fmt.Println(arr)
}

func variable_array() {
	arr := []string{}
	arr = append(arr, "test1")
	arr = append(arr, "test2")
	arr = append(arr, "test3")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func ensured_array() {
	arr := make([]string, 0, 1024)
	arr = append(arr, "test")
	fmt.Println(len(arr))
}

func _map() {
	mp := map[string]int{
		"x": 100,
		"y": 200,
	}
	fmt.Println(mp["x"])

	_, isExist := mp["z"]
	fmt.Println(isExist)

	mp["z"] = 300
	_, isExist = mp["z"]
	fmt.Println(isExist)
}

func _if(isOk bool) {
	if isOk {
		println("ok")
	} else {
		println("not ok")
	}
}

func _switch(mode int) {
	switch mode {
	case 1:
		fmt.Println("dev mode")
	case 2:
		fmt.Println("prod mode")
	case 3:
		fmt.Println("special mode")
		fallthrough
	case 4:
		fmt.Println("liar mode")
	}
}

func _for(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func _args(b ...int) {
	for _, num := range b {
		fmt.Println(num)
	}
}

type SUSHI_INTERFACE interface {
	ToString() string
}

type Sushi struct {
	rice string
	neta string
}

func (sushi Sushi) ToString() string {
	fmt.Println(sushi)
	return sushi.neta
}

func _pointer(a *int, b int) {
	*a++
	b++
}

func _parallel(chA chan<- string) {
	for i := 0; i < 10; i++ {
		println(i)
		time.Sleep(2 * time.Millisecond)
	}
	chA <- "Finished"
}

func _exec_parallel() {
	for i := 0; i < 5; i++ {
		println("a")
		time.Sleep(time.Millisecond)
	}
	chA := make(chan string)
	defer close(chA)
	go _parallel(chA)
	msg := <-chA
	fmt.Println(msg)

}
