function moveRover(direction){
	$.ajax({
	  url: "http://cameraserver:7070/rover?c=" + direction + "&v=1",
	  context: document.body
	}).done(function() {
	  alert("made it");
	});
};
