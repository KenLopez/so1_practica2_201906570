package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("sh", "-c", "cat /proc/stat | grep cpu | tail -1 | awk '{print ($5*100)/($2+$3+$4+$5+$6+$7+$8+$9+$10)}' | awk '{print 100-$1}'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	} else {
		output := string(out[:])
		fmt.Println(output)
	}
}
