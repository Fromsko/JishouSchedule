# -*- encoding: utf-8 -*-
# @File     : app
# @Time     : 2023-09-15 18:23:34
# @Docs     : 吉首大学教务系统爬虫 🥳
import asyncio

from aiofiles import open
from playwright.async_api import async_playwright
from apscheduler.schedulers.asyncio import AsyncIOScheduler

from tools.paser_cname import Parser
from tools.gen_picture import DrawPicture
from tools import TESTDIR, ViewInfo, CACHEDIR, log, ErrorStatus, DATADIR, IMGDIR


class Spider(DrawPicture, Parser):
    """ 爬虫主类 """

    def __init__(self, file: str = "config.json") -> None:
        super().__init__(file)
        self.content_save = str(TESTDIR.joinpath("Index.html"))
        self.main_page_save = str(TESTDIR.joinpath("MainPage.png"))
        self.table_page_save = str(TESTDIR.joinpath("TablePage.png"))
        self.weather_info_save = str(TESTDIR.joinpath("weather.png"))
        self.pepole_info_save = str(TESTDIR.joinpath("pepoleInfo.png"))

    async def fetch(self, page, context):
        Option: dict = self.load_config()
        # 登录
        await page.goto(ViewInfo.LoginPageURL)
        await page.get_by_placeholder("学工号").click()
        await page.get_by_placeholder("学工号").fill(Option['username'])
        await page.get_by_placeholder("密码").click()
        await page.get_by_placeholder("密码").fill(Option['password'])
        await page.get_by_role("button", name="立即登录").click()

        # 等待加载
        await page.wait_for_selector("html")
        await page.wait_for_url(page.url)

        try:
            await page.screenshot(path=self.main_page_save)
            await page.locator(ViewInfo.IndexPeopleInfo).screenshot(path=self.pepole_info_save)
            await page.locator(ViewInfo.IndexWeatherPage).screenshot(path=self.weather_info_save)
        except Exception as err:
            log.exception(err)
        else:
            await self.download_content(self.content_save, page)

        # 获取课表数据
        async with page.expect_popup() as page_info:
            await page.get_by_role("link", name="教务系统（师生入口）").click()
            await page.wait_for_url(page.url)
        page_two = await page_info.value

        # 等待加载完成
        await page_two.wait_for_load_state("load")
        await page_two.wait_for_timeout(300)
        await page_two.screenshot(path=self.table_page_save)

        try:
            href_attribute = await page_two.locator(ViewInfo.TableHref).get_attribute("href")

            if href_attribute != "":
                table_page = await context.new_page()
                await table_page.goto("https://webvpn.jsu.edu.cn" + href_attribute)

                for weekly in range(1, 20):
                    try:
                        await self.download_task(table_page, weekly)
                    except Exception as err:
                        log.exception(f"|{weekly}|下载失败=> {err}")
            else:
                raise RuntimeError(ErrorStatus.ServerError)
        except (RuntimeError, Exception) as err:
            log.exception(err)

        # 退出
        await page_two.close()
        await page.get_by_role("button", name="设置").click()
        await page.locator("li").filter(has_text="退出").click()

    async def download_content(self, filename: str, page):
        async with open(filename, "w", encoding="utf-8") as f:
            result = await page.content()
            await f.write(result)

    async def download_task(self, table_page, weekly):
        """ 下载任务 """
        # 展开页面
        await table_page.locator("#zc").select_option(f"{weekly}")
        # 下载
        async with table_page.expect_download() as download_info:
            await table_page.get_by_role("button", name="打 印").click()
            # 下载信息
            download = await download_info.value
            await download.save_as(f'{CACHEDIR.joinpath(f"第{weekly}周课表.xls")}')

        if download_info.is_done():
            log.info(f"[任务队列 {weekly}/{20}] -第{weekly}周课表- 下载完成")

    async def __task(self) -> None:
        """ 启动爬虫任务 """
        async with async_playwright() as playwright:
            browser = await playwright.chromium.launch(headless=True)
            context = await browser.new_context()
            page = await context.new_page()

            await self.fetch(page, context)

            await context.close()
            await browser.close()
        log.info("最新数据获取成功!")

    def _update(self):
        """ 转换更新 """
        for i in CACHEDIR.iterdir():
            data = self.gen_data(i)
            self.save_json_file(
                DATADIR.joinpath(f"{data['周次']}.json"),
                content=data,
            )
            self.draw_photo(data)
        log.info("数据生成完毕!")

    def AutoTask(self):
        """ 定时任务执行 """
        cheduler = AsyncIOScheduler()

        cheduler.add_job(
            self.__task, 'interval',
            hours=12, max_instances=1,
        )
        cheduler.add_job(self._update, 'interval', hours=2)
        cheduler.start()

        asyncio.get_event_loop().run_forever()


class CanmeTable(Spider):
    def get_cname_data(self, week_: str, type_: str = "json"):
        if type_ == 'json':
            dirs = DATADIR.iterdir()
        elif type_ == 'img':
            dirs = IMGDIR.iterdir()

        for dir_ in dirs:
            if week_ in dir_.name:
                if type_ == "json":
                    return self.load_json_file(dir_)
                elif type_ == 'img':
                    return dir_
        else:
            raise ValueError(f"没有查询到 {week_}")
