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

	var articles = new Vue({
	  el: '#articles',
	  data: {
	    articles:[]
	  }
	})




	$.get('http://my-no-1.appspot.com/article', function(res) {
		articles.articles = res;
	})




});
