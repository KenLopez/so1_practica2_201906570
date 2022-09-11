package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Ram struct {
	Totalram int `json:"totalram"`
	Freeram  int `json:"freeram"`
}

type Proc struct {
	Pid      int    `json:"pid",omitempty`
	Nombre   string `json:"nombre",omitempty`
	Usuario  int    `json:"usuario",omitempty`
	Estado   int    `json:"estado",omitempty`
	Ram      int    `json:"ram",omitempty`
	Children []int  `json:"children",omitempty`
}

type Data struct {
	Cpu   float32 `json:"cpu",omitempty`
	Ram   Ram     `json:"ram",omitempty`
	Procs []Proc  `json:"procs",omitempty`
}

func main() {
	var out [3][]byte
	var err [3]error
	var output [3]string

	// for {
	cmd := exec.Command("sh", "-c", "cat /proc/stat | grep cpu | tail -1 | awk '{print ($5*100)/($2+$3+$4+$5+$6+$7+$8+$9+$10)}' | awk '{print 100-$1}'")
	out[0], err[0] = cmd.CombinedOutput()

	if err[0] != nil {
		fmt.Println(err[0])
	}
	output[0] = string(out[0][:len(out[0])-1])

	cmdram := exec.Command("sh", "-c", "cat /proc/ram_201906570")
	out[1], err[1] = cmdram.CombinedOutput()
	if err[1] != nil {
		fmt.Println(err[1])
	}
	output[1] = string(out[1][:])

	cmdcpu := exec.Command("sh", "-c", "cat /proc/cpu_201906570")
	out[2], err[2] = cmdcpu.CombinedOutput()
	if err[2] != nil {
		fmt.Println(err[2])
	}
	output[2] = string(out[2][:])

	fmt.Println(output[2])

	jsonstring := fmt.Sprintf("{\"cpu\":%s,\"ram\":%s}", output[0], output[1])

	fmt.Println(jsonstring)

	arr := output[2]

	var d Data
	var x []Proc
	json.Unmarshal([]byte(jsonstring), &d)
	json.Unmarshal([]byte(arr), &x)

	fmt.Printf("%+v\n", d)
	fmt.Printf("%+v\n", arr)

	// time.Sleep(2 * time.Second)
	//}
}
