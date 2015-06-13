#!/usr/bin/env python3
#-*- coding:utf-8 -*-


from tornado.testing import AsyncTestCase, AsyncHTTPTestCase
from mongoengine import connect
from server import Application, db_name

class BaseCase(AsyncHTTPTestCase):
    
    def setUp(self):
        super(BaseCase, self).setUp()
        self.db = connect(db_name)
        self.db.drop_database(db_name)
    
    def get_app(self):
        self.app = Application
        self.app.settings['xsrf_cookies'] = False
        return self.app()
    
    def tearDown(self):
        self.db.drop_database(db_name)
        super(BaseCase, self).tearDown()
        

