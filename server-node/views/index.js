let url = "http://localhost:3000/";

var state = {
'querySet': null,

'page': 1,
'rows': 3,
'window': 5,
}

buildTable()

function pagination(querySet, page, rows) {

var trimStart = (page - 1) * rows
var trimEnd = trimStart + rows

var trimmedData = querySet.slice(trimStart, trimEnd)

var pages = Math.round(querySet.length / rows);

return {
    'querySet': trimmedData,
    'pages': pages,
}
}

function pageButtons(pages) {
    var wrapper = document.getElementById('pagination-wrapper')

    wrapper.innerHTML = ``
    console.log('Pages:', pages)

    var maxLeft = (state.page - Math.floor(state.window / 2))
    var maxRight = (state.page + Math.floor(state.window / 2))

    if (maxLeft < 1) {
        maxLeft = 1
        maxRight = state.window
    }

    if (maxRight > pages) {
        maxLeft = pages - (state.window - 1)
        
        if (maxLeft < 1){
            maxLeft = 1
        }

        maxRight = pages
    }

    for (var page = maxLeft; page <= maxRight; page++) {
        wrapper.innerHTML += `<button value=${page} class="page btn btn-success">${page}</button>`
    }

    if (state.page != 1) {
        wrapper.innerHTML = `<button value=${1} class="page btn btn-success">&#171; First</button>` + wrapper.innerHTML
    }

    if (state.page != pages) {
        wrapper.innerHTML += `<button value=${pages} class="page btn btn-success">Last &#187;</button>`
    }

    $('.page').on('click', function() {
        $('#table-body').empty()

        state.page = Number($(this).val())

        buildTable()
    })
}


function buildTable() {
var table = $('#table-body')

axios.get(url)
        .then(function(response) {
            var data = pagination(response.data, state.page, state.rows)
            var myList = data.querySet

            for (var i = 1 in myList) {
                //Keep in mind we are using "Template Litterals to create rows"
                var row = `<tr>
                        <td>${myList[i].Nombre}</td>
                        <td>${myList[i].Edad}</td>
                        <td>${myList[i].Departamento}</td>
                        <td>${myList[i].Forma}</td>
                        <td>${myList[i].Estado}</td>
                        `
                table.append(row)
            }

            pageButtons(data.pages)

        })
        .catch(function(error) {
            console.log(error);
        })
        .then(function() {}); 

}

let topThree = function top3(){
    axios.get(url+"top3")
        .then(function(response) {
            let table2 =  $('#table-body2')
            for (let i=0; i<response.data.length; i++) {
                //Keep in mind we are using "Template Litterals to create rows"
                let row = `<tr>
                        <td>${response.data[i]._id}</td>
                        <td>${response.data[i].total}</td>
                        `
                table2.append(row)
            }
        })
        .catch(function(error) {
            console.log(error);
        })
        .then(function() {}); 
}()

let last = function() {
    axios.get(url+"last")
        .then(function(response) {
            let table3 =  document.getElementById("table-body3")
            table3.innerHTML=`<tr>
                        <td>${response.data.Nombre}</td>
                        <td>${response.data.Edad}</td>
                        <td>${response.data.Departamento}</td>
                        <td>${response.data.Forma}</td>
                        <td>${response.data.Estado}</td>
            `
            
        })
        .catch(function(error) {
            console.log(error);
        })
        .then(function() {}); 
}()

let pie = function() {
    axios.get(url+"pie")
        .then(function(response) {
            const datos = []
            const labelsD = []

            for(let i=0;i<response.data.length; i++){
                datos.push(response.data[i].total)
                labelsD.push(response.data[i]._id)
            }

            let ctx = document.getElementById("pie").getContext("2d")
            let myPieChart = new Chart(ctx, {
                type: 'pie',
                data: {
                    labels:labelsD,
                    datasets:[{
                        label: "Departamentos Infectados",
                        data: datos,
                        backgroundColor:[
                            '#FF5733',
                            '#FFE933',
                            '#78FF33',
                            '#33FFF4'
                        ]
                    }]
                },
                options: {
                    responsive: true
                }
            });

        })
        .catch(function(error) {
            console.log(error);
        })
        .then(function() {}); 
}()

let bar = function() {
    axios.get(url+"bar")
        .then(function(response) {
            let datos = [response.data.menor10,response.data.menor20,response.data.menor30,response.data.menor40,response.data.menor50,response.data.menor60,response.data.menor70,response.data.menor80,response.data.menor90,response.data.menor100,response.data.mayor100]

            let ctx = document.getElementById("bar").getContext("2d")
            let myPieChart = new Chart(ctx, {
                type: 'bar',
                data: {
                    labels:["<10","<20","<30","<40","<50","<60","<70","<80","<90","<100"],
                    datasets:[{
                        label: "Rangos",
                        data: datos,
                        backgroundColor:[
                            '#FF5733',
                            '#FFE933',
                            '#78FF33',
                            '#33FFF4',
                            "#339BFF",
                            "#FF33F5",
                            "#FF336E",
                            "#3D33FF",
                            '#FF5733',
                            '#FFE933',
                            '#33FFF4',
                        ]
                    }]
                },
                options: {
                    responsive: true
                }
            });

        })
        .catch(function(error) {
            console.log(error);
        })
        .then(function() {}); 
}()