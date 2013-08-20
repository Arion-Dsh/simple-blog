#!/usr/bin/python
# -*- coding: utf-8 -*-
# author :Arion

def truncate(text, length=255, killwords=False, end='...'):
    """Return a truncated copy of the string. The length is specified
    with the first parameter which defaults to ``255``. If the second
    parameter is ``true`` the filter will cut the text at length. Otherwise
    it will discard the last word. If the text was in fact
    truncated it will append an ellipsis sign (``"..."``). If you want a
    different ellipsis sign than ``"..."`` you can specify it using the
    third parameter.

    .. sourcecode::mako

        ${ "foo bar"|truncate(5) }
            -> "foo ..."
        ${ "foo bar"|truncate(5, True) }
            -> "foo b..."
    """
    if len(text) <= length:
        return text
    elif killwords:
        return text[:length] + end
    words = text.split(' ')
    result = []
    m = 0
    for word in words:
        m += len(word) + 1
        if m > length:
            break
        result.append(word)
    result.append(end)
    return u' '.join(result)