# -*- encoding: utf-8 -*-
# @File     : api
# @Time     : 2023-09-19 02:17:40
# @Docs     : 程序入口🚀
import multiprocessing
from uvicorn import run as StartApp

from api.base_api import Base_api, Cname


if __name__ == '__main__':
    multiprocessing.Process(target=Cname.AutoTask).start()
    StartApp(
        app="main:Base_api",
        host="0.0.0.0",
        port=2000,
        reload=False
    )
