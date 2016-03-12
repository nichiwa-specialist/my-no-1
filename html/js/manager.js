//
$(document).ready(function() {

	$('#detail').trigger('autoresize');

	$('select').material_select();

	$("#post").on("touchend mouseup", function() {
		var url = "/article/";

		var data = {
			"title" : $("#title").val(),
			"lat" : "35.039419",
			"lng" : "135.7915279",
			"detail" : $("#detail").val()
		};

		console.log(JSON.stringify(data));

		$.ajax({
			type: "POST",
			url: url,
			data: JSON.stringify(data),
			contentType: 'application/JSON',
			dataType: 'JSON',
			scriptCharset: 'utf-8',
			success: function(data) {
				alert("投稿が完了しました");
			},
			error: function(data) {
				console.log(data);
				alert("投稿が失敗しました");
			}
		});
	});

});
