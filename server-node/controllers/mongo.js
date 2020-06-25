const Person = require("../models/person");

const get = async (req, res) => {
  const person = await Person.find()
  res.json(person)
}

const getTop3 = async (req, res) => {
  //agrupa los datos por departamento y cuenta cuantos hay en cada departamento
  const agg = [
    {
      $group: {
          _id: "$Departamento",
          total: { $sum: 1 }
      }
    },
    {
      $sort : { total : -1 } //lo ordena por el total en forma descendente
    },
    { $limit : 3 }  //Obtiene los 3 primeros
  ]

  //Se obtienen los datos agrupados y se muestran en formato json
  Person.aggregate(agg, function(err, logs){
    if (err) { return def.reject(err); }

    res.json(logs)
  });
}

const getPie = async (req, res) => {
  const agg = [
    {
      $group: {
          _id: "$Departamento",
          total: { $sum: 1 }
      }
    }
  ]

  Person.aggregate(agg, function(err, logs){
    if (err) { 
      return def.reject(err); 
    }

    res.json(logs)
  });
}

module.exports = {
  get: get,
  getTop3: getTop3,
  getPie: getPie
}