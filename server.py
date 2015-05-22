#!/usr/bin/env python3
#-*- coding:utf-8 -*-

import os.path
import tornado.web
import tornado.ioloop
import tornado.httpserver

from tornado.options import define, options
from views.common_handlers import ImageHandler
from views.admin_handlers import AuthLoginHandler, AuthLogoutHandler, AdminHomeHandler, \
                                 AdminArticlesHandler, AdminArticleSigleHandler, \
                                 AdminImageHandler, AdminCategoriesHandler, \
                                 AdminCategoryHandler, AdminQuotesHandler, AdminQuoteHandler, \
                                 AdminNovelCaptersHandler, AdminNovelCapterHandler, \
                                 AdminSiglePageHandler, AdminSiglePagesHandler, \
                                 AdminNovelsHandler, AdminNovelHandler
from views.front_handlers import HomeHandler, ArticlesHandler, ArticleHandler, \
                                    PageHandler, NovelChaptersHandler, NovelChapterHandler
# 链接数据库
from mongoengine import connect
db_name = 'blog_test'
try:
    from local_settings import db_name
except:
    pass
connect(db_name)

define("port", default=8000, help="run on the given port", type=int)

class Application (tornado.web.Application):

    def __init__(self):
        handlers = [
            tornado.web.URLSpec(r'/', HomeHandler, name='home'),
            tornado.web.URLSpec(r'/(zh-hans|en-us)/([0-9]+)', ArticlesHandler, name='archives'),
            tornado.web.URLSpec(r'/(zh-hans|en-us)/([\d]+-[\d]+)/([^/]+)', ArticleHandler, name="article"),
            tornado.web.URLSpec(r'/novel/([^/]+)', NovelChaptersHandler, name='novel_chapters'),
            tornado.web.URLSpec(r'/novel/([^/]+)/([\d]+)', NovelChapterHandler, name='novel_chapter'),
            tornado.web.URLSpec(r'/page/([^/]+)', PageHandler, name='sigle_page'),
            tornado.web.URLSpec(r'/auth/login', AuthLoginHandler, name='login'),
            tornado.web.URLSpec(r'/auth/logout', AuthLogoutHandler, name='logout'),
            tornado.web.URLSpec(r'/file/images/([^/]+)', ImageHandler, name='file_image'),
            tornado.web.URLSpec(r'/admin', AdminHomeHandler, name='admin_home'),
            tornado.web.URLSpec(r'/admin/articles/([\d]+)', AdminArticlesHandler, name="admin_articles"),
            tornado.web.URLSpec(r'/admin/article/add', AdminArticleSigleHandler, name="admin_article_add"),
            tornado.web.URLSpec(r'/admin/article/([\d]+)/edit', AdminArticleSigleHandler, name="admin_article_edit"),
            tornado.web.URLSpec(r'/admin/article/([\d]+)/del', AdminArticleSigleHandler, name="admin_article_del"),
            tornado.web.URLSpec(r'/admin/pages/([\d]+)', AdminSiglePagesHandler, name="admin_sigle_pages"),
            tornado.web.URLSpec(r'/admin/page/add', AdminSiglePageHandler, name="admin_sigle_page_add"),
            tornado.web.URLSpec(r'/admin/page/([^/]+)', AdminSiglePageHandler, name="admin_sigle_page"),
            tornado.web.URLSpec(r'/admin/categories', AdminCategoriesHandler, name="admin_categories"),
            tornado.web.URLSpec(r'/admin/category/([\d]+)', AdminCategoryHandler, name="admin_category"),
            tornado.web.URLSpec(r'/admin/quotes/([\d]+)', AdminQuotesHandler, name="admin_quotes"),
            tornado.web.URLSpec(r'/admin/quote/([\d]+)', AdminQuoteHandler, name="admin_quote"),
            tornado.web.URLSpec(r'/admin/novel/([\d]+)', AdminNovelHandler, name="admin_novel"),
            tornado.web.URLSpec(r'/admin/novels/([\d]+)', AdminNovelsHandler, name="admin_novels"),
            tornado.web.URLSpec(r'/admin/novel-chapters/([\d]+)', AdminNovelCaptersHandler, name="admin_novel_chapters"),
            tornado.web.URLSpec(r'/admin/novel-chapter/([\d]+)', AdminNovelCapterHandler, name="admin_novel_chapter"),
            tornado.web.URLSpec(r'/admin/novel-chapter/sigle', AdminNovelCapterHandler, name="admin_novel_chapter_sigle"),
            tornado.web.URLSpec(r'/admin/images/upload', AdminImageHandler, name="admin_image_upload"),
            tornado.web.URLSpec(r'/admin/images/([\w]+)/del', AdminImageHandler, name="admin_image_del")
        ]
        settings = dict(
            template_path=os.path.join(os.path.dirname(__file__), 'templates'),
            static_path=os.path.join(os.path.dirname(__file__), 'static'),
            xsrf_cookies=True,
            cookie_secret= u'Arion&kelly',
            login_url='/auth/login',
            debug=True,
            gzip=True,
        )
        tornado.web.Application.__init__(self, handlers, **settings)

def main():
    tornado.options.parse_command_line()
    http_server = tornado.httpserver.HTTPServer(Application())
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.instance().start()

if __name__ == '__main__':
    main()


