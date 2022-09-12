const express = require('express');
const data = require('./routes/data');
const cors = require("cors");

const app = express();
const port = 3000;



app.use(express.json());
app.use(cors());

app.use('/data', data);

app.listen(port, ()=>{
    console.log(`App listening on port ${port}`)
})

module.exports = app;
