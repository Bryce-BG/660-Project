// ==UserScript==
// @name              Censorship Measurement Script
// @namespace         https://github.com/Bryce-BG/660-Project
// @version           2.2.6
// @description       Measure Censorship by periodically polling our server for a random image on a potentially blocked site 
// @author            Bryce Bodley-Gomes
// @include           *
// @require           https://cdn.bootcss.com/jquery/1.12.4/jquery.min.js
// @run-at            document-idle
// @grant             unsafeWindow
// @grant             GM_xmlhttpRequest
// @grant             GM_download
// @grant             GM_setClipboard
// ==/UserScript==
function loadDoc() {
	var xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() 
	{
	  if (this.readyState == 4 && this.status == 200) 
	  {
		var res = document.getElementById("footer").innerHTML =this.responseText;
					 
		
		eval(document.getElementById('sc1').innerHTML); //actually register new javascript with browser

		CensorshipObject.measure(); //call the function to preform te measurement   
	  }
  };
  // TODO THIS SHOULD BE A ABSOLUTE URL NOT a relative URL because client
  xhttp.open("GET", "https://localhost:8888/task.js", true);
  xhttp.send();
}
loadDoc();