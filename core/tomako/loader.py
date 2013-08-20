# coding: utf-8
import os

from mako.lookup import TemplateLookup
from tornado.template import Loader


class MakoTemplateLoader(Loader):
    def __init__(self, root_directory, **kwargs):
        super(MakoTemplateLoader, self).__init__(root_directory, **kwargs)
        path = os.path.abspath(root_directory)
        self._lookup = TemplateLookup(directories=path, input_encoding='utf-8',
                                      output_encoding='utf-8',
                                     )

    def _create_template(self, name):
        template = self._lookup.get_template(name)
        template.generate = template.render

        return template
