/**
 *  ueditor完整配置项
 *  可以在这里配置整个编辑器的特性
 */

(function () {
    //这里你可以配置成ueditor目录在您网站中与根目录之间的相对路径或者绝对路径（指以http开头的绝对路径）
    //window.UEDITOR_HOME_URL可以在外部配置，这里就不用配置了
    //场景：如果你有多个页面使用同一目录下的editor,因为路径不同，所以在使用editor的页面上配置这个路径写在这个js外边
    //var URL = window.UEDITOR_HOME_URL || '../';
    //var tmp = window.location.pathname,
    //        URL = window.UEDITOR_HOME_URL||tmp.substr(0,tmp.lastIndexOf("\/")+1).replace("_examples/","").replace("website/","");//这里你可以配置成ueditor目录在您网站的相对路径或者绝对路径（指以http开头的绝对路径）
	var URL = window.UEDITOR_HOME_URL || '/static/admin/js/ueditor/';//修改路径
	
	UEDITOR_CONFIG = {
        imagePath:URL + "server/upload/", //图片文件夹所在的路径，用于显示时修正后台返回的图片url！具体图片保存路径需要在后台设置。！important
        compressSide:0,                   //等比压缩的基准，确定maxImageSideLength参数的参照对象。0为按照最长边，1为按照宽度，2为按照高度
        maxImageSideLength:900,          //上传图片最大允许的边长，超过会自动等比缩放,不缩放就设置一个比较大的值
        relativePath:true,                //是否开启相对路径。开启状态下所有本地图片的路径都将以相对路径形式进行保存.强烈建议开启！

        filePath:URL + "server/upload/",  //附件文件夹保存路径
	    catchRemoteImageEnable:false,                                   //是否开启远程图片抓取
        catcherUrl:URL +"server/submit/php/getRemoteImage.php",             //处理远程图片抓取的地址
        localDomain:"baidu.com",                                        //本地顶级域名，当开启远程图片抓取时，除此之外的所有其它域名下的图片都将被抓取到本地
	    imageManagerPath:URL + "server/submit/php/imageManager.php",       //图片在线浏览的处理地址
        UEDITOR_HOME_URL:URL,                                          //为editor添加一个全局路径
        //工具栏上的所有的功能按钮和下拉框，可以在new编辑器的实例时选择自己需要的从新定义
        toolbars:[
                   ['Bold','Underline','StrikeThrough','|','InsertUnorderedList','InsertOrderedList','BlockQuote','|','Link','Unlink','|','Source']
        ],
        //当鼠标放在工具栏上时显示的tooltip提示
        labelMap:{
		'bold':'加粗','underline':'下划线','strikethrough':'删除线','insertunorderedlist':'无序列表','insertorderedlist':'有序列表','blockquote':'引用','link':'超链接','unlink':'取消链接','source':'源码'
        },
        //dialog内容的路径 ～会被替换成URL
        iframeUrlMap:{
            'anchor':'~/dialogs/anchor/anchor.html',
            'insertimage':'~/dialogs/image/image.html',
            'inserttable':'~/dialogs/table/table.html',
            'link':'~/dialogs/link/loftlink.html',
            'spechars':'~/dialogs/spechars/spechars.html',
            'searchreplace':'~/dialogs/searchreplace/searchreplace.html',
            'map':'~/dialogs/map/map.html',
            'gmap':'~/dialogs/gmap/gmap.html',
            'insertvideo':'~/dialogs/video/video.html',
            'help':'~/dialogs/help/help.html',
            'highlightcode':'~/dialogs/code/code.html',
            'emotion':'~/dialogs/emotion/emotion.html',
            'wordimage':'~/dialogs/wordimage/wordimage.html',
            'attachment':'~/dialogs/attachment/attachment.html',
            'insertframe':'~/dialogs/insertframe/insertframe.html',
            'edittd':'~/dialogs/table/edittd.html',
            'snapscreen': '~/dialogs/snapscreen/snapscreen.html'
        },
        //所有的的下拉框显示的内容
        listMap:{
            //字体
            'fontfamily':['宋体', '楷体', '隶书', '黑体', 'andale mono', 'arial', 'arial black', 'comic sans ms', 'impact', 'times new roman'],
            //字号
            'fontsize':[10, 11, 12, 14, 16, 18, 20, 24, 36],
            //段落格式 值:显示的名字
            'paragraph':['p:段落', 'h1:标题 1', 'h2:标题 2', 'h3:标题 3', 'h4:标题 4', 'h5:标题 5', 'h6:标题 6'],
            //段间距 值和显示的名字相同
            'rowspacing':['5', '10', '15', '20', '25'],
            //行内间距 值和显示的名字相同
            'lineheight':['1', '1.5','1.75','2', '3', '4', '5'],
            //block的元素是依据设置段落的逻辑设置的，inline的元素依据BIU的逻辑设置
            //尽量使用一些常用的标签
            //参数说明
            //tag 使用的标签名字
            //label 显示的名字也是用来标识不同类型的标识符，注意这个值每个要不同，
            //style 添加的样式
            //每一个对象就是一个自定义的样式
            'customstyle':[
                {tag:'h1', label:'居中标题', style:'border-bottom:#ccc 2px solid;padding:0 4px 0 0;text-align:center;margin:0 0 20px 0;'},
                {tag:'h1', label:'居左标题', style:'border-bottom:#ccc 2px solid;padding:0 4px 0 0;margin:0 0 10px 0;'},
                {tag:'span', label:'强调', style:'font-style:italic;font-weight:bold;color:#000'},
                {tag:'span', label:'明显强调', style:'font-style:italic;font-weight:bold;color:rgb(51, 153, 204)'}
            ]
        },
        //字体对应的style值
        fontMap:{
            '宋体':['宋体', 'SimSun'],
            '楷体':['楷体', '楷体_GB2312', 'SimKai'],
            '黑体':['黑体', 'SimHei'],
            '隶书':['隶书', 'SimLi'],
            'andale mono':['andale mono'],
            'arial':['arial', 'helvetica', 'sans-serif'],
            'arial black':['arial black', 'avant garde'],
            'comic sans ms':['comic sans ms'],
            'impact':['impact', 'chicago'],
            'times new roman':['times new roman']
        },
        //定义了右键菜单的内容
//        contextMenu:[
//            {
//                label:'删除',
//                cmdName:'delete'
//
//            },
//            {
//                label:'全选',
//                cmdName:'selectall'
//
//            },
//            {
//                label:'删除代码',
//                cmdName:'highlightcode',
//                icon:'deletehighlightcode'
//
//            },
//            {
//                label:'清空文档',
//                cmdName:'cleardoc',
//                exec:function () {
//
//                    if ( confirm( '确定清空文档吗？' ) ) {
//
//                        this.execCommand( 'cleardoc' );
//                    }
//                }
//            },
//            '-',
//            {
//                label:'取消链接',
//                cmdName:'unlink'
//            },
//            '-',
//            {
//                group:'段落格式',
//                icon:'justifyjustify',
//
//                subMenu:[
//                    {
//                        label:'居左对齐',
//                        cmdName:'justify',
//                        value:'left'
//                    },
//                    {
//                        label:'居右对齐',
//                        cmdName:'justify',
//                        value:'right'
//                    },
//                    {
//                        label:'居中对齐',
//                        cmdName:'justify',
//                        value:'center'
//                    },
//                    {
//                        label:'两端对齐',
//                        cmdName:'justify',
//                        value:'justify'
//                    }
//                ]
//            },
//            '-',
//            {
//                label:'表格属性',
//                cmdName:'edittable',
//                exec:function () {
//                    this.tableDialog.open();
//                }
//            },
//            {
//                label:'单元格属性',
//                cmdName:'edittd',
//                exec:function () {
//
//                    this.ui._dialogs['tdDialog'].open();
//                }
//            },
//            {
//                group:'表格',
//                icon:'table',
//
//                subMenu:[
//                    {
//                        label:'删除表格',
//                        cmdName:'deletetable'
//                    },
//                    {
//                        label:'表格前插行',
//                        cmdName:'insertparagraphbeforetable'
//                    },
//                    '-',
//                    {
//                        label:'删除行',
//                        cmdName:'deleterow'
//                    },
//                    {
//                        label:'删除列',
//                        cmdName:'deletecol'
//                    },
//                    '-',
//                    {
//                        label:'前插入行',
//                        cmdName:'insertrow'
//                    },
//                    {
//                        label:'前插入列',
//                        cmdName:'insertcol'
//                    },
//                    '-',
//                    {
//                        label:'右合并单元格',
//                        cmdName:'mergeright'
//                    },
//                    {
//                        label:'下合并单元格',
//                        cmdName:'mergedown'
//                    },
//                    '-',
//                    {
//                        label:'拆分成行',
//                        cmdName:'splittorows'
//                    },
//                    {
//                        label:'拆分成列',
//                        cmdName:'splittocols'
//                    },
//                    {
//                        label:'合并多个单元格',
//                        cmdName:'mergecells'
//                    },
//                    {
//                        label:'完全拆分单元格',
//                        cmdName:'splittocells'
//                    }
//                ]
//            },
//            {
//                label:'复制(ctrl+c)',
//                cmdName:'copy',
//                exec:function () {
//                    alert( "请使用ctrl+c进行复制" );
//                }
//            },
//            {
//                label:'粘贴(ctrl+v)',
//                cmdName:'paste',
//                exec:function () {
//                    alert( "请使用ctrl+v进行粘贴" );
//                }
//            }
//        ],

        initialStyle://编辑器内部样式
        	'.selectTdClass{background-color:#3399FF !important}'+
        	'table{margin-bottom:10px;border-collapse:collapse;}td{padding:2px;}'+
        	'.pagebreak{display:block;clear:both !important;cursor:default !important;width: 100% !important;margin:0;}'+
        	'.anchorclass{background: url("' + URL + 'themes/default/images/anchor.gif") no-repeat scroll left center transparent;border: 1px dotted #0000FF;cursor: auto;display: inline-block;height: 16px;width: 15px;}' +
        	'.view{padding:0;word-wrap:break-word;word-break:break-all;cursor:text;height:100%;}' +
        	'html{background:#fff url("http://l.bst.126.net/rsc/img/shadow-in.png") no-repeat;}' +
        	'body{margin:8px;font-family:"Hiragino Sans GB","Hiragino Sans GB W3","Microsoft YaHei","微软雅黑",tahoma,arial,simsun,"宋体";font-size:14px;color:#333;text-align:left;}' +
        	'li{clear:both}' +
        	'*{line-height:1.4;}' +
        	'a{color:#7594b3;}' +
        	'p{margin:10px 0}' +
        	'ol,ul{margin:10px 0 10px 40px;padding:0;}'+
        	'ol li p,ul li p{margin:0;}'+
        	'blockquote{padding:0 0 0 10px;margin:10px 0 10px 10px;border-left:3px solid #ccc;}',
        //初始化编辑器的内容,也可以通过textarea/script给值，看官网例子
        initialContent:'',
        autoClearinitialContent:false, //是否自动清除编辑器初始内容，注意：如果focus属性设置为true,这个也为真，那么编辑器一上来就会触发导致初始化的内容看不到了
        iframeCssUrl:'', //要引入css的url
        removeFormatTags:'b,big,code,del,dfn,em,font,i,ins,kbd,q,samp,small,span,strike,strong,sub,sup,tt,u,var', //清除格式删除的标签
        removeFormatAttributes:'class,style,lang,width,height,align,hspace,valign', //清除格式删除的属性
        enterTag:'p', //编辑器回车标签。p或br
        maxUndoCount:20, //最多可以回退的次数
        maxInputCount:20, //当输入的字符数超过该值时，保存一次现场
        selectedTdClass:'selectTdClass', //设定选中td的样式名称
        pasteplain:false, //是否纯文本粘贴。false为不使用纯文本粘贴，true为使用纯文本粘贴
        //提交表单时，服务器获取编辑器提交内容的所用的参数，多实例时可以给容器name属性，会将name给定的值最为每个实例的键值，不用每次实例化的时候都设置这个值
        textarea:'editorValue',
        focus:false, //初始化时，是否让编辑器获得焦点true或false
        indentValue:'2em', //初始化时，首行缩进距离
        pageBreakTag:'_baidu_page_break_tag_', //分页符
        minFrameHeight:'', //最小高度
        autoHeightEnabled:false, //是否自动长高
        autoFloatEnabled:false, //是否保持toolbar的位置不动
        elementPathEnabled:false, //是否启用elementPath
        wordCount:false, //是否开启字数统计
        maximumWords:10000, //允许的最大字符数
        tabSize:4, //tab的宽度
        tabNode:'&nbsp;', //tab时的单一字符
        imagePopup:true, //图片操作的浮层开关，默认打开
        emotionLocalization:false, //是否开启表情本地化，默认关闭。若要开启请确保emotion文件夹下包含官网提供的images表情文件夹
        sourceEditor:"codemirror", //源码的查看方式，codemirror 是代码高亮，textarea是文本框
        tdHeight:'20', //单元格的默认高度
        highlightJsUrl:URL + "third-party/SyntaxHighlighter/shCore.js",
        highlightCssUrl:URL + "third-party/SyntaxHighlighter/shCoreDefault.css",
        codeMirrorJsUrl:URL + "third-party/codemirror2.15/codemirror.js",
        codeMirrorCssUrl:URL + "third-party/codemirror2.15/codemirror.css",
        zIndex : 99, //编辑器z-index的基数
        fullscreen : false, //是否上来就是全屏
        snapscreenHost: '127.0.0.1', //屏幕截图的server端文件所在的网站地址或者ip，请不要加http://
        snapscreenServerFile: URL +"server/upload/php/snapImgUp.php", //屏幕截图的server端保存程序，UEditor的范例代码为“URL +"server/upload/php/snapImgUp.php"”
        snapscreenServerPort: 80,//屏幕截图的server端端口
        snapscreenImgAlign: 'center', //截图的图片默认的排版方式
        snapscreenImgIsUseImagePath: 1, //是否使用上面定义的imagepath，如果为否，那么server端需要直接返回图片的完整路径
        messages:{
            pasteMsg:'编辑器已过滤掉您粘贴内容中不支持的格式！', //粘贴提示
            wordCountMsg:'当前已输入 {#count} 个字符，您还可以输入{#leave} 个字符 ', //字数统计提示，{#count}代表当前字数，{#leave}代表还可以输入多少字符数。
            wordOverFlowMsg:'你输入的字符个数已经超出最大允许值，服务器可能会拒绝保存！', //超出字数限制
            pasteWordImgMsg:'您粘贴的内容中包含本地图片，需要转存后才能正确显示！',
            snapScreenNotIETip: '截图功能需要在ie浏览器下使用',
            snapScreenMsg:'截图上传失败，请检查你的PHP环境。 '
        },
        serialize:function () {                              //配置过滤标签
        	return {
        		blackList: {table:1,style:1,script:1,link:1,applet:1,input:1,meta:1,base:1,button:1,select:1,textarea:1,'#comment':1,'map':1,'area':1},

        		//白名单，编辑器会根据此配置保留对应标签下的对应标签或者属性
        		whiteList:{
        			        'div':{'$':{},'div':1,'p':1,'a':1,'span':1,'ul':1,'ol':1,'li':1,'embed':1,'object':1,'br':1,'BR':1,'u':1,'b':1,'strong':1,'strike':1,'img':1,'blockquote':1},
        			        'a':{'$':{'href':1,'target':1},'span':1,'embed':1,'object':1,'br':1,'BR':1,'u':1,'b':1,'strong':1,'strike':1,'img':1},
        		            'p': {'$':{'reblogfrom':1},'div':1,'p':1,'a':1,'span':1,'ul':1,'ol':1,'li':1,'embed':1,'object':1,'br':1,'BR':1,'u':1,'b':1,'strong':1,'strike':1,'img':1,'blockquote':1},
        		            'span': {'$':{'style':1},'a':1,'span':1,'embed':1,'object':1,'br':1,'BR':1,'u':1,'b':1,'strong':1,'strike':1,'img':1},
        		            'ul':{'$':{},'li':1},
        		            'ol':{'$':{},'li':1},
        		            'li':{'$':{},'div':1,'p':1,'a':1,'span':1,'ul':1,'ol':1,'li':1,'embed':1,'object':1,'br':1,'BR':1,'u':1,'b':1,'strong':1,'strike':1,'img':1,'blockquote':1},
        		            'embed':{'$':{'id':1,'name':1,'width':1,'height':1,'src':1,'type':1,'wmode':1,'quality':1,'pluginspage':1,'flashvars':1,'allowscriptaccess':1}},
        		            'object':{'$':{},'param':1,'embed':1,'$':{'id':1,'name':1,'width':1,'height':1,'type':1,'classid':1,'codebase':1}},
        		            'param':{'$':{'name':1,'value':1}},
        		            'br':{'$':{}},
        		            'u':{'$':{},'a':1,'span':1,'embed':1,'object':1,'br':1,'BR':1,'u':1,'b':1,'strong':1,'strike':1,'img':1},
        		            'b':{'$':{},'a':1,'span':1,'embed':1,'object':1,'br':1,'BR':1,'u':1,'b':1,'strong':1,'strike':1,'img':1},
        		            'strong':{'$':{},'a':1,'span':1,'embed':1,'object':1,'br':1,'BR':1,'u':1,'b':1,'strong':1,'strike':1,'img':1},
        		            'strike':{'$':{},'a':1,'span':1,'embed':1,'object':1,'br':1,'BR':1,'u':1,'b':1,'strong':1,'strike':1,'img':1},
        		            'img':{'$':{'src':1,'smallsrc':1}},
        		            'blockquote':{'$':{},'div':1,'p':1,'a':1,'span':1,'ul':1,'ol':1,'li':1,'embed':1,'object':1,'br':1,'BR':1,'u':1,'b':1,'strong':1,'strike':1,'img':1,'blockquote':1}
        		            
        		           }
        		};
        }(),
        //下来框默认显示的内容
        ComboxInitial:{
            FONT_FAMILY:'字体',
            FONT_SIZE:'字号',
            PARAGRAPH:'段落格式',
            CUSTOMSTYLE:'自定义样式'
        },
        //自动排版参数
        autotypeset:{
            mergeEmptyline : true,          //合并空行
            removeClass : true,            //去掉冗余的class
            removeEmptyline : false,        //去掉空行
            textAlign : "left",             //段落的排版方式，可以是 left,right,center,justify 去掉这个属性表示不执行排版
            imageBlockLine : 'center',      //图片的浮动方式，独占一行剧中,左右浮动，默认: center,left,right,none 去掉这个属性表示不执行排版
            pasteFilter : true,             //根据规则过滤没事粘贴进来的内容
            clearFontSize : false,           //去掉所有的内嵌字号，使用编辑器默认的字号
            clearFontFamily : false,         //去掉所有的内嵌字体，使用编辑器默认的字体
            removeEmptyNode : false,         // 去掉空节点
            //可以去掉的标签
            removeTagNames : {div:1,a:1,abbr:1,acronym:1,address:1,b:1,bdo:1,big:1,cite:1,code:1,del:1,dfn:1,em:1,font:1,i:1,ins:1,label:1,kbd:1,q:1,s:1,samp:1,small:1,span:1,strike:1,strong:1,sub:1,sup:1,tt:1,u:1,'var':1},
            indent : false,                  // 行首缩进
            indentValue : '2em'             //行首缩进的大小
        }
    };
})();
