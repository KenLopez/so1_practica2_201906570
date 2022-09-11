package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("top", "-bn2 | grep '%Cpu' | tail -1 | grep", "-P '(....|...) id,'|awk '{print 100-$8 \"%\"}")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	} else {
		output := string(out[:])
		fmt.Println(output)
	}
}
