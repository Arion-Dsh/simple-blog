#!/usr/bin/env python
# -*- coding:utf-8 -*-

from tornado.escape import url_escape
from models.model import *
from test.base_test import BaseCase

__all__ = ['FrontendViewsTest']


class FrontendViewsTest(BaseCase):

    def setUp(self):
        super(FrontendViewsTest, self).setUp()

        self.home_page = SiglePage(slug="home",
                                   title="home_page",
                                   md_content="test",
                                   html_content="<p>test</p>")
        self.home_page.save()

        self.category_1 = Category(name="zh-hans",
                                   description="just chinese.")
        self.category_2 = Category(name="en-us",
                                   description="english category.")
        self.category_1.save()
        self.category_2.save()
        self.article_1 = Article(category=self.category_1,
                                 title="test article 1",
                                 md_content="article 中国",
                                 html_content="<p>article 中国</p>")
        self.article_1.save()
        self.article_2 = Article(category=self.category_1,
                                 title="test article 2",
                                 md_content="article 中国",
                                 html_content="<p>article 中国</p>",
                                 active=0)
        self.article_2.save()

        self.article_3 = Article(category=self.category_2,
                                 title="test article english",
                                 md_content="article english",
                                 html_content="<p>article english</p>"
                                 )
        self.article_3.save()

        self.novel = Novel(name="江湖")
        self.novel.save()

        self.chapter_1 = Chapter(title="chapter test 1",
                                 md_content="中文小说test",
                                 html_content="<p>中文小说test</p>",
                                 novel=self.novel)
        self.chapter_2 = Chapter(title="chapter test 2",
                                 md_content="中文小说test2",
                                 html_content="<p>中文小说test 2</p>",
                                 novel=self.novel,
                                 active=0)
        self.chapter_1.save()
        self.chapter_2.save()

    def tearDown(self):
        super(FrontendViewsTest, self).tearDown()

    def test_homepage(self):
        response = self.fetch('/')
        self.assertIn(b'<p>test</p>', response.body)
        self.assertIn(b'test article 1', response.body)
        self.assertNotIn(b'test article 2', response.body)
        self.assertNotIn(b'test article english', response.body)

    def test_zh_hans_articles(self):
        url = '/%s/1' % self.category_1.name
        response = self.fetch(url)
        self.assertIn(b'test article 1', response.body)
        self.assertNotIn(b'test article 2', response.body)
        self.assertNotIn(b'test article english', response.body)

    def test_zh_hans_articles_1(self):
        url = '/%s/2' % self.category_1.name
        response = self.fetch(url)
        self.assertIn(b'<p>no article yet!</p>', response.body)
        self.assertNotIn(b'test article 1', response.body)

    def test_en_us_articles(self):
        url = '/%s/1' % self.category_2.name
        response = self.fetch(url)
        body = response.body.decode()
        self.assertNotIn('test article 1', body)
        self.assertIn('test article english', body)
        self.assertIn('%s' % self.article_3.create_time.strftime('%d-%m'),
                      body)

    def test_article(self):
        url = '/zh-hans/%s/%s' % (self.article_1.create_time\
                                  .strftime('%d-%m'), self.article_1.id_no)
        response = self.fetch(url)
        body = response.body.decode()
        self.assertIn('<p>article 中国</p>', body)

    def test_novel_page(self):
        url = '/novel/%s' % url_escape(self.novel.name)
        response = self.fetch(url)
        body = response.body.decode()
        self.assertIn('chapter test 1', body)
        self.assertNotIn('chapter test 2', body)
        self.assertIn('%s' % self.chapter_1.id_no, body)

    def test_chapter(self):
        url = '/novel/%s/%s' % (url_escape(self.novel.name),
                                self.chapter_1.id_no)
        response = self.fetch(url)
        body = response.body.decode()
        self.assertIn('<p>中文小说test</p>', body)
