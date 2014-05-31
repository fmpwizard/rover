function moveRover(direction){
	$.ajax({
	  url: "/rover?c=" + direction + "&v=1",
	  context: document.body
	}).done(function() {
	  console.log("made it");
	});
};
