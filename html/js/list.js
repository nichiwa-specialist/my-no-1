//
$(document).ready(function() {

	$("#articles").perfectScrollbar();
	$('.parallax').parallax();

	var articles = new Vue({
	  el: '#articles',
	  data: {
	    articles:[]
	  }
	})

	$.getJSON('/article/', function(res) {
		articles.articles = res;
	})
});
