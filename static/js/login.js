(function($){

$(document).ready(function(){
	var post = function(d) {
		d = d||{};
		$.ajaxUpload({
			form:$('form'), 
			data: d,
			type: 'post',
			success: function(data){
				if(data.code != 0) {
					var msg = $('<div class="alert">' + data.msg + '</div>');
					msg.appendTo($('.bottom'));
					msg.fadeOut(2000, function(){
						$(this).remove();
					});
					return;
				}
				window.location = data.redirect;
		}});
	}

	$('.submit').click(function(){
		post();
	});


	WB2.anyWhere(function(W){
		W.widget.connectButton({
			id: "sina",	
			type:'4,4',
			callback : {
				login:function(o){
					window.location = '/'
				},
				logout:function(){
					window.location = '/'
				}
			}
		});
	});
});

})(jQuery);