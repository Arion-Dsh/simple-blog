#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os.path
import tornado.web
import tornado.ioloop
import tornado.httpserver

from tornado.web import url
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
db = connect(db_name)

define("port", default=8000, help="run on the given port", type=int)


class Application (tornado.web.Application):
    handlers = [
        url(r'/', HomeHandler, name='home'),
        url(r'/(zh-hans|en-us)/([0-9]+)', ArticlesHandler, name='archives'),
        url(r'/(zh-hans|en-us)/([\d]+-[\d]+)/([^/]+)', ArticleHandler, name="article"),
        url(r'/novel/([^/]+)', NovelChaptersHandler, name='novel_chapters'),
        url(r'/novel/([^/]+)/([\d]+)', NovelChapterHandler, name='novel_chapter'),
        url(r'/page/([^/]+)', PageHandler, name='sigle_page'),
        url(r'/auth/login', AuthLoginHandler, name='login'),
        url(r'/auth/logout', AuthLogoutHandler, name='logout'),
        url(r'/file/images/([^/]+)', ImageHandler, name='file_image'),
        url(r'/admin', AdminHomeHandler, name='admin_home'),
        url(r'/admin/articles/([\d]+)', AdminArticlesHandler, name="admin_articles"),
        url(r'/admin/article/add', AdminArticleSigleHandler, name="admin_article_add"),
        url(r'/admin/article/([\d]+)/edit', AdminArticleSigleHandler, name="admin_article_edit"),
        url(r'/admin/article/([\d]+)/del', AdminArticleSigleHandler, name="admin_article_del"),
        url(r'/admin/pages/([\d]+)', AdminSiglePagesHandler, name="admin_sigle_pages"),
        url(r'/admin/page/add', AdminSiglePageHandler, name="admin_sigle_page_add"),
        url(r'/admin/page/([^/]+)', AdminSiglePageHandler, name="admin_sigle_page"),
        url(r'/admin/categories', AdminCategoriesHandler, name="admin_categories"),
        url(r'/admin/category/([\d]+)', AdminCategoryHandler, name="admin_category"),
        url(r'/admin/quotes/([\d]+)', AdminQuotesHandler, name="admin_quotes"),
        url(r'/admin/quote/([\d]+)', AdminQuoteHandler, name="admin_quote"),
        url(r'/admin/novel/([\d]+)', AdminNovelHandler, name="admin_novel"),
        url(r'/admin/novels/([\d]+)', AdminNovelsHandler, name="admin_novels"),
        url(r'/admin/novel-chapters/([\d]+)', AdminNovelCaptersHandler, name="admin_novel_chapters"),
        url(r'/admin/novel-chapter/([\d]+)', AdminNovelCapterHandler, name="admin_novel_chapter"),
        url(r'/admin/novel-chapter/sigle', AdminNovelCapterHandler, name="admin_novel_chapter_sigle"),
        url(r'/admin/images/upload', AdminImageHandler, name="admin_image_upload"),
        url(r'/admin/images/([\w]+)/del', AdminImageHandler, name="admin_image_del")
    ]
    settings = dict(
        template_path=os.path.join(os.path.dirname(__file__), 'templates'),
        static_path=os.path.join(os.path.dirname(__file__), 'static'),
        xsrf_cookies=True,
        cookie_secret='Arion&kelly',
        login_url='/auth/login',
        debug=True,
        gzip=True,
    )

    def __init__(self):
        handlers = self.handlers
        settings = self.settings
        super(Application, self).__init__(handlers, **settings)


def main():
    tornado.options.parse_command_line()
    http_server = tornado.httpserver.HTTPServer(Application())
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.instance().start()


if __name__ == '__main__':
    main()
