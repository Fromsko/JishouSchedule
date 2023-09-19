# -*- encoding: utf-8 -*-
# @File     : api
# @Time     : 2023-09-19 02:17:40
# @Docs     : 📦 对外接口
from fastapi import FastAPI, Request
from fastapi.responses import RedirectResponse

from tools.spider_cname import DATADIR, log
from .method_api import MethodRouter, Cname

Base_api = FastAPI(
    title="吉首大学课表查询小工具",
    description="""
    本站点用于{吉首大学} 学生查询课程表信息
    请注意以下免责声明：

    - 作者和开发者不对使用本应用程序导致的任何错误或损失负责。
    - 使用本应用程序的用户应自行核实课程信息，以免发生任何误解或错误。
    - 作者和开发者保留随时更改或终止本应用程序的权利，不另行通知。
    """,
    docs_url=None,
    version="5.2.0",
    # redoc_url="/redocs",
)
Base_api.include_router(MethodRouter)


@Base_api.get("/ping")
async def ping(request: Request):
    """
    ### 测试程序是否正常运行

    用于测试程序是否正常运行的简单路由。

    **返回**:
    - `dict`: 包含应答信息的字典，格式示例:
      ```json
      {
          "ping": "pong",
          "client_ip": "客户端的IP地址"
      }
      ```
    """
    return {
        "ping": "pong",
        "client_ip": request.client.host,
    }


@Base_api.exception_handler(404)
async def redirect_to_login(request: Request, exc: Exception):
    # 重定向到主页面
    return RedirectResponse(url="/redoc")


@Base_api.on_event("startup")
async def startup():
    if list(DATADIR.iterdir()) == []:
        Cname._update()
    log.info("程序已经启动!")
