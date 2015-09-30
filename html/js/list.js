// 
$(document).ready(function() {
	
// 	$.ajax({
// //		url: "http://my-no-1.appspot.com/list",
// 		cache: false,
// 		success: function(html)　{
// 			//$("#results").append(html);
// 			console.log(html)
// 		}
// 	});

	$.get('http://my-no-1.appspot.com/article', function(res) {
		for (var i = 0; i < res.length; i++) {
			var data = res[i];
			console.log(data);
		}
	})

});
