let addressElement = document.getElementById("address");
addressElement.focus();
let valueElement = document.getElementById("value");
valueElement.focus();
let typeElement = document.getElementById("typeform");

window.onPush = function () {

  let address = addressElement.value;
  let value = valueElement.value;
  let type = typeElement.type.value;

  if (address === "" || value === "" || type === "") {
    document.getElementById("text").innerHTML = "> Warn: Field is empty.";
    return;
  }

  try{
    window.go.main.App.Transmitter(address,value,type)
      .then((result) => {
        // Update result with data back from App.Greet()
        document.getElementById("text").innerHTML = `${result}`;
      })
      .catch((err) => {
        console.error(err);
      });
  } catch(err) {
    document.getElementById("text").innerHTML = `> Error! ${err} <br> Report to developer please.`;
  }

};