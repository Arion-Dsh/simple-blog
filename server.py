#!/usr/bin/python
# -*- coding: utf-8 -*-
# author :Arion

# join the path
import os, sys

import datetime
import tornado.ioloop
import tornado.web
from mongotor.database import Database
import tornado.httpserver
from tornado.options import define, options
from core.tomako import MakoTemplateLoader

define("port", default=8000, help="run on the given port", type=int)
Database.connect('localhost:27017', 'Arion-blog-db')

from handlers import (AdminLogin, Logout, Admin, CategoryAdd, CategoryList, CategoryEdit,
                            PostAdd, PostEdit, PostDel, Index, PostCategory, Node, FileHandler)





class Application(tornado.web.Application):
    """docstring for Application"""
    def __init__(self):
        """docstring for __init__"""
        handlers = [
            tornado.web.URLSpec(r"/upload/", FileHandler),
            tornado.web.URLSpec(r"/logout", Logout, name="logout"),
            tornado.web.URLSpec(r"/admin/login", AdminLogin, name="admin_login"),
            tornado.web.URLSpec(r"/admin", Admin, name="admin_index"),            
            tornado.web.URLSpec(r"/admin/page-([0-9]+)/", Admin, name="admin"),
            tornado.web.URLSpec(r"/admin/category", CategoryList, name ="admin_category_list"),
            tornado.web.URLSpec(r"/admin/category/add", CategoryAdd, name ="category_add"),
            tornado.web.URLSpec(r"/admin/category/edit/(.*)", CategoryEdit, name ="category_edit"),
            tornado.web.URLSpec(r"/admin/post/add/", PostAdd, name="post_add"),
            tornado.web.URLSpec(r"/admin/post/edit/(.*)", PostEdit, name="post_edit"),
            tornado.web.URLSpec(r"/admin/post/del/(.*)", PostDel, name="post_del"),
            tornado.web.URLSpec(r"/post/(.*)", Node, name="post"),
            tornado.web.URLSpec(r"/", Index, name = "index"),
            tornado.web.URLSpec(r"/pages/([0-9]+)/", Index, name="page_list"),
            tornado.web.URLSpec(r"/(.*)/", PostCategory, name="category"),
            tornado.web.URLSpec(r"/(.*)/([0-9]+)", PostCategory ,name="category_list"),
            (r"/src/(.*)", tornado.web.StaticFileHandler, {'path': 'static/src/'}),
            (r'/upload/(.*)', tornado.web.StaticFileHandler, {'path': 'upload/'})
            
        ]
        settings = {
            'template_loader': MakoTemplateLoader(os.path.join(os.path.dirname(__file__), "templates")),
            'static_path': os.path.join(os.path.dirname(__file__), "static"),
            'login_url': "/admin/login",
            'debug':True,
            "cookie_secret": "bZJc2sWbQLKos6GkHn/VB9oXwQt8S0R0kRvJ5/xJ89E=",
            'gzip': True,
            
        }
        tornado.web.Application.__init__(self, handlers, **settings)


if __name__ == "__main__":
	tornado.options.parse_command_line()
	http_server = tornado.httpserver.HTTPServer(Application())
	http_server.listen(options.port)
	tornado.ioloop.IOLoop.instance().start()
