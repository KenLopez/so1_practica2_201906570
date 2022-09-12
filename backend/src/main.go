package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
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

var conn = MySQLConn()

func MySQLConn() *sql.DB {
	connString := "root:123@tcp(34.134.233.102:3306)/practica2"
	conn, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Conn MySQL")
	}
	return conn
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

	jsonstring := fmt.Sprintf("{\"cpu\":%s,\"ram\":%s,\"procs\":%s}", output[0], output[1], output[2])

	var d Data
	e := json.Unmarshal([]byte(jsonstring), &d)
	if e != nil {
		fmt.Println(e)
	}

	query := `INSERT INTO LOG(fecha, cpu, ram) VALUES (NOW(),?,?);`
	result, er := conn.Exec(query, d.Cpu, float64(d.Ram.Freeram)*100/float64(d.Ram.Totalram))
	if er != nil {
		fmt.Println(er)
	}

	fmt.Printf("%+v\n", result)

	// time.Sleep(2 * time.Second)
	// }
}
