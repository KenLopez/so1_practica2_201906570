var mysql = require('mysql')
var conn = mysql.createConnection({
    host: '34.134.233.102',
    user: 'root',
    password: '123',
    database: 'practica2',
    port: '3306'
 })
conn.connect(function(err){
    if(err)throw err
    console.log("Conn Mysql")
})

module.exports = conn;