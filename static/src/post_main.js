define(function(require, exports, module) {
	
	//加入代码高亮  
	var $ = require('$');
	$(function(){                   //dom ready调用的另外一种方式
	  	var code =$('.artcle-cont code');
	  	if (code) {
			require.async('libs/rainbow/themes/baby-blue.css');
			seajs.use(['rainbow'],function(Rainbow){
			Rainbow.color();
			})
	  	}

	        });

	
});