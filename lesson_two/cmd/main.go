package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"os"
	ex "solvery/lesson_two/internal/string_ex"
)

func main() {
	var str string
	var daemon bool
	var pack bool
	var unpack bool

	pflag.StringVar(&str, "input", "", "string to unpack")
	pflag.BoolVar(&daemon, "daemon", false, "run in daemon mode")
	pflag.BoolVar(&pack, "pack", false, "pack mode")
	pflag.BoolVar(&unpack, "unpack", false, "unpack mode")
	pflag.Parse()

	if unpack {
		if daemon {
			daemonMode()
		} else {
			unpackMode(str)
		}
	} else if pack {
		packMode(str)
	} else {
		fmt.Println("Please specify a mode: --pack, --unpack or --daemon")
	}

}

func daemonMode() {
	for {
		fmt.Print("Введите строку: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		unpackedString, err := ex.UnpackString(input)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		fmt.Println(unpackedString)
		fmt.Println()
	}
}

func unpackMode(str string) {
	unpackedString, err := ex.UnpackString(str)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(unpackedString)
}

func packMode(str string) {
	unpackedString := ex.PackString(str)

	fmt.Println(unpackedString)
}

// LinkedList
//list := str.LinkedList[int32]{}
//fmt.Println(list.GetValues())
//list.Append(1)
//fmt.Println(list.GetValues())
//list.Append(2)
//list.Append(3)
//list.Append(4)
//list.Append(3)
//list.Append(5)
//list.Append(3)
//fmt.Println(list.GetValues())
//
//list.Prepend(4)
//list.Prepend(5)
//fmt.Println(list.GetValues())
//
//val, _ := list.RemoveTail()
//fmt.Println(val, list.GetValues())
//list.RemoveFront()
//fmt.Println(list.GetValues())
//
//val2, _ := list.FindVal(4)
//fmt.Println(val2)
//
//list.RemoveAll(3)
//fmt.Println(list.GetValues())
//
//list.Clear()
//list.Clear()
//fmt.Println(list.GetValues())
//
//val3, ok := list.RemoveTail()
//fmt.Println(val3, ok, list.GetValues())
//val4, ok := list.RemoveFront()
//fmt.Println(val4, ok, list.GetValues())
//val5, ok := list.FindVal(4)
//fmt.Println(val5, ok)
//
//list.RemoveAll(3)
//fmt.Println(list.GetValues())

//Stack
//stack := str.Stack[string]{}
//fmt.Println(stack.IsEmpty())
//stack.Push("1")
//stack.Push("2")
//stack.Push("3")
//
//fmt.Println(stack.Size())
//fmt.Println(stack.Peek())
//fmt.Println(stack.Pop())
//fmt.Println(stack.Pop())
//fmt.Println(stack.Pop())
//fmt.Println(stack.IsEmpty())
//
//stack.Push("4")
//stack.Push("5")
//stack.Push("6")
//fmt.Println(stack)
//stack.Clear()
//fmt.Println(stack.Size())

//Queue
//queue := str.Queue[float64]{}
//fmt.Println(queue.GetValues())
//queue.Push(1.1)
//queue.Push(2.2)
//queue.Push(3.3)
//fmt.Println(queue.GetValues())
//
//fmt.Println(queue.Pop())
//fmt.Println(queue.GetValues())
//
//queue.Clear()
//fmt.Println(queue.GetValues())
