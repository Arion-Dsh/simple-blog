# -*- coding: utf-8 -*-
# join the path
import os, sys
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

import datetime
import tornado.ioloop
import tornado.web
from tornado import gen
from mongotor.orm import collection 
from mongotor.orm.field import *
#from mongotor.orm.types.compound import *
#from mongotor.orm.model.models import Model
from mongotor.database import Database
from tools.pagination import Paginate
from wtforms.validators import ValidationError, email, required
from wtforms import *
from tools.form import Form
from tools.field import TagListField
from blog.tools.tomako import MakoTemplateLoader

import tornado.httpserver
from tornado.options import define, options

define("port", default=8000, help="run on the given port", type=int)

# A connection to the MongoDB database needs to be
# established before perform operations
Database.connect('localhost:27017', 'test_db')

class Tag(collection.Collection):
    __collection__ = "user"

    _id = ObjectIdField()
    name =StringField()
    des = StringField()

class User(collection.Collection):

    __collection__ = "user"
    _id = ObjectIdField()
    name = StringField()
    active = BooleanField()
    created = DateTimeField()
    
class Post(collection.Collection):
    """docstring for Post"""
    __collection__ = "post"
    _id = ObjectIdField()
    a = StringField()
    b = ListField()
    c = BooleanField()
 
class LoginForm(Form):
    
    email = TextField('Email', validators=[required(),email('must be a Email!')])
    password =PasswordField()
    
    @tornado.web.asynchronous
    @gen.engine
    def get_user(self):
        yield gen.Task(User.objects.find_one, {'email':self.email.data})
    
    def validate_login(self, field):
        user = self.get_user()
        
        if (user is None) or (self.email.data != ADMIN_NAME):
            raise ValidationError("have no this user")
    

class MainHandler(tornado.web.RequestHandler):
    
    @tornado.web.asynchronous
    def get(self):
        form = LoginForm()
        self.render('1.html',form=form)


class Index(tornado.web.RequestHandler):

    def on_count_response(self, count):
        self.write(count)
        self.finish()
    @tornado.web.asynchronous
    @gen.engine
    def get(self):
        user = User()
        #user.name = 'sdfsdaf'
        #user.active = True
        #user.created = datetime.now()

        #yield gen.Task(user.save)
        
        form = LoginForm()
        tag = Tag()
        tag.name ='absdf'
        tag.des = 'sdfjsdjflsdf'
#         user.tags.append(tag)
        
        #yield gen.Task(user.save)
        post = Post()
        post.a='2321314'
        post.b=['s','sdf']
        post.c = True
        #post.save()
        #post.tags.append(tag.as_dict())
        #post._id = "12"
        yield gen.Task(post.save)
        #user_found = yield gen.Task(Post.objects.find_one)
        
        
        #user_found.name ='wq'
        #yield gen.Task( user_found.remove)
        
        # update date
        # yield gen.Task(user_found.update,{'$push':{
#             'tags':tag.serialize()
#         }})
        #yield gen.Task(user_found.update)
        # find one object
        #user_found = yield gen.Task(Post.objects.find,{'$or':[{'tags.name':{"$regex":"中"}},{'name':{"$regex":"ds"}}]}, sort=[('_id', 'ASC')]   )
       # user_found = yield gen.Task(Post.objects.find,{}, sort=({"_id": -1}),limit=2)        
       
        #user_found = Post.objects.count(callback=self.on_count_response)
        user_found= yield gen.Task(Post.objects.find,{})
        
        #posts = user_found[0:1]
        #paginate = Paginate(user_found, 1, 2)
        #posts = paginate._list
        #ps = paginate.iter_pages
        #self.write(user_found.name)
        # find many objects
       # new_user = User()
        #new_user.name = "new user name"
        #new_user.user.active = True
        #new_user.created = datetime.now()
        #yield gen.Task(user_found.remove)
        #users_actives = yield gen.Task(User.objects.find,{'active': True}, skip=5)
        #for a in users_actives:
            #self.write(a.name)
        self.render("1.html", title="My title", posts=user_found)
        #users_actives[0].active = False
       # yield gen.Task(users_actives[0].save)

        # remove object
       



class Application(tornado.web.Application):
    """docstring for Application"""
    def __init__(self):
        """docstring for __init__"""
        handlers = [
            tornado.web.URLSpec("/", Index),
            tornado.web.URLSpec("/add",MainHandler),
            #tornado.web.URLSpec(r'/coo', MainHandler)
        ]
        settings = {
            'template_loader': MakoTemplateLoader(os.path.join(os.path.dirname(__file__), "templates")),
           # 'static_path': os.path.join(os.path.dirname(__file__), "static"),
            'debug':True,
            "cookie_secret": "bZJc2sWbQLKos6GkHn/VB9oXwQt8S0R0kRvJ5/xJ89E=",
            
        }
        tornado.web.Application.__init__(self, handlers, **settings)


if __name__ == "__main__":
	tornado.options.parse_command_line()
	http_server = tornado.httpserver.HTTPServer(Application())
	http_server.listen(options.port)
	tornado.ioloop.IOLoop.instance().start()
