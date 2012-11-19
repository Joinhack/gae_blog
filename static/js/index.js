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

	var bindLoginDlgEvent = function(pos) {
		$('.cancel', pos).click(function(){
			pos.fadeOut(2000, function(){
				pos.remove();
			})
		});

		$('.submit', pos).click(function(){
			var showMsg = function(msg) {
				var msg = $('<div class="alert">' + msg + '</div>');
					msg.appendTo($('.bottom', pos));
					msg.fadeOut(2000, function(){
						$(this).remove();
				});
			}
			if($('input[name=loginId]', pos).val() == '') {
				showMsg("username can't be null");
				return;
			}
			if($('input[name=password]', pos).val() == '') {
				showMsg("password can't be null");
				return;
			}

			$.ajaxUpload({
				form:$('form[name="pop"]'),
				type: 'post',
				dataType: 'json',
				success: function(data){
					if(data.code != 0) {
						showMsg(data.msg);
						return;
					}
			}});
		});
	}

	$('.newTopic').click(function(){
		var border = $('<div class="border"/>');
		var dlg = $('<div class="pdlg"/>').append(border);
		var title = $('<div class="content">Loading...</div>').appendTo(border);
		title.css("padding", "15px");
		var pos = $('<div class="positioner"/>');
		var wrapper = $('<div class="wrapper"/>').append(dlg).appendTo(pos);
		pos.appendTo($('.pops'));
		$.getJSON('/new_topic', {}, function(data){
			if(data.code != 0) {
				showMsg(data.msg);
				return;
			}
			title.contents().remove();
			var content = $(data.content);
			title.append(content);
			bindLoginDlgEvent(pos);
		});

	});

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