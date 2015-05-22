#!/usr/bin/env python
#-*- coding:utf-8 -*-

from models.model import ImageDoc
from views import BaseHandler

class ImageHandler(BaseHandler):
    
    def get(self, id):
        img = ImageDoc.objects.get(id=id)
        if not img:
            self.flash('have not this Image', 'Error')
            return
        d = img.image.read()
        self.write(d)
        self.set_header('content-type', img.image.content_type)


