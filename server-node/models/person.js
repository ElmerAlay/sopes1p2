const mongoose = require("mongoose")
const Schema = mongoose.Schema

const personSchema = new Schema({
    Nombre: String,
    Departamento: String,
    Edad: Number,
    FormaContagio: String,
    Estado: String
})

module.exports = mongoose.model("persons", personSchema)