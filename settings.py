#!/usr/bin/python
# -*- coding: utf-8 -*-
# author :Arion

import os, sys


# amdin 用户名 必须为 email
ADMIN_NAME = "EinMagic@gmail.com"
ADMIN_PASSWORD = "123"

#category
CATEGORIES = ["Life","Work","Rabbit-Hole"]


DEBUG = True

CACHED = True

CACHE_HOST = ["localhost:11211"]

CACHE_USER = ''

CACHE_PWD = ''

CACHE_TIMEOUT = 0   # 缓存过期时间,单位为秒, 0不过期

ROOT_PATH = os.path.join(os.path.abspath(os.path.dirname(__file__)))

STATIC_PATH = os.path.join(ROOT_PATH, 'static')

TEMPLATE_PATH = os.path.join(ROOT_PATH, 'template')

UPLOAD_DIR = "upload"

UPLOAD_PATH = os.path.join(ROOT_PATH, UPLOAD_DIR)

