#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import bson.py3compat
import tornado.web

from plus.markdown import markdown
from models.model import Article, Quote, User, ImageDoc, Category,\
    Novel, SiglePage, Chapter
from views import BaseHandler


class AuthLoginHandler(BaseHandler):

    def get(self):
        self.render('admin/login.html')

    def post(self):
        email = self.get_argument('email', None)
        pass_word = self.get_argument('pass_word', None)

        if not (email or pass_word):
            self.flash('validationError.', 'Error')
            return
        query = User.objects
        user = query(email=email).first()
        # 如果没有用户添加进去
        if not user and query.count() == 0:
            user = User(
                email=email,
                name=email,
                pass_word=pass_word
            )
            user.save()

        if user.pass_word != pass_word:
            self.flash('password is wrong!', 'Error')
            return
        self.set_secure_cookie('login_user', email)
        self.redirect(self.get_argument('next',
                                        self.reverse_url("admin_home")))


class AuthLogoutHandler(BaseHandler):

    def get(self):
        self.clear_cookie("login_user")
        self.redirect(self.get_argument('next', '/'))


class AdminHomeHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self):
        articles = Article.objects(is_del=False)[:5].all()
        self.render('admin/home.html', articles=articles)

    @tornado.web.authenticated
    def post(self):
        quote_body = self.get_argument('quote_body', None)
        quote_author = self.get_argument('quote_author', None)

        quote = Quote()
        quote.body = quote_body
        quote.author = quote_author

        try:
            quote.save()
        except:
            self.flash('validationError.', 'Error')
            return

        self.redirect(self.reverse_url("admin_home"))


class AdminArticlesHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self, page=1, per_page=10):
        objects = Article.objects(is_del=False).paginate(page, per_page)
        self.render('admin/articles.html', articles=objects)


class AdminArticleSigleHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self, id=None):
        article = Article()
        images = []
        if id:
            article = Article.objects.get(id_no=int(id))
            imgs = ImageDoc.objects(id__in=article.img_list).all()
            for img in imgs:
                _id = bson.py3compat.text_type(img.id)
                images.append(dict(
                    id=_id,
                    url=img.url,
                    description=img.description,
                    del_url=self.reverse_url('admin_image_del', _id)
                ))
        categories = Category.objects.all()
        self.render('admin/article.html', article=article, images=images,
                    categories=categories)

    @tornado.web.authenticated
    def post(self, id=None):

        title = self.get_argument('title')
        active = int(self.get_argument('active', 1))
        category = self.get_argument('category', '')
        create_time = self.get_argument('create_time')
        md_content = self.get_argument('md_content', '')
        translate = self.get_argument('translate', '')
        img_list = bson.json_util.loads(self.get_argument('img_list', "[]")
                                        .replace("\'", "\""))

        article = Article()
        print(create_time)
        if id:
            article = Article.objects.get(id_no=int(id))
        article.title = title
        article.active = active
        article.create_time = create_time
        article.md_content = md_content
        article.translate = translate
        article.img_list = img_list
        article.html_content = markdown(md_content)
        article.category = Category.objects(name=category).first()
        article.save()
        self.redirect(self.reverse_url('admin_article_edit', article.id_no))

    @tornado.web.authenticated
    def delete(self, id):
        article = Article.objects.get(id_no=int(id))
        article.is_del = 1
        article.save()
        self.write(dict(result=1))
        self.set_header('content-type', 'application/json')


class AdminImageHandler(BaseHandler):

    @tornado.web.authenticated
    def post(self):
        _file = self.request.files.get('file', None)
        if not _file:
            self.flash('must be have file', 'Error')
            return
        img_data = _file[0].get('body')
        description = self.get_argument('description')
        image = ImageDoc(description=description)
        image.image.put(img_data, content_type=_file[0].get('content_type'))
        image.save()
        _id = bson.py3compat.text_type(image.id)
        image.url = self.reverse_url('file_image', _id)
        image.save()
        out_data = dict(
            url=image.url,
            description=image.description,
            id=_id,
            del_url=self.reverse_url('admin_image_del', _id)
        )
        self.write(out_data)
        self.set_header('content-type', 'application/json')

    @tornado.web.authenticated
    def delete(self, id):
        image = ImageDoc.objects.get(id=id)
        image.image.delete()
        image.delete()
        self.write(dict(result=1))
        self.set_header('content-type', 'application/json')


class AdminCategoriesHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self):
        categories = Category.objects.all()
        self.render('admin/categories.html', categories=categories)

    @tornado.web.authenticated
    def post(self):
        name = self.get_argument('name')
        description = self.get_argument('description')
        category = Category(
            name=name,
            description=description
        )
        try:
            category.save()
        except:
            self.flash('validationError.', 'Error')
        self.redirect(self.reverse_url('admin_categories'))


class AdminCategoryHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self, id):
        category = Category.objects.get(id_no=int(id))
        self.render('admin/category.html', category=category)

    @tornado.web.authenticated
    def post(self, id):
        name = self.get_argument('name')
        description = self.get_argument('description')
        category = Category.objects.get(id_no=id)
        category.name = name
        category.description = description
        try:
            category.save()
        except:
            self.flash('validationError.', 'Error')
            self.redirect(self.reverse_url('admin_category', id))
            return
        self.redirect(self.reverse_url('admin_categories'))

    @tornado.web.authenticated
    def delete(self, id):
        category = Category.objects.get(id_no=id)
        category.delete()
        self.write(dict(result=1))
        self.set_header('content-type', 'application/json')


class AdminQuotesHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self, page=1, per_page=10):
        quotes = Quote.objects.paginate(page, per_page)
        self.render('admin/quotes.html', quotes=quotes)

    @tornado.web.authenticated
    def post(self, page=None):
        quote_body = self.get_argument('quote_body')
        quote_author = self.get_argument('quote_author')

        quote = Quote()
        quote.body = quote_body
        quote.author = quote_author

        try:
            quote.save()
        except:
            self.flash('validationError.', 'Error')
        self.redirect(self.reverse_url('admin_quotes', page))


class AdminQuoteHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self, id):
        quote = Quote.objects.get(id_no=id)
        self.render('admin/quote.html', quote=quote)

    @tornado.web.authenticated
    def post(self, id):
        quote_body = self.get_argument('quote_body')
        quote_author = self.get_argument('quote_author')
        quote = Quote.objects.get(id_no=id)
        quote.body = quote_body
        quote.author = quote_author
        try:
            quote.save()
        except:
            self.flash('validationError.', 'Error')
            self.redirect(self.reverse_url('admin_quote', id))
            return
        self.redirect(self.reverse_url('admin_quotes', 1))


class AdminNovelsHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self, page=1, per_page=10):
        novels = Novel.objects.paginate(page, per_page)
        self.render('admin/novels.html', novels=novels)

    @tornado.web.authenticated
    def post(self, page=None):
        name = self.get_argument('name')
        description = self.get_argument('description')

        novel = Novel()
        novel.name = name
        novel.description = description

        try:
            novel.save()
        except:
            self.flash('validationError.', 'Error')
            return
        self.redirect(self.reverse_url('admin_novels', page))


class AdminNovelHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self, id):
        novel = Novel.objects(id_no=id).first()
        self.render('admin/novel.html', novel=novel)

    @tornado.web.authenticated
    def post(self, id):
        name = self.get_argument('name')
        description = self.get_argument('description')

        novel = Novel.objects(id_no=id).first()
        novel.name = name
        novel.description = description

        try:
            novel.save()
        except:
            self.flash('validationError.', 'Error')
            self.redirect(self.reverse_url('admin_novel', id))
            return
        self.redirect(self.reverse_url('admin_novels', 1))


class AdminNovelCaptersHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self, page=1, per_page=10):
        novel_chapters = Chapter.objects.paginate(page, per_page)
        self.render('admin/novel_chapters.html', novel_chapters=novel_chapters)


class AdminNovelCapterHandler (BaseHandler):

    @tornado.web.authenticated
    def get(self, id=None):
        novel_chapter = Chapter()
        novels = Novel.objects.all()
        if id:
            novel_chapter = Chapter.objects.get(id_no=id)
        self.render('admin/novel_chapter.html', novel_chapter=novel_chapter,
                    novels=novels)

    @tornado.web.authenticated
    def post(self, id=None):
        title = self.get_argument('title')
        create_time = self.get_argument('create_time')
        active = int(self.get_argument('active', 1))
        novel = self.get_argument('novel')
        md_content = self.get_argument('md_content', '')

        novel_chapter = Chapter()
        if id:
            novel_chapter = Chapter.objects.get(id_no=id)
        novel_chapter.title = title
        novel_chapter.active = active
        novel_chapter.create_time = create_time
        novel_chapter.md_content = md_content
        novel_chapter.html_content = markdown(md_content)
        novel_chapter.novel = Novel.objects(name=novel).first()
        novel_chapter.save()
        self.redirect(self.reverse_url('admin_novel_chapter',
                                       novel_chapter.id_no))


class AdminSiglePagesHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self, page=1, per_page=10):
        sigle_pages = SiglePage.objects(is_del=0).paginate(page, per_page)
        self.render('admin/sigle_pages.html', sigle_pages=sigle_pages)


class AdminSiglePageHandler(BaseHandler):

    @tornado.web.authenticated
    def get(self, slug=None):
        sigle_page = SiglePage()
        images = []
        if slug:
            sigle_page = SiglePage.objects.get(slug=slug)
            imgs = ImageDoc.objects(id__in=sigle_page.img_list).all()
            for img in imgs:
                _id = bson.py3compat.text_type(img.id)
                images.append(dict(
                    id=_id,
                    url=img.url,
                    description=img.description,
                    del_url=self.reverse_url('admin_image_del', _id)
                ))
        self.render('admin/sigle_page.html', sigle_page=sigle_page,
                    images=images)

    @tornado.web.authenticated
    def post(self, slug=None):
        title = self.get_argument('title')
        _slug = self.get_argument('_slug', '')
        category = self.get_argument('category', '')
        create_time = self.get_argument('create_time')
        md_content = self.get_argument('md_content', '')
        translate = self.get_argument('translate', '')
        img_list = bson.json_util.loads(self.get_argument('img_list', "[]")
                                            .replace("\'", "\""))

        sigle_page = SiglePage()
        if slug:
            sigle_page = SiglePage.objects(slug=slug).first()
        sigle_page.title = title
        sigle_page.slug = slug if slug else _slug
        sigle_page.create_time = create_time
        sigle_page.category = category
        sigle_page.md_content = md_content
        sigle_page.translate = translate
        sigle_page.img_list = img_list
        sigle_page.html_content = markdown(md_content)
        try:
            sigle_page.save()
        except:
            self.flash('validationError.', 'Error')
        self.redirect(self.reverse_url('admin_sigle_page', sigle_page.slug))

    @tornado.web.authenticated
    def delete(self, slug):
        sigle_page = SiglePage.objects.get(slug=slug)
        sigle_page.is_del = 1
        sigle_page.save()
        self.write(dict(result=1))
        self.set_header('content-type', 'application/json')
