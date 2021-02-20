package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	test6()
}

func test6() {
	input, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(input)
	output := fmt.Sprintf("back: %s", input)
	err = clipboard.WriteAll(output)
	if err != nil {
		panic(err)
	}
}

func test5() {
	input := os.Args[1]
	fmt.Println(input)
	// for _, i := range input {
	// 	fmt.Println(i)
	// }
}

func test4() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func test3() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Println("input=", input)
		}
	}
}

// func test2(r io.Reader) {
// 	scanner := bufio.NewScanner(bufio.NewReader(r))
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }

func test1() {
	fmt.Println("play")
	cmd := toCommand(os.Args[1:])
	fmt.Println("cmd=", cmd)
}

func toCommand(args []string) string {
	cmd := ""
	// cmd = strings.Join(args, " ")
	for _, arg := range args {
		fmt.Println("==>>", arg)
	}
	return cmd
}
