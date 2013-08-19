#!/usr/bin/python
# -*- coding: utf-8 -*-
# author :Arion


import os, datetime, math
from tornado import gen
import tornado.web
from core.pagination import Paginate
from core.flash import FlashMessageMixIn
from models import Post, Category, Comment, User
from forms import CommentForm, LoginForm, PostForm, CategoryForm
from settings import ADMIN_NAME, ADMIN_PASSWORD, UPLOAD_PATH, UPLOAD_DIR



class BaseHandler(tornado.web.RequestHandler, FlashMessageMixIn):


    
    def get_current_user(self):
        return self.get_secure_cookie('user')
    
class Index(BaseHandler):
       
    @tornado.web.asynchronous
    @gen.engine
    def get(self, page=1):
        posts_found = yield gen.Task(Post.objects.find,{'active': True}, sort=({"_id": -1}))
        posts = Paginate(posts_found, page, 7)
        self.render('post/list.html', posts=posts)

class PostCategory(BaseHandler):
    
       
    @tornado.web.asynchronous
    @gen.engine
    def get(self, slug, page=1):
        posts_found = yield gen.Task(Post.objects.find,{'active': True, 'category':slug}, sort=({"_id": -1}))
        posts = Paginate(posts_found, page, 7)
        categorydes = yield gen.Task(Category.objects.find_one, {'name':slug})
        self.set_secure_cookie("paginate_page", unicode(page))
        self.render('post/list.html', posts=posts, categorydes=categorydes, category=slug)

class Node(BaseHandler):
    
        
    @tornado.web.asynchronous    
    @gen.engine     
    def get(self, slug):
        node = yield gen.Task(Post.objects.find_one,{'slug':slug}) 
        self.render('post/node.html',node=node)
        
            
        


class AdminLogin(BaseHandler):
    
    @tornado.web.asynchronous
    @gen.engine
    def get(self):
        
        form = LoginForm()
        self.render('admin/login.html',form=form)
    
    @tornado.web.asynchronous
    @gen.engine
    def post(self):        
        form = LoginForm(self.request.arguments)
        user = yield gen.Task(User.objects.find_one, {'email':form.email.data})
        if form.validate():
            if form.email.data != ADMIN_NAME :
                form.email.errors.append("Provide a valid username")
            elif ADMIN_PASSWORD != form.password.data :
                form.password.errors.append("Provide a valid password")
            else:                
                self.set_secure_cookie("user", form.email.data)
                self.redirect('/admin')
        self.render('admin/login.html',form=form)

class Logout(BaseHandler):
    
    @tornado.web.asynchronous 
    def get(self):
        self.clear_cookie("user")
        self.redirect(self.get_argument('next', '/'))
            
        
class Admin(BaseHandler):
    
    @tornado.web.authenticated
    @tornado.web.asynchronous
    @gen.engine
    def get(self,page=1):
        posts_found = yield gen.Task(Post.objects.find,{},sort=({"_id": -1}) )
        posts = Paginate(posts_found, page, 13)
        self.render('admin/post_list.html', posts=posts)

            
class PostAdd(BaseHandler):
    
    @tornado.web.authenticated
    @tornado.web.asynchronous
    @gen.engine
    def get(self):
        form = PostForm()
        category_list = yield gen.Task(Category.objects.find,{})
        for category in category_list:
            form.category.choices.append((category.alias,category.alias))
        self.render('admin/post_add.html', form=form)
    
    @tornado.web.authenticated    
    @tornado.web.asynchronous
    @gen.engine
    def post(self):
        post =Post()
        form = PostForm(self.request.arguments)

        slug = form.title.data.strip(' ').replace(' ', '-')
        _post = yield gen.Task(Post.objects.find_one,{'slug':slug}) 
        if form.validate():
            if _post :
                form.title.errors.append("Provide a valid title")
            else:
                form.populate_obj(post)
                post.created_at = comment_change = datetime.datetime.now()
                post.slug = slug       
                yield gen.Task(post.save)            
                self.redirect(self.reverse_url("admin_index"))
        self.render('admin/post_add.html', form=form, post =post)
            

class PostDel(BaseHandler):
    
    @tornado.web.authenticated
    @tornado.web.asynchronous
    @gen.engine
    def get(self, slug):
        post = yield gen.Task(Post.objects.find_one,{'slug':slug})
        
        yield gen.Task(post.remove)
        self.flash("success", 'success')
        
        self.redirect(self.reverse_url("admin_index"))
        
class PostEdit(BaseHandler):
    
    @tornado.web.authenticated
    @tornado.web.asynchronous
    @gen.engine
    def get(self,slug):
        post = yield gen.Task(Post.objects.find_one,{'slug':slug})
        form = PostForm(obj=post)
        self.render('admin/post_add.html', form=form)
    
    @tornado.web.authenticated
    @tornado.web.asynchronous
    @gen.engine
    def post(self, slug):
        post = yield gen.Task(Post.objects.find_one,{'slug':slug})
        form = PostForm(self.request.arguments)
        if form.validate():
            form.populate_obj(post)
        
            yield gen.Task(post.update)
            self.flash("success", 'success')
            self.redirect(self.reverse_url("post_edit", slug))
        self.render('admin/post_add.html', form=form, post =post)


class CategoryList(BaseHandler):
    
    @tornado.web.authenticated
    @tornado.web.asynchronous
    @gen.engine
    def get(self,page=1):
        posts = yield gen.Task(Category.objects.find,{},sort=({"_id": -1}) )
        self.render('admin/category_list.html', posts=posts)

class CategoryAdd(BaseHandler):
    
    @tornado.web.authenticated
    @tornado.web.asynchronous
    @gen.engine
    def get(self):
        form = CategoryForm()
        self.render('admin/category_add.html', form = form)
    
    @tornado.web.authenticated
    @tornado.web.asynchronous
    @gen.engine
    def post(self):
        category = Category()
        form = CategoryForm(self.request.arguments)
        if form.validate():
            form.populate_obj(category)
            yield gen.Task(category.save)
        self.render('admin/category_add.html', form = form,category=category)         
        
   
        
class CategoryEdit(BaseHandler):
    
    @tornado.web.authenticated
    @tornado.web.asynchronous
    @gen.engine
    def get(self,name):
        category = yield gen.Task(Category.objects.find_one,{'name':name})
        form = CategoryForm(obj=category)
        self.render('admin/category_add.html', form=form)

    @tornado.web.authenticated
    @tornado.web.asynchronous
    @gen.engine
    def post(self, name):
        category = yield gen.Task(Category.objects.find_one,{'name':name})
        form = CategoryForm(self.request.arguments)
        if form.validate():
            form.populate_obj(category)

            yield gen.Task(category.update)
            self.flash("success", 'success')
            self.redirect(self.reverse_url("post_edit", name))
        self.render('admin/category_add.html', form=form, category =category)          
        
        

class FileHandler(tornado.web.RequestHandler):

    def post(self):
        f = self.request.files.get('file', None)
        if not f:
            self.write('{"stat": "fail", "msg": "no file"}')
            return
        f = f[0]
        filename = f.get('filename', '')
        fullpath = UPLOAD_PATH+"/"+filename
        rpath = '/' + UPLOAD_DIR + '/' + filename
        b = open(fullpath, 'w')
        b.write(f.get('body', ''))
        b.close
        self.write(filename)

class FileDel(tornado.web.RequestHandler):
 
    def get(self):
        filename = self.get_argument('filename')
        fullpath = UPLOAD_PATH+"/"+filename
        os.remove(fullpath)
        


                  
        
              


        
        
        
        