async function get() {
  let url = 'http://localhost:8090/temp'
  let obj = await (await fetch(url)).json();

  //console.log(obj);
  let res = ""
  for (let i = 0; i < obj.length; i++) {
    res+="name: "+obj[i].name+ " - "+ "value: "+ obj[i].value+ " - passed healthcheck: "+ obj[i].health+"<br>";
  }
  return res;
}

(async () => {
  while (true) {
    //console.log(tags)
    document.getElementById("sensors").innerHTML = await get();
    await sleep(500)
  }})()

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}
