$(document).ready(function() {
    //cookies
    var Cookie = {
        set: function(name,value,expHour,domain,path){
            document.cookie = name+"="+encodeURIComponent(value==undefined?"":value)+(expHour?"; expires="+new Date(new Date().getTime()+(expHour-0)*3600000).toUTCString():"")+"; domain="+(domain?domain:document.domain)+"; path="+(path?path:"/");
        },
        get: function(name){
            return document.cookie.match(new RegExp("(^| )"+name+"=([^;]*)(;|$)"))==null ? null : decodeURIComponent(RegExp.$2);
        },
        remove: function(name){
            if(this.get(name) != null) this.set(name,null,-1);
        }
    };
    // _xsrf
    var _xsrf = Cookie.get('_xsrf');
    // 列表删除
    $('.art-del').on('click', function(e) {
        var url = $(this).attr('url-bind')
        $.ajax({
            type: "delete",
            url: url,
            data: {
                _xsrf: _xsrf
            },
            success: function(data) {
                window.location.reload()
            }})
    });
    //input img list 处理
    var setImages = function(id, del) {
        del = del ? del: false;
        var _input =  $('.article-form input.form-imgs')
        var images = _input.val() ? _input.val(): '[]';
        //坑爹的数据格式
        images = $.parseJSON(images.replace(/\'/g, "\""))
        if(!del){
            images.push(id)
        }else{
            images.shift(id)
        }
        _input.val(JSON.stringify(images))
    }
    //添加文章页面图片处理
   $('#img-form').on('submit', function(e){
        e.preventDefault()
        var $form = $(this)
        var file = $('#img-form input[name=file]')[0].files[0];
        var description = $('#img-form input[name=description]').val();
        var form_data = new FormData()
        form_data.append('file', file)
        form_data.append('description', description)
        form_data.append('_xsrf', _xsrf)
        //ajax 成功操作
        var success = function(data) {
            var html = $("<div class='pure-g' >\
                  <div class='pure-u-1-4'>\
                        <img class='pure-img' src=''>\
                    </div>\
                    <div class='pure-u-1-2'>\
                        <span class='img-description'></span><br />\
                        <span class='img-url'></span>\
                    </div>\
                    <div class='pure-u-1-4'>\
                        <button class='pure-button button-error button-small img-del' >delete</button>\
                    </div>\
                </div>")
         html.find('img').attr('src', data.url)
         html.find('.img-description').text('description'+data.description)
         html.find('.img-url').text('url:'+ data.url)
         html.find('button').attr('data-id', data.id)
         html.find('button').attr('data-del-url', data.url)
         html.addClass('id-' + data.id)
         $('.img-block').append(html)
         //id 加入 article-form
         setImages(data.id)
        }
        $.ajax({
            type: 'post',
            url: $form.attr('url-bind'),
            contentType: false,
            processData: false,
            data:form_data,
            success: success,
            dataType: 'json'
        });
    });
    // 删除图片
    $('.img-block .img-del').live('click', function(e) {
        var data_id = $(this).attr('data-id');
        var data_del_url = $(this).attr('data-del-url');
        $.ajax({
            type:'delete',
            url: data_del_url,
            data: {
                _xsrf:_xsrf
            },
            success: function (data) {
                $('.id-'+data_id).remove()
                //删除input img list
                setImages(data_id, true)
            }
        })
    });
})
