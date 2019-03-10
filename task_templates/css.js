

var serverUrl = "localhost";
var measurementId = 1;

var CensorshipMeter = new Object();


CensorshipMeter.baseUrl = "{serverUrl}/submit";
CensorshipMeter.measurementId = encodeURIComponent("{measurementId}");
CensorshipMeter.submitResult = 
    function(state, message) {
        this.submitted = state; 
        var params = {
            "cmh-id": this.measurementId,
            "cmh-result": state,
        };
        if (message != null) {
            params["cmh-message"] = String(message).substring(0, this.maxMessageLength);
        }
        $.ajax({
            url: this.baseUrl + "?" + $.param(params),
        });
    }
CensorshipMeter.sendSuccess = function() {
  this.submitResult("success");
}
CensorshipMeter.sendFailure = function() {
  this.submitResult("failure");
}
CensorshipMeter.sendException = function(err) {
  this.submitResult("exception", err);
}

