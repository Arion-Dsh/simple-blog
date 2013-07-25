#!/usr/bin/python
# -*- coding: utf-8 -*-
# author :Arion

import datetime
from mongotor.orm.collection import Collection as Model
from mongotor.orm.field import *

class User(Model):

    __collection__ = "user"
    _id = ObjectIdField()
    name = StringField()
    password = StringField()
    level = IntegerField()
    actived = BooleanField()
        


class Tag(Model):
    """Tag Model"""    
    __collection__ = "tags"
    _id = ObjectIdField()
    name = StringField()
    slug = StringField()

class Category(Model):
    """Tag Model"""    
    __collection__ = "category"
    _id = ObjectIdField()
    name = StringField()
    alias = StringField()
    images = ListField()
    description = StringField()

class Comment(Model):
    """Tag  Model"""
    __collection__ = "comment"
    _id = ObjectIdField()
    created_at =  DateTimeField()
    slug =  StringField()
    name = StringField()
    email = StringField()
    body = StringField()
        
class Post(Model):
    """ Posts Model"""
    __collection__ = "post"
    _id = ObjectIdField()
    #_type = StringField()
    created_at = DateTimeField()
    comment_change = DateTimeField()
    active = BooleanField()
    slug = StringField()
    title = StringField()
    slug = StringField()
    category = StringField()
    tags = ListField()
    description = StringField()
    body = StringField()
    images = ListField()
    comments = ListField()
    

    

    
    
    
    
        