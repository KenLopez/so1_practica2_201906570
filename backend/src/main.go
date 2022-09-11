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

	// arr := "[{\"pid\":1,\"nombre\":\"systemd\",\"usuario\":0,\"estado\":1,\"ram\":11,\"children\":[170,201,327,449,453,488,501,506,510,521,525,533,546,549,722,728,730,821,828,831,893,998,1013,1165]},{\"pid\":2,\"nombre\":\"kthreadd\",\"usuario\":0,\"estado\":1,\"children\":[3,4,5,7,9,10,11,12,13,14,15,16,18,19,20,21,22,24,25,26,27,29,30,31,32,33,34,80,81,82,83,84,85,86,87,88,90,92,93,95,96,97,98,99,100,101,102,104,105,115,118,119,126,127,128,131,224,323,324,325,326,843,1806,3068,5591,6466,7701,8550]},{\"pid\":3,\"nombre\":\"rcu_gp\",\"usuario\":0,\"estado\":1026,\"children\":[]},{\"pid\":4,\"nombre\":\"rcu_par_gp\",\"usuario\":0,\"estado\":1026,\"children\":[]}]"
	arr := output[2]

	var d Data
	var x []Proc
	json.Unmarshal([]byte(jsonstring), &d)
	e := json.Unmarshal([]byte(arr), &x)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Printf("%+v\n", d)
	fmt.Printf("%+v\n", x)

	// time.Sleep(2 * time.Second)
	//}
}
