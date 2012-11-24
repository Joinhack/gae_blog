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

	var showMsg = function(pos, msg) {
		var msg = $('<div class="alert">' + msg + '</div>');
			msg.appendTo($('.bottom', pos));
			msg.fadeOut(2000, function(){
				$(this).remove();
		});
	}

	var bindnewTopicDlgEvent = function(pos, dlgContent) {
		$('form[name=newTopic] .cancel', pos).click(function(){
			pos.fadeOut(2000, function(){
				pos.remove();
			})
		});

		if($('form[name=newTopic]').size() > 0) {
			$('.wrapper', pos).css('margin-top','30px');
			$('.wrapper', pos).width($(window).width() - 340);
			$('.wrapper textarea', pos).height(280);
		}

		$('form[name=newTopic] .submit', pos).click(function(){
			if($('input[name=title]', pos).val() == '') {
				showMsg(pos, "title can't be null");
				return;
			}
			$.ajaxUpload({
				form:$('form[name="newTopic"]'),
				type: 'post',
				dataType: 'json',
				success: function(data){
					alert(data);
				}
			});
		});

	}

	var bindLoginDlgEvent = function(pos, dlgContent) {
		$('form[name=login] .cancel', pos).click(function(){
			pos.fadeOut(2000, function(){
				pos.remove();
			})
		});

		$('form[name=login] .submit', pos).click(function(){
			if($('input[name=loginId]', pos).val() == '') {
				showMsg(pos, "username can't be null");
				return;
			}
			if($('input[name=password]', pos).val() == '') {
				showMsg(pos, "password can't be null");
				return;
			}

			$.ajaxUpload({
				form:$('form[name="login"]'),
				type: 'post',
				dataType: 'json',
				success: function(data){
					if(data.code != 0) {
						showMsg(pos, data.msg);
						return;
					}
					dlgContent.contents().remove();
					dlgContent.append("Loading...");
					$.getJSON('/new_topic', {}, function(data){
						if(data.code != 0) {
							showMsg(pos, data.msg);
							return;
						}
						$('wrapper', pos).css('width', '600px');
						dlgContent.contents().remove();
						var content = $(data.content);
						dlgContent.append(content);
						bindnewTopicDlgEvent(pos, dlgContent);
					});
			}});
		});
	}

	$('.newTopic').click(function(){
		var border = $('<div class="border"/>');
		var dlg = $('<div class="pdlg"/>').append(border);
		var dlgContent = $('<div class="content">Loading...</div>').appendTo(border);
		dlgContent.css("padding", "15px");
		var pos = $('<div class="positioner"/>');
		var wrapper = $('<div class="wrapper"/>').append(dlg).appendTo(pos);
		pos.appendTo($('.pops'));
		$.getJSON('/new_topic', {}, function(data){
			if(data.code != 0) {
				showMsg(pos, data.msg);
				return;
			}
			dlgContent.contents().remove();
			var content = $(data.content);
			dlgContent.append(content);
			bindLoginDlgEvent(pos, dlgContent);
			bindnewTopicDlgEvent(pos, dlgContent);
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