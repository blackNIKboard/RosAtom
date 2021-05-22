async function get() {
    let url = 'http://'+location.host+'/api/temp'
    // let url = 'http://localhost:8090/temp'
    let obj = await (await fetch(url)).json();

    //console.log(obj);
    let res = ""
    for (let i = 0; i < obj.length; i++) {
        // res+="name: "+obj[i].name+ " - "+ "value: "+ obj[i].value+ " - passed healthcheck: "+ obj[i].health+"<br>";

        // res += "<div class='card text-white bg-dark mb-3' style='max-width: 60rem;'>    <div class='card-header'>       <tr>        <td>"+obj[i].name+"</td>        <td><button type='button' class='btn btn-success' style='float:right; scale: 80%;'>"+obj[i].health+"</button></td>      </tr>    </div><div class='card-body'>     Temperature, C: "+ obj[i].value+ "     <br>      Brightness, L: 50      <br>      Energy, J: 50     </div>  </div>"


        res += "<tr>" +
            "<th scope='row'>" + obj[i].name + "</th>" +
            "<td" + checkValue(obj[i].health.temperature, false) + ">" + obj[i].values.temperature + "</td>" +
            "<td" + checkValue(obj[i].health.brightness, false) + ">" + obj[i].values.brightness + "</td>" +
            "<td" + checkValue(obj[i].health.energy, false) + ">" + obj[i].values.energy + "</td>" +
            "<td" + checkValue(obj[i].health.health, true) + ">" + parseHealth(obj[i].health.health) + "</td>" +
            "</tr>"
    }
    return res;
}

function checkValue(health, ok) {
    let healthColorOk = " style='background-color: #52b788;'"
    let healthColor = ""

    let unhealthColorOk = " style='background-color: #e56b6f;'"
    let unhealthColor = " style='color: #e56b6f; font-weight: bold'"

    if (health) {
        if (ok) {
            return healthColorOk
        } else return healthColor
    } else{
        if (ok) {
            return unhealthColorOk
        } else return unhealthColor
    }
}

function parseHealth(health) {
    if (health) {
        return "OK"
    } else {
        return "Strange"
    }
}

(async () => {
    while (true) {
        //console.log(tags)
        document.getElementById("sensors").innerHTML = await get();
        await sleep(500)
    }
})()

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
