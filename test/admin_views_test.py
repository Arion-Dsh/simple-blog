#!/usr/bin/env python3
#-*- coding:utf-8 -*-

from tornado.web import create_signed_value, urlencode, decode_signed_value
from tornado.escape import url_escape
from tornado.httputil import HTTPHeaders
from models.model import *
from test.base_test import BaseCase

__all__ = ['AdminViewsTest']

class AdminViewsTest (BaseCase):
    
    def setUp(self):
        super(AdminViewsTest, self).setUp()
   
    def tearDown(self):
        super(AdminViewsTest, self).tearDown()
    
    def get_headers(self):
        cookie = create_signed_value(self.app.settings['cookie_secret'],
                                    'login_user', '1234')
        headers = dict(COOKIE=cookie)
        return headers 
    
    def test_login(self):
        url = "/auth/login"
        response = self.fetch(url)
        self.assertIn(b'email', response.body)
        
        body = urlencode(dict(
            email = 'einmagic@gmail.com',
            pass_word = '12345'
        ))
        response = self.fetch(url, method="POST", body=body)
    
    def test_home_get(self):
        headers = self.get_headers()
        response = self.fetch('/admin', headers=headers)
        body = response.body.decode()
        self.assertIn('Arion. All rights reserved.', body)
    
    def home_post(self):
        body = urlencode(dict(quote_body="test quo body", quote_author= "arion"))
        response = self.fetch('/admin', method="POST", body=body)
        self.assertEqual(response.code, 200)
