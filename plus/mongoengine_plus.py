# !/usr/bin/env python3
# -*- coding:utf-8 -*-

import math
from mongoengine import Document as Doc
from mongoengine.queryset import QuerySet


class QueryPaginateError(Exception):
    """docstring for QueryPaginateError"""
    pass


class BaseQuerySet(QuerySet):
    """docstring for BaseQuerySet"""

    def paginate(self, page=1, per_page=10):
        page = int(page)
        per_page = int(per_page)
        if page < 1:
            page = 1
        items = self.limit(per_page).skip((page-1)*per_page).all()
        if page == 1 and len(items) < per_page:
            total = len(items)
        else:
            total = self.count()
        return Paginate(self, page, per_page, items, total)


class Document(Doc):
    """docstring for Document"""
    meta = {'abstract': True,
            'queryset_class': BaseQuerySet}


class Paginate(object):

    def __init__(self, query, page, per_page, items, total):
        """docstring for __init__"""
        self.page = int(page)
        self.per_page = int(per_page)
        self.total = total
        self.query = query
        self._list = items

    @property
    def pages(self):
        """The total number of pages"""
        return int(math.ceil(self.total / float(self.per_page)))

    def prev(self):
        return self.query.paginate(self.page-1, self.per_page)

    @property
    def prev_num(self):
        """Number of the previous page."""
        return self.page - 1

    @property
    def has_prev(self):
        """True if a previous page exists"""
        return self.page > 1

    def next(self):
        return self.query.paginate(self.page+1, self.per_page)

    @property
    def has_next(self):
        """True if a next page exists."""
        return self.page < self.pages

    @property
    def next_num(self):
        """Number of the next page"""
        return self.page + 1

    def iter_pages(self, left_edge=2, left_current=2,
                   right_current=5, right_edge=2):
        """Iterates over the page numbers in the pagination.  The four
        parameters control the thresholds how many numbers should be produced
        from the sides.  Skipped page numbers are represented as `None`.
        This is how you could render such a pagination in the templates:

        .. sourcecode:: html+jinja

            {% macro render_pagination(pagination, endpoint) %}
              <div class=pagination>
              {%- for page in pagination.iter_pages() %}
                {% if page %}
                  {% if page != pagination.page %}
                    <a href="{{ url_for(endpoint, page=page) }}">{{ page }}</a>
                  {% else %}
                    <strong>{{ page }}</strong>
                  {% endif %}
                {% else %}Paginate
                  <span class=ellipsis>…</span>
                {% endif %}
              {%- endfor %}
              </div>
            {% endmacro %}
        """
        last = 0
        for num in range(1, self.pages + 1):
            if num <= left_edge or \
               (num > self.page - left_current - 1 and
                num < self.page + right_current) or \
               num > self.pages - right_edge:
                if last + 1 != num:
                    yield None
                yield num
                last = num
