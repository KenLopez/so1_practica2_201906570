package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Parent struct {
	Value    *Proc
	Children []*Proc
}

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
	var err [5]error
	var output [3]string

	for {
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
		result, er := conn.Exec(query, d.Cpu, float64(d.Ram.Totalram-d.Ram.Freeram)*100/float64(d.Ram.Totalram))
		if er != nil {
			fmt.Println(er)
		}

		logId, _ := result.LastInsertId()

		var parents []*Parent

		for i := 0; i < len(d.Procs); i++ {
			p := d.Procs[i]
			if len(p.Children) > 0 {
				var parent Parent
				parent.Value = &p
				for j := 0; j < len(p.Children); j++ {
					c := p.Children[j]
					for k := 0; k < len(d.Procs); k++ {
						pr := d.Procs[k]
						if pr.Pid == c {
							parent.Children = append(parent.Children, &pr)
						}
					}

				}
				parents = append(parents, &parent)
			}
		}

		for i := 0; i < len(parents); i++ {
			pro := *parents[i].Value
			query := `INSERT INTO PROCESO(pid, nombre, usuario, estado, ram, log) VALUES (?,?,?,?,?,?);`
			state := ""
			switch pro.Estado {
			case 0:
				state = "EN EJECUCION"
			case 1:
				state = "SUSPENDIDO"
			case 2:
				state = "SUSPENDIDO"
			case 4:
				state = "DETENIDO"
			case 32:
				state = "ZOMBIE"
			default:
				state = "SUSPENDIDO"
			}
			res, er2 := conn.Exec(query, pro.Pid, pro.Nombre, pro.Usuario, state, float64(pro.Ram)/float64(d.Ram.Totalram), logId)
			if er2 != nil {
				fmt.Println(er2)
			}
			parentId, _ := res.LastInsertId()
			for j := 0; j < len(parents[i].Children); j++ {
				query2 := `INSERT INTO PROCESO(pid, nombre, usuario, estado, ram, padre, log) VALUES (?,?,?,?,?,?,?);`
				ch := *parents[i].Children[j]
				chState := ""
				switch ch.Estado {
				case 0:
					chState = "EN EJECUCION"
				case 1:
					chState = "SUSPENDIDO"
				case 2:
					chState = "SUSPENDIDO"
				case 4:
					chState = "DETENIDO"
				case 8:
					chState = "ZOMBIE"
				default:
					chState = "SUSPENDIDO"
				}
				_, er3 := conn.Exec(query2, ch.Pid, ch.Nombre, ch.Usuario, chState, float64(ch.Ram)/float64(d.Ram.Totalram), parentId, logId)
				if er3 != nil {
					fmt.Println(er3)
				}
			}

		}
		fmt.Println(logId)

		time.Sleep(1 * time.Second)
	}
}
