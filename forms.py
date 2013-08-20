#!/usr/bin/python
# -*- coding: utf-8 -*-
# author :Arion

import tornado.web
from tornado import gen
from wtforms import *
from core.form import Form
from core.field import TagListField 
from wtforms.validators import ValidationError, email, required


from settings import ADMIN_NAME, ADMIN_PASSWORD



class LoginForm(Form):
    
    email = TextField('Email', validators=[required(),email('must be a Email!')])
    password =PasswordField()
    
class CommentForm(Form):
    
    name = TextField()
    email = TextField()
    body = TextField()

class CategoryForm(Form):
    
    name = TextField()
    alias = TextField()
    images = FieldList(TextField())
    description = TextAreaField()
        
        

class PostForm(Form):
    title = TextField(validators=[required()])
    category = SelectField(choices=[])
    tags = TagListField()    
    active = SelectField(choices=[('True', 'Ture'), ('False', 'False')])
    description = TextAreaField(validators=[required()])
    images = FieldList(HiddenField())
    body = TextAreaField(validators=[required()])
    
    
    
    