(function($){
var resizeW = function(){
	var w = $(window).width();
	$('.row .left').width(w*0.65);
	$('.row .right').width(w*0.25);
};
$(window).resize(function(){
	resizeW();
});
$(document).ready(function(){
	resizeW();

	var showMsg = function(msg) {
		alert(msg);
	}

	$('button.login').click(function(){
		if($('.title-right input[name=loginId]').val() == '') {
			showMsg("username can't be null");
			return;
		}
		if($('.title-right input[name=password]').val() == '') {
			showMsg("password can't be null");
			return;
		}
		$.ajaxUpload({
			form:$('form'), 
			type: 'post',
			dataType: 'json',
			success: function(data){
				if(data.code != 0) {
					showMsg(data.msg);
					return;
				}
				$('.title-right').children().remove();
				var div = $("<div>Welcome " + data.loginId + "</div>");
				div.append($("<span class='btn  btn-primary logout'>Logout</span>"));
				$('.title-right').append(div);
		}});
	});
});
})(jQuery)