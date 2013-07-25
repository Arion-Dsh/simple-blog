# coding: utf-8
from mako.runtime import supports_caller


@supports_caller
def ui(context, uiclass, **kwargs):
    parts = uiclass.split('.')
    class_ = parts.pop()
    module = '.'.join(parts)

    uimodule = __import__(module, globals(), locals(), [class_], -1)
    instance = getattr(uimodule, class_)(context['handler'])
    return instance.render(**kwargs)
