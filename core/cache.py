#!/usr/bin/env python
# -*- coding:utf-8 -*-
#
#

import settings
import pylibmc


class Cache(object):
    @classmethod
    def func(cls, *args, **kwargs):
        return None

    def __getattr__(self, key):
        return Cache.func

    def __setattr__(self, key, value):
        return Cache.func

class Memcached(object):
    _mc = pylibmc.client.Client(settings.CACHE_HOST, binary = True)

    def __enter__(self):
        if settings.CACHED:
            return Memcached
        else:
            return Cache()

    def __exit__(self, exc_type, exc_val, exc_tb):
        pass

    @classmethod
    def get_cache(cls):
        return cls._mc

    @classmethod
    def get(cls, key, default = None):
        r = cls._mc.get(key)
        if not r:
            r = default
        return r

    @classmethod
    def set(cls, key, value, timeout = 0):
        timeout = timeout if timeout else settings.CACHE_TIMEOUT
        return cls._mc.set(key, value, timeout)

    @classmethod
    def delete(cls, key):
        return cls._mc.delete(key)

    @classmethod
    def flush(cls):
        return cls._mc.flush_all()

    def __getattr__(self, key):
        return Memcached.get(key)

    def __setattr__(self, key, value):
        return Memcached.set(key, value)

    def __delattr__(self, key):
        return Memcached.delete(key)

    def __setitem__(self, key, value):
        return Memcached.set(key, value)

    def __getitem__(self, key):
        return Memcached.get(key)

