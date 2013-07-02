seajs.use(['$'], function($) {
	
	var menuButton = $('.nav-menu-button'),
		nav        = $('#nav');
    menuButton.on('click', function (e) {
        nav.toggleClass('active');
    });
	
	
	$('body').on('click','.pure-paginator li a',function(e){
		e.preventDefault()
		var path = $(this).attr("href");
		$("#list").load(path+" #list .content")
	})
	$('body').on('click','.email-subject a',function(e){
		e.preventDefault()
		var path = $(this).attr("href");
		$("#main").load(path+" #main .content")
		$(".email-item").removeClass("email-item-selected");
		$(this).parent().parent().parent().addClass("email-item-selected");

	})
	
	$(".email-subject a").click(function(e){

	});


});