
      
      var CensorshipObject = new Object();
      CensorshipObject.myID = 10;
      CensorshipObject.sendSuccess = function() 
      {
        console.log("image was loaded correctly");
        // this.submitResult("success");
      }
      CensorshipObject.sendFailure = function() 
      {
        console.log("image failed to load correctly");
        // this.submitResult("failure");
      }
      CensorshipObject.measure = function() {
        var img = new Image(); // width, height values are optional params 
        img.src = 'https://www.w3schools.com/tags/smiley.gif';
        img.onerror = CensorshipObject.sendFailure;
        img.onload = CensorshipObject.sendSuccess;
        // img.css('display', 'none');   //hide the image we are loading

        document.getElementById("testingDiv").appendChild(img);

      }
      CensorshipObject.measure();//actually call the function to execute measurement task
   
    