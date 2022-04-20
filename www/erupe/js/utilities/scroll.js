// Create New Variables
var time = 100; // Timer to react
var target;
var targetOffset;

// Button Top
$("#charsel_char_list_up").on("click" ,function(e){
	e.preventDefault();
	target = $('.charsel_char.active').prev('.charsel_char');
	if (target.length == 0)
		target = $('.charsel_char:last');
	scrollTo(target);
	$('.active').removeClass('active');
	target.addClass('active');
});

// Button Bottom
$("#charsel_char_list_down").on("click" ,function(e){
	e.preventDefault();
	target = $('.charsel_char.active').next('.charsel_char');
	if (target.length == 0)
		target = $('.charsel_char:first');
	scrollTo(target);
	$('.active').removeClass('active');
	target.addClass('active');
});

// Work Animation
function scrollTo(selector) {
	var offset = $(selector).offset();
    var $charsel_list = $('#charsel_list');
    $charsel_list.animate({
		scrollTop: offset.top - ($charsel_list.offset().top - $charsel_list.scrollTop())
    }, time);
}





