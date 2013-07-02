seajs.use(['$', 'uploader'], function($, Uploader) {
  	
	function upload (value){
	    new Uploader({
	        trigger: '#upload-'+value,
			accept: 'image/*',
	        action: '/upload/'
	    }).success(function(data) {
			var wrapper = $(".upload-img-list-"+value);
			var next = value + 1;
			if (value<4) {
				wrapper.after("<div class='upload-img-list-"+next+"'><button class='pure-button pure-button-success' id='upload-"+next+"'>more</button></div>");
			}			
			wrapper.append("<img src='"+ data +"' /><i class='icon-remove' path='"+data+"'></i>");			
			wrapper.find(".pure-button").addClass("hidden");
			var form =$('form[target*=iframe-uploader]').eq(value)
			form.addClass('hidden')
			$(".images-list").append("<input id='images-"+value+"' name='images-"+value+"' type='text' value='"+data+ "'>")
    
	    });	
	};
	$(".upload-wrapper").on('mouseenter','button',function(e){
		var length= $('.pure-button-success').length
		upload(length-1)
	})

});