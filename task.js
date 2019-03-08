var M = Object();



// A measurement ID is a unique identifier
// linking all submissions of a measurement.
M.measurementId = 1; // a UUID.
// This function embeds an image from a remote
// origin, hides it, and sets up callbacks to
// detect success or failure to load the image.
M.measure = function() 
{
var img = $('<img>');
img.attr('src', '//target/image.png');
img.style('display', 'none');
img.on('load', M.sendSuccess);
img.on('error', M.sendFailure);
img.appendTo('html');
}
// This function submits a result using
// a cross-origin AJAX request. The server
// must allow such cross-origin submissions.
M.submitToCollector = function(state) {
$.ajax({
url: "//collector/submit" +"?cmh-id=" + this.measurementId + "&cmh-result=" + state,
});
}
M.sendSuccess = function() {
M.submitToCollector("success");
}
M.sendFailure = function() {
M.submitToCollector("failure");
}
// Submit to the server as soon as the client
// loads the page, regardless of the
// measurement result. This indcates which
// clients attempted to run the measurement,
// even if they don't submit a final result.
M.submitToCollector("init");
// Run the measurement when the page loads.
$(M.measure);