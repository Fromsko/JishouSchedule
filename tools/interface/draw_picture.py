# -*- encoding: utf-8 -*-
# @File     : draw_picture
# @Time     : 2023-09-19 01:05:09
# @Docs     : 接口定义
from pathlib import Path
from abc import ABC, abstractmethod

from PIL import Image


class DrawBase(ABC):
    """ 绘制的抽象接口 """

    @abstractmethod
    def base_photo(self) -> Image:
        """基础模板

        Returns:
            img: 图片对象
        """
        pass

    @abstractmethod
    def save_draw(self, img_class: Image, save_load: Path):
        """存储图片

        Args:
            img_class (Image): 图片对象
            save_load (Path): 存储路径
        """
        pass

    @abstractmethod
    def draw_photo(self, cname_data: dict) -> Path:
        """ 绘制图片 """
        pass

    @abstractmethod
    def check_data(self, data: dict) -> dict:
        """ 校验绘制数据 """
        pass
