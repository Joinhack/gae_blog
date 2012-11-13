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
});
})(jQuery)