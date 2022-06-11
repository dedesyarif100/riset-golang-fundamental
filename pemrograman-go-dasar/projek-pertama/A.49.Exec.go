package main

import (
    "fmt"
	"os/exec"
)

func main() {
	// A.49.1. Penggunaan Exec
	fmt.Println("# - A.49.1. Penggunaan Exec");
		var output1, _ = exec.Command("ls").Output()
		fmt.Printf(" -> ls\n%s\n", output1)

		var output2, _ = exec.Command("pwd").Output()
		fmt.Print("OUTPUT :", string(output2))
		fmt.Printf(" -> pwd\n%s\n", output2)

		var output3, _ = exec.Command("git", "config", "user.name").Output()
		fmt.Print("OUTPUT :", string(output3))
		fmt.Printf(" -> git config user.name\n%s\n", output3)

		// var output4

	// A.49.2. Rekomendasi Penggunaan Exec
	fmt.Println("# - A.49.2. Rekomendasi Penggunaan Exec");
		// if runtime.GOOS == "windows" {
		// 	output, err = exec.Command("cmd", "/C", "git config user.name").Output()
		// } else {
		// 	output, err = exec.Command("bash", "-c", "git config user.name").Output()
		// }

	// A.49.3. Method Exec Lainnya
	fmt.Println("# - A.49.3. Method Exec Lainnya");
	// https://golang.org/pkg/os/exec/
}