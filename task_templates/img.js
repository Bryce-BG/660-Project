<script id="sc1" type="text/javascript">
		var CensorshipObject = new Object();

		CensorshipObject.sendSuccess = function()
		{
		console.log("image was loaded correctly");
		CensorshipObject.submitResult("success");
		}
		CensorshipObject.sendFailure = function()
		{
		console.log("image failed to load correctly");
		CensorshipObject.submitResult("failure");
		}
		CensorshipObject.submitResult = function(message){
			var result_json = { "measurement_id" : CensorshipObject.myID,"result_status" : message};
	
			$.ajax({
				url: "/submit",
				type: "POST",
				data: JSON.stringify(result_json),
				contentType: "application/json"
			});
		}
		CensorshipObject.measure = function() {
		var img = new Image(); // width, height values are optional params
		img.src = "%s";
		img.onerror = CensorshipObject.sendFailure;
		img.onload = CensorshipObject.sendSuccess;
		// img.css('display', 'none');   //hide the image we are loading

		CensorshipObject.myID = %d;
		document.body.appendChild(img);

		}
		// CensorshipObject.measure();//actually call the function to execute measurement task
		</script>