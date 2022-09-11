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
	Children []Proc `json:"children",omitempty`
}

type Data struct {
	Cpu   float32 `json:"cpu",omitempty`
	Ram   Ram     `json:"ram",omitempty`
	Procs []Proc  `json:"procs",omitempty`
}

func main() {
	// for {
	cmd := exec.Command("sh", "-c", "cat /proc/stat | grep cpu | tail -1 | awk '{print ($5*100)/($2+$3+$4+$5+$6+$7+$8+$9+$10)}' | awk '{print 100-$1}'")
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])

	cmdram := exec.Command("sh", "-c", "cat /proc/ram_201906570")
	out2, err2 := cmdram.CombinedOutput()
	if err2 != nil {
		fmt.Println(err2)
	}
	output2 := string(out2[:])

	cmdcpu := exec.Command("sh", "-c", "cat /proc/cpu_201906570")
	out3, err3 := cmdcpu.CombinedOutput()
	if err3 != nil {
		fmt.Println(err3)
	}
	output3 := string(out3[:])

	jsonstring := fmt.Sprintf("{\"cpu\":%s,\"ram\":%s,%s", output, output2, output3)

	fmt.Println(jsonstring)

	var d Data
	json.Unmarshal([]byte(jsonstring), &d)

	fmt.Printf("%+v\n", d)

	// time.Sleep(2 * time.Second)
	//}
}
