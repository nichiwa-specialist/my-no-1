// 
$(document).ready(function() {

	$("#list").perfectScrollbar();

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
