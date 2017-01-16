#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import base64
import tornado.web

from plus.flash import FlashMessageMixIn
# from models.model import User


class BaseHandler (tornado.web.RequestHandler, FlashMessageMixIn):

    def get_current_user(self):
        return self.get_secure_cookie('login_user')

    def get_image_b64data(self, data):
        return str(base64.b64encode(data), encoding='utf-8')
