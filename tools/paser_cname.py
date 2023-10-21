# -*- encoding: utf-8 -*-
# @File     : paser_cname
# @Time     : 2023-09-19 04:23:18
# @Docs     : 解析数据文件
from pathlib import Path

import pandas as pd

from tools import ErrorStatus, log


class Parser:
    # @MethodLogger(skip_save=False, skip_log=True)
    def parser_excel(self, parser_file: Path) -> dict:
        try:
            # 读取 Excel 文件
            raw_data = pd.read_excel(parser_file, header=None)

            # 将dataframe类型转化为list类型
            cname_base = raw_data.head().values.tolist()[0][0]

            # 获取第一行第一列的班级信息
            cname_info = raw_data.iat[1, 0].split("        ")

            # 获取星期信息（第三行的第二列到第八列）
            weekdays = raw_data.iloc[2, 1:8].tolist()

            # 获取节次信息（第一列的第四行到第八行）
            time_slots = raw_data.iloc[3:8, 0].tolist()

            # 提取课表数据（第三行的第二列到第八列）
            course_data = raw_data.iloc[3:8, 1:8].values.tolist()

            # 备注
            remark = raw_data.iat[8, 1]

            # 构造基础参数
            param = {
                "学校": cname_base.split(" ")[0],
                "姓名": cname_base.split(" ")[1],
                "周次": parser_file.name.split(".")[0][:-2] or None,
                "课程信息": {
                    "课程数据": course_data,
                    "星期数据": weekdays,
                    "节次数据": time_slots,
                },
            }
            param.update(
                self._dict_gen(cname_info),
            )
            param.update(
                self._dict_gen(remark),
            )
        except Exception as err:
            log.exception(err)
            return {}
        return param

    # @MethodLogger(skip_log=False, skip_save=False)
    def parser_table(self, mate_data: dict) -> dict:
        """解析课表

        Args:
            mate_data (dict): 元课程数据

        Returns:
            _type_: dict
        """
        course_schedule = {}
        course_data = mate_data["课程信息"]["课程数据"]
        week_days = mate_data["课程信息"]["星期数据"]
        time_slots = mate_data["课程信息"]["节次数据"]

        try:
            # 使用嵌套循环遍历时间段和星期信息
            for y, week_day in enumerate(week_days):
                course_schedule[week_day] = {}
                for x, time_slot in enumerate(time_slots):
                    # 找到对应的数据
                    course = course_data[x][y].strip().split("\n")
                    if (course_size := len(course)) == 1:
                        # 当前节次没有课程
                        course = "没课哟"
                    elif (course_size == 3 and "（网络）" in course[0]):
                        # 选修课
                        course[0] = course[0][4:]
                        course = {
                            "课程名": course[0] or "",
                            "老师": "",
                            "周次": course[1] or "",
                            "教室": course[2] or "",
                        } 
                    else:
                        # 正常课程
                        course = {
                            "课程名": course[0] or "",
                            "老师": course[1] or "",
                            "周次": course[2] or "",
                            "教室": course[3] or "",
                        }
                    course_schedule[week_day][time_slot] = course
        except Exception as err:
            course_schedule["err"] = err
            log.exception(ErrorStatus.ParserError)
        return course_schedule

    def _dict_gen(self, data: list):
        orgin = {}
        try:
            if isinstance(data, str):
                data = [data]

            for info in data:
                avg = info.split("：")
                orgin[avg[0]] = avg[1]

        except Exception as err:
            orgin["err"] = err
            log.exception(ErrorStatus.ParserError)
        return orgin

    def gen_data(self, parser_file: Path):
        param = {}
        try:
            param = self.parser_excel(parser_file)
            param["课程信息"]["课程数据"] = self.parser_table(param)
        except Exception:
            log.exception(ErrorStatus.ParserError)
        return param
