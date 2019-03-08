// class someClass {

	// constructor(){
// var i;
// for(i =0; i<10; i++)
// {
// 	console.log(i);

// }
// 1. Create the button
var button = document.createElement("button");
button.innerHTML = "Do Preform Measurement";

// 2. Append somewhere
var body = document.getElementsByTagName("body")[0];
body.appendChild(button);

// 3. Add event handler
button.addEventListener ("onClick", preformMesurement());
// }



function preformMesurement() {
	console.log("need to preform a measurement");
	// 1.contact our server get task

	
	// 2. executre
	//retur result

}
// }