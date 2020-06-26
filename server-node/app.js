require("dotenv").config()
const express = require("express");
const { json } = require("express");
const mongoose = require("mongoose")

const Routes = require("./routes/routes");

const app = express();

mongoose.connect("mongodb://"+process.env.MONGO_HOST+":"+process.env.MONGO_PORT+"/",{
   dbName: process.env.MONGO_DB,
   user: process.env.MONGO_USER,
   pass: process.env.MONGO_PASSWORD,
   useNewUrlParser: true,
   useUnifiedTopology: true
})
.then(db => console.log("Db Mongo connected!"))
.catch(err => console.log(err))

app.use(function(req, res, next) {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    next();
});
app.use(json());

app.use("/", Routes);

module.exports = app;