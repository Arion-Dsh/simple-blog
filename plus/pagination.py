# -*- coding:utf-8 -*-
import math
from tornado import gen


class Paginate(object):
    
    def __init__(self, date, page, per_page ):
        """docstring for __init__"""
        self.page = int(page)
        self.per_page = int(per_page)
        self.total = len(date)
        self.date = date
        
       
    @property   
    def _list(self):
        
        start_index = (self.page - 1) * self.per_page
        end_index = self.page * self.per_page
        
        return self.date[start_index:end_index].all()
        
    @property 
    def pages(self):
        """The total number of pages"""
        return int(math.ceil(self.total / float(self.per_page)))
    @property
    def prev_num(self):
        """Number of the previous page."""
        return self.page - 1
    @property
    def has_prev(self):
        """True if a previous page exists"""
        return self.page > 1
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
