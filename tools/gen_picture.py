# -*- encoding: utf-8 -*-
# @File     : draw_picture
# @Time     : 2023-09-19 01:05:09
# @Docs     : 接口定义
import random
from pathlib import Path

from PIL import Image, ImageDraw

from .interface.draw_picture import DrawBase
from tools import Config, log, RESDIR, IMGDIR


class DrawPicture(DrawBase, Config):
    """ 绘制图片的具体类 """

    def base_photo(self) -> Image.Image:
        """基础模板

        Returns:
            img: 图片对象
        """

        # 画布 && 画笔
        base_img = Image.new(mode="RGBA", size=(2000, 1000), color="white")
        draw = ImageDraw.Draw(base_img, "RGBA")
        font = self.load_font("Deng", 30)

        # 画格子
        draw.line(xy=(250, 0, 250, 1000), fill="black")  # 最左的竖线
        for i in range(500, 1751, 250):
            draw.line(xy=(i, 0, i, 890), fill="black")  # 竖线
        for i in [100, 235, 375, 515, 655, 750, 890]:
            draw.line(xy=(0, i, 2000, i), fill="black")  # 横线

        # 备注
        draw.text(xy=(95, 930), text="备注", font=font, fill=(47, 79, 79, 255))
        draw.text(xy=(95, 690), text="晚上", font=font, fill="black")  # 晚上

        # 节次
        draw.text(xy=(70, 140), text="第 1-2 节", font=font, fill="black")
        draw.text(xy=(70, 275), text="第 3-4 节", font=font, fill="black")
        draw.text(xy=(70, 415), text="第 5-6 节", font=font, fill="black")
        draw.text(xy=(70, 555), text="第 7-8 节", font=font, fill="black")
        draw.text(xy=(55, 790), text="第 9-10-11 节", font=font, fill="black")

        # 时间
        draw.text(xy=(50, 180), text="08:00—09:40", font=font, fill="black")
        draw.text(xy=(50, 315), text="10:10—11:50", font=font, fill="black")
        draw.text(xy=(50, 455), text="15:00—16:40", font=font, fill="black")
        draw.text(xy=(50, 595), text="16:50—18:30", font=font, fill="black")
        draw.text(xy=(60, 830), text="19:30—21:10", font=font, fill="black")

        # 星期
        days = ["星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期天"]
        date_num = [330, 580, 830, 1080, 1330, 1580, 1830]
        # 星期
        for local, text in zip(date_num, days):
            draw.text(xy=(local, 40), text=text, font=font, fill="black")

        return base_img

    def save_draw(self, img_class: Image.Image, save_load: Path):
        """存储图片

        Args:
            img_class (Image): 图片对象
            save_load (Path): 存储路径
        """
        p_load = save_load.parent
        if not Path.exists(p_load):
            Path.mkdir(p_load)
        img_class.save(save_load)

    def draw_photo(self, cname_data: dict) -> Path:
        # 星期位置
        w_place_data = [250, 500, 750, 1000, 1250, 1500, 1750]
        # 时间位置
        h_place_data = {"第一大节": 100, "第二大节": 235,
                        "第三大节": 375, "第四大节": 515, "第五大节": 750}
        # 颜色列表
        color_list = [
            (251, 255, 242, 200),
            (192, 192, 192, 200),
            (255, 255, 0, 200),
            (244, 164, 95, 200),
            (127, 255, 0, 200),
            (218, 112, 214, 200),
            (156, 147, 133, 200),
            (186, 164, 48, 200),
            (15, 56, 154, 200),
            (49, 65, 196, 200),
            (153, 51, 250, 200),
            (34, 139, 34, 200),
            (255, 192, 203, 200),
            (255, 127, 80, 200),
            (237, 145, 33, 200),
        ]

        # 颜色数据
        color_data = {}
        # 课表数据处理
        data_class: dict = self.check_data(cname_data)
        cname = data_class['data']["班级"]
        week_ = data_class['data']["周次"]

        # 读取模板
        img = Image.open(data_class['init_photo'], "r")
        img.convert("RGBA")
        draw = ImageDraw.Draw(img, "RGBA")
        # 字体
        font = self.load_font("萝莉体", 23)
        deng_ttf = self.load_font("Deng", 30)

        # 遍历绘制
        for week_day, day_name in enumerate(data_class["data"]["课程信息"]["星期数据"]):
            try:
                day_info = data_class["data"]["课程信息"]["课程数据"].get(day_name, {})
                for time_slot, course_info in day_info.items():
                    if isinstance(course_info, str) and course_info == "没课哟":
                        continue  # 没课就跳过
                    if isinstance(course_info, dict):
                        lesson_name = course_info.get("课程名", "")
                        teacher = course_info.get("老师", "")
                        week_info = course_info.get("周次", "")
                        place = course_info.get("教室", "")
                    else:
                        continue  # 非字典类型的数据跳过

                    if lesson_name:
                        # 从课表中获取 课程名 lesson_name
                        if lesson_name in color_data.keys():
                            color = color_data[lesson_name]
                        else:
                            color = random.choice(
                                list(set(color_list) - set(color_data.values()))
                            )
                            color_data[lesson_name] = color

                        draw.rectangle(
                            (
                                w_place_data[week_day],
                                h_place_data[time_slot],
                                w_place_data[week_day] + 249,
                                h_place_data[time_slot] + 140,
                            ),
                            fill=color,
                            outline=color,
                        )

                        draw.text(
                            (w_place_data[week_day] + 5,
                             h_place_data[time_slot] + 15),
                            text=lesson_name,
                            font=font,
                            fill=(0, 0, 0, 255),
                        )
                        draw.text(
                            (w_place_data[week_day] + 5,
                             h_place_data[time_slot] + 40),
                            text=teacher,
                            font=font,
                            fill=(0, 0, 0, 255),
                            align="center",
                        )
                        draw.text(
                            (w_place_data[week_day] + 10,
                             h_place_data[time_slot] + 65),
                            text=week_info,
                            font=font,
                            fill=(0, 0, 0, 255),
                            align="center",
                        )
                        draw.text(
                            (w_place_data[week_day] + 10,
                             h_place_data[time_slot] + 90),
                            text=place,
                            font=font,
                            fill=(0, 0, 0, 255),
                            align="center",
                        )
            except Exception as err:
                log.exception(err)

        draw.text(
            (50, 25),  # 表头数据
            cname.endswith("班") or f"{cname}班",  # 文本信息
            font=font,  # 采用字体
            fill=(0, 0, 0, 255),  # 头部数据绘制
        )
        draw.text(
            (80, 60),
            week_,
            font=font,
            fill=(0, 0, 0, 255),
        )
        draw.text(
            (300, 930), text=data_class["data"]["备注"], font=deng_ttf, fill=(40, 79, 79, 200)
        )

        save_path: Path = IMGDIR.joinpath(f"{cname}{week_}课表.png")
        self.save_draw(img, save_path)
        return save_path

    def check_data(self, data: dict) -> dict:
        """ 校验绘制数据 """
        data_class = {"bool": True, "data": {}}

        if not (init_photo := RESDIR.joinpath("init_photo.png")).exists():
            img = self.base_photo()
            self.save_draw(img, init_photo)

        if data is None:
            raise RuntimeError("传入数据为空")
        else:
            data_class['init_photo'] = init_photo.__str__()
            data_class["data"] = data
        return data_class
