(function($){
	$.ajaxUpload = function(s){
		var data = {};
		var addVal = function(data, name, val) {
			if(data[name] == null)
				data[name] = val;
			else {
				if(typeof(data[name]) == 'string') {
					var vals = [];
					vals.push(data[name]);
					data[name] = vals;
					vals.push(val);
				} else {
					data[name].push(val);
				}
			}
		};
		$("input[type='text'],input[type='hidden'],input[type='password'],input:radio:checked,input:checkbox:checked,textarea", s.form).each(function() {
			var item = $(this);
			var name = item.attr('name');
			var val = item.val();
			addVal(data, name, val);
		});
		$('select option:selected').each(function(){
			addVal(data, $(this).parent().attr('name'), $(this).val());
		});
		var url = $(s.form).attr('action');
		var uploadData = $.extend({}, s.data||{}, data);
		$.ajax($.extend({url: url}, s, {data: uploadData}));
	}
})(jQuery);