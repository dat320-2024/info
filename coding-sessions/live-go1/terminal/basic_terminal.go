package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TM struct {
	cmd map[string]func()
}

func NewTM() *TM {
	return &TM{
		cmd: map[string]func(){
			"ls":   ls,
			"pwd":  pwd,
			"exit": exit,
		},
	}
}

func (tm *TM) AddCommand(name string, f func()) {
	tm.cmd[name] = f
}

func (tm *TM) Run(name string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" Basic Shell")
	fmt.Println("---------------------")
	ColorReset := "\033[0m"
	for {
		fmt.Println(string(ColorReset))
		fmt.Printf("[%s]>", name)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		//tm.Read()
		//fmt.Println("Command:", text)
		value, ok := tm.cmd[text]
		if ok {
			value()
		} else {
			//fmt.Println("Command:", text)
			fmt.Printf("Command %s not found\n", text)
		}
	}

}
func exit() {
	os.Exit(0)
}
func ls() {
	fmt.Println("** Calling func ls() **")
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}
	ColorGreen := "\033[32m"
	ColorRed := "\033[31m"
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(string(ColorGreen), file.Name())
		} else {
			fmt.Println(string(ColorRed), file.Name())
		}
	}

}
func pwd() {
	fmt.Println("** Calling func pwd() **")
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}

func (tm *TM) Read() {
	arg := os.Args[1]
	value, ok := tm.cmd[arg]
	if ok {
		value()
	} else {
		fmt.Println("Command not found")
	}
	fmt.Println(arg)

}
