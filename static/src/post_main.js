define(function(require, exports, module) {
	
	//加入代码高亮  
	var $ = require('$');
	$(function(){                   //dom ready调用的另外一种方式
	  	var code =$('.email-content-body code');
	  	if (code) {
			require.async('libs/rainbow/themes/blackboard.css');
			seajs.use(['rainbow'],function(Rainbow){
			Rainbow.color();
			})
	  	}

	        });
 	//主要js
	 seajs.use(['$'], function($) {
		
		//小屏幕上菜单。
		var menuButton = $('.nav-menu-button'),
			nav        = $('#nav');
	    menuButton.on('click', function (e) {
	        nav.toggleClass('active');
	    });
		
		//所有页面 ajax
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
});

	
	
});