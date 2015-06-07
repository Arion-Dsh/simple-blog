#!/usr/bin/env python3
#-*- coding:utf-8 -*-

import bson.json_util
import random
import tornado.web

from plus.pagination import Paginate
from models.model import Article, Quote, ImageDoc, Category, Novel, Chapter, SiglePage
from views import BaseHandler


class FEBaseHandler(BaseHandler):
    
    def get_random_quote(self):
        _quote = Quote.objects
        quote = _quote(id_no=random.randrange(_quote.count() or 1)).first()
        return quote or _quote.first()


class HomeHandler(FEBaseHandler):

    def get(self):
        page = SiglePage.objects(slug='home').first()
        quote = self.get_random_quote()
        image = ImageDoc()
        if page:
            image = ImageDoc.objects(id__in=page.img_list).first()
        category = Category.objects(name='zh-hans').first()
        articles = Article.objects(category=category, active=1)[:5].all()
        self.render('home.html', page=page, articles=articles, image=image, quote=quote)


class ArticlesHandler(FEBaseHandler):

    def get(self, category, page=1, per_page=10):
        category = Category.objects(name=category).first()
        articles=Paginate(Article.objects(category=category, active=1), page, per_page)
        quote = self.get_random_quote()
        self.render('list.html', category=category, articles=articles, quote=quote)


class ArticleHandler(FEBaseHandler):

    def get(self, category, ime_line, id):
        article = Article.objects(id_no=id).first()
        quote = self.get_random_quote()
        self.render('page.html', article=article, quote=quote)


class PageHandler(FEBaseHandler):

    def get(self, slug):
        page = SiglePage.objects(slug=slug).first()
        quote = self.get_random_quote()
        self.render('page.html', page=page, quote=quote)


class NovelChaptersHandler(FEBaseHandler):

    def get(self, name):
        novel = Novel.objects(name=name).first()
        chapters = Chapter.objects(novel=novel).all()
        self.render('novel_chapters.html', novel=novel, chapters=chapters)


class NovelChapterHandler(FEBaseHandler):

    def get(self, novel, id):
        quote = self.get_random_quote()
        novel = Novel.objects(name=novel).first()
        chapter = Chapter.objects(id_no=int(id), novel=novel).first()
        prev_chapter = Chapter.objects(id_no__lt=int(id), novel=novel).first()
        next_chapter = Chapter.objects(id_no__gt=int(id), novel=novel).first()
        self.render('novel_chapter.html', chapter = chapter, prev_chapter=prev_chapter,
                    next_chapter=next_chapter, quote=quote)
        
