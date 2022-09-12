const express = require('express');
const router = express.Router();
var conn = require('../db')

router.get('/', (req, res) => {
    conn.query(
        `SELECT *FROM LOG l ORDER BY l.id DESC LIMIT 2;`,
        (err, result) => {
            if(err){ throw err }
            if(result.length>1){
                res.send(result[1])
            }else{
                res.send(result[0])
            }
        }
    )
})

router.get('/process', async (req, res) => {
    conn.query(
        `WITH logid AS (SELECT (MAX(id)-1) AS id FROM LOG)SELECT p.id, p.pid, p.nombre, p.usuario, p.estado, p.ram, p.padre FROM PROCESO p, logid WHERE p.log = logid.id;`,
        (err, result) => {
            if(err){ throw err }
            let processes = [];
            let children = [];
            let exec = 0;
            let sus = 0;
            let stp = 0;
            let zmb = 0;
            let total = result.length;
            for (const proc of result) {
                switch (proc.estado) {
                    case 'EN EJECUCION':
                        exec++;
                        break;
                    case 'SUSPENDIDO':
                        sus++;
                        break;
                    case 'DETERNIDO':
                        stp++;
                        break;
                    case 'ZOMBIE':
                        zmb++;
                    default:
                        zmb++;
                        break;
                }
                if(proc.padre == null){
                    processes.push({...proc, children: []})
                }
                else{
                    children.push({...proc});
                } 
            }
            for (const proc of children) {
                const parent = processes.find(item => item.id === proc.padre);
                parent.children.push({...proc});
            }
            res.send({
                ejecucion: exec,
                suspendidos: sus,
                detenidos: stp,
                zombie: zmb,
                total: total,
                procs: processes
            });
        }
    )
})

module.exports = router;
