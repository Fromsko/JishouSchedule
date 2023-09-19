# -*- encoding: utf-8 -*-
# @File     : method_api
# @Time     : 2023-09-19 19:35:20
# @Docs     : 📦 拓展接口
from fastapi import Request, APIRouter, Response
from fastapi.responses import FileResponse

from tools.spider_cname import CanmeTable, log

Cname = CanmeTable()
MethodRouter = APIRouter(prefix="/api/v1")


@MethodRouter.get('/get_cname_table')
async def get_cname_table(request: Request, response: Response, week: int):
    """
    ### 获取课表图片

    获取指定周数的课表图片。

    **参数**:
    - `request` (Request): FastAPI Request对象
    - `response` (Response): FastAPI Response对象
    - `week` (int): 周数，表示要获取哪一周的课表

    **返回**:
    - `FileResponse` or `dict`: 
      - 如果成功，返回课表图片的FileResponse。
      - 如果失败，返回包含错误信息的字典。

    **错误格式**:
    ```json
    {"code": 404, "err": "错误消息"}
    ```
    """
    msg = {}
    try:
        img_path = Cname.get_cname_data(f"第{week}周", "img")
        return FileResponse(img_path, media_type="image/jpeg")
    except Exception as err:
        msg['code'] = 404
        msg['err'] = err.__str__()
        log.error(err)
    return msg


@MethodRouter.get('/get_cname_data')
async def get_cname_data(request: Request, response: Response, week: int):
    """
    ### 获取课表数据

    获取指定周数的课表数据，返回格式为JSON。

    **参数**:
    - `request` (Request): FastAPI Request对象
    - `response` (Response): FastAPI Response对象
    - `week` (int): 周数，表示要获取哪一周的课表数据

    **返回**:
    - `dict`: 包含课表数据的字典，格式示例: 
      ```json
      {"code": 200, "data": {...}, "week": "第X周"}
      ```
      如果发生错误，字典中会包含错误信息。

    **错误格式**:
    ```json
    {"code": 404, "err": "错误消息"}
    ```
    """
    msg = {"code": 200, "data": None, "week": f"第{week}周"}
    try:
        result = Cname.get_cname_data(msg['week'], 'json')
        msg['data'] = result
    except Exception as err:
        msg['code'] = 404
        msg['err'] = err.__str__()
    return msg
