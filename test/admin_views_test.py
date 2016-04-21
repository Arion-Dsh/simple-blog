#!/usr/bin/env python3
# -*- coding:utf-8 -*-

from tornado.web import create_signed_value, urlencode
# from tornado.escape import url_escape
# from tornado.httputil import HTTPHeaders
from models.model import *
from test.base_test import BaseCase

__all__ = ['AdminViewsTest']


class AdminViewsTest (BaseCase):

    def setUp(self):
        super(AdminViewsTest, self).setUp()

    def tearDown(self):
        super(AdminViewsTest, self).tearDown()

    @property
    def headers(self):
        cookie = create_signed_value(self.app.settings['cookie_secret'],
                                     'login_user', '1234').decode()
        headers = dict(COOKIE='login_user=%s' % cookie)
        return headers

    def test_login(self):
        url = "/auth/login"
        response = self.fetch(url)
        self.assertIn(b'email', response.body)

        body = urlencode(dict(
            email='einmagic@gmail.com',
            pass_word='12345'
        ))
        response = self.fetch(url, method="POST", headers=self.headers, body=body)
        self.assertIn(b'Add Quote', response.body)

    def test_home(self):
        response = self.fetch('/admin', headers=self.headers)
        body = response.body.decode()
        self.assertIn('Add Quote', body)

        body = urlencode(dict(quote_body="test quo body",
                              quote_author="arion"))
        response = self.fetch('/admin', method="POST",
                              headers=self.headers, body=body)
        self.assertEqual(response.code, 200)

    def test_article_about(self):
        category = Category(name="zh-hans")
        category.save()
        body = urlencode(dict(
            title="sdf是地方",
            active=1,
            category='zh-hans',
            create_time='2015-06-14 16:07',
            md_content='test 1 2 3 <pre><code>def123</code></pre>',
        ))
        response = self.fetch('/admin/article/add', method="POST",
                              headers=self.headers, body=body)
        self.assertIn(b'&lt;code&gt;', response.body)
        article = Article.objects(title='sdf是地方').first()
        self.assertIn('<code>', article.html_content)

        response = self.fetch('/admin/articles/1', headers=self.headers)
        self.assertIn('sdf是地方', response.body.decode())
        self.assertIn('/zh-hans/%s/%s' % (article.create_time.strftime('%d-%m'),
                                          article.id_no),
                      response.body.decode())
        self.assertIn('/admin/article/%s/edit' % article.id_no,
                      response.body.decode())
        response = self.fetch('/admin', headers=self.headers)
        self.assertIn('sdf是地方', response.body.decode())
        self.assertIn('/admin/article/%s/edit' % article.id_no,
                      response.body.decode())

        body = urlencode(dict(
            # title="sdf是地方",
            active=1,
            category='zh-hans',
            create_time='2015-06-14 16:07',
            md_content='test 1 2 3 <pre><code>def123</code></pre>',
        ))
        response = self.fetch('/admin/article/add', method="POST",
                              headers=self.headers, body=body)
        self.assertEqual(400, response.code)

        body = urlencode(dict(
            title="sdf是地方",
            active=1,
            category='zh-hans',
            create_time='2015-06-14 24:07fc',
            md_content='test 1 2 3 <pre><code>def123</code></pre>',
        ))
        response = self.fetch('/admin/article/add', method="POST",
                              headers=self.headers, body=body)
        self.assertEqual(500, response.code)

        response = self.fetch('/admin/article/%s/del' % article.id_no,
                              method="DELETE", headers=self.headers)
        self.assertIn('1', response.body.decode())
        article = Article.objects(title='sdf是地方').first()
        self.assertEqual(1, article.is_del)

    def test_quote_about(self):
        quote = Quote(
            body="test quote 1",
            author="arion_______")
        quote.save()

        response = self.fetch('/admin/quotes/1', headers=self.headers)
        self.assertIn(b'test quote 1', response.body)
        self.assertIn(b'arion_______', response.body)

        response = self.fetch('/admin/quote/%s' % quote.id_no,
                              headers=self.headers)
        self.assertIn(b'test quote 1', response.body)
        self.assertIn(b'arion_______', response.body)

        body = urlencode(dict(
            quote_body="test quote edit",
            quote_author="arion______"
        ))
        response = self.fetch('/admin/quote/%s' % quote.id_no, method="POST",
                              headers=self.headers, body=body)
        self.assertIn(b'test quote edit', response.body)
        self.assertIn(b'arion______', response.body)

    def test_chapter_about(self):
        pass
