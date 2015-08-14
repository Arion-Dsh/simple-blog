#!/usr/bin/env python
# -*- coding:utf-8 -*-

import copy
import re
import mistune


class MyReader(mistune.Renderer):

    def article_img(self, url, dec, sdec=None):
        if sdec == 'None':
            sdec = ''
        out = """<div class='cont-img pos-r'>
                <div class='cont-img-dec pos-a'>%s<p>%s</></div>
                <img src='%s' alt='%s' class="pure-img bdr-s"/></div>""" \
                % (dec, sdec, url, dec)
        return out


class MyBlockGrammar(mistune.BlockGrammar):
    article_img = re.compile(
        r'\[\['                             # [[
        r'([\s\S]+?\|[\s\S]+?\|[\s\S]+?)'   # Page 2|Page 2|page 2
        r'\]\](?!\])'                       # ]]
    )


class MyBlockLexer(mistune.BlockLexer):
    default_rules = copy.copy(mistune.BlockLexer.default_rules)
    default_rules.insert(3, 'article_img')

    def __init__(self, rules=None, **kwargs):
        if rules is None:
            rules = MyBlockGrammar()

        super(MyBlockLexer, self).__init__(rules, **kwargs)

    def parse_article_img(self, m):
        text = m.group(1)
        dec, sdec, url = text.split('|')
        self.tokens.append({
            'type': 'article_img',
            'dec': dec,
            'url': url,
            'sdec': sdec
        })


class MyMarkdown(mistune.Markdown):

    def output_article_img(self):
        dec = self.token['dec']
        url = self.token['url']
        sdec = self.token['sdec']
        return self.renderer.article_img(url, dec, sdec)


renderer = MyReader()
block = MyBlockLexer()
markdown = MyMarkdown(renderer, block=block)
