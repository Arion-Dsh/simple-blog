seajs.use(['$', 'uploader'], function($, Uploader) {
	
	function getCookie(name) {
	    var r = document.cookie.match("\\b" + name + "=([^;]*)\\b");
	    return r ? r[1] : undefined;
	}
	var xsrf = getCookie("_xsrf");
	//alert(xsrf)
    new Uploader({
        trigger: '#upload-0',
        accept: 'image/*',
        action: '/upload/',
        //data: {'_xsrf': xsrf}
    }).success(function(data) {
       	//alert(data);
        var image = '<img src="/upload/'+data+'"  rola="'+data+'">'
        $('.upload-wrapper').append(image);
		var number = $('.upload-wrapper img').size()-1
		$('.images-list').append('<input id="images-'+number+'" name="images-'+number+'" type="hidden" value="'+data+'"> ')
		 
    });


    $('.upload-wrapper').on('click','img',function(e, args){
    	var	data = $(this).attr('rola')
		var items = $('.upload-wrapper img')
		var number = $(this).index(items)
		
    	$.ajax({
		 type: "get",
		 url: "/delete/",
		 data: { filename: data, }
		}).done(function() {
			
			$('.upload-wrapper img').eq(number).remove()
			$('.images-list input').eq(number).remove()
			
		});
    })

});