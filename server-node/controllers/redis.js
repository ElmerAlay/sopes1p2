require("dotenv").config()

const redis = require("redis")
const client = redis.createClient({
    port      : process.env.REDIS_PORT,               
    host      : process.env.REDIS_HOST,
    password  : process.env.REDIS_PASSWORD
  })

const get = (req, res) => {
    let c1 = 0, c2=0, c3=0, c4=0, c5=0, c6=0, c7=0, c8=0, c9=0, c10=0, c11 = 0

    client.LRANGE(process.env.REDIS_LIST,0,-1,function(err, value) {
        if (err) { 
          throw err;
        } else {
            for(i=0; i<value.length; i++){
                value[i] = Number(value[i])
                if(value[i]<11)
                    c1++
                else if(value[i]<21 && value[i]>10)
                    c2++
                else if(value[i]<31 && value[i]>20)
                    c3++
                else if(value[i]<41 && value[i]>30)
                    c4++
                else if(value[i]<51 && value[i]>40)
                    c5++
                else if(value[i]<61 && value[i]>50)
                    c6++
                else if(value[i]<71 && value[i]>60)
                    c7++
                else if(value[i]<81 && value[i]>70)
                    c8++
                else if(value[i]<91 && value[i]>80)
                    c9++
                else if(value[i]<101 && value[i]>90)
                    c10++
                else if(value[i]>100)
                    c11++
            }

            res.json({
                menor10: c1,
                menor20: c2,
                menor30: c3,
                menor40: c4,
                menor50: c5,
                menor60: c6,
                menor70: c7,
                menor80: c8,
                menor90: c9,
                menor100: c10,
                mayor100: c11
            })
        }
    })
}

const getLast = (req, res) => {
    client.HGETALL(process.env.REDIS_HASH, (err, values)=>{
        if (err) { 
            throw err;
        } else {
            res.json(values)
        }
    })
}

module.exports = {
  get: get,
  getLast: getLast
}