package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	// A.49.1. Penggunaan Exec
	// output1, err := exec.Command("ls").Output()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("-> ls\n%s\n", string(output1))

	// output2, err := exec.Command("pwd").Output()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("-> pwd\n%s\n", string(output2))

	// output3, err := exec.Command("git", "config", "user.name").Output()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("-> git config user.name\n%s\n", string(output3))

	// A.49.2. Rekomendasi Penggunaan Exec
	var output []byte
	var err error
	if runtime.GOOS == "windows" {
		output, err = exec.Command("cmd", "/C", "git config user.name").Output()
	} else {
		output, err = exec.Command("bash", "-c", "git config user.name").Output()
	}

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Output: %s", output)
}
