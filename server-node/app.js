const express = require("express");
const { json } = require("express");
const mongoose = require("mongoose")

const Routes = require("./routes/routes");

const app = express();

mongoose.connect("mongodb://35.192.186.35:27017/",{
   dbName: "persons",
   useNewUrlParser: true,
   useUnifiedTopology: true
})
.then(db => console.log("Db connected!"))
.catch(err => console.log(err))

app.use(function(req, res, next) {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    next();
});
app.use(json());

app.use("/", Routes);

module.exports = app;