import axios from 'axios'

export const ApiURL = ""

// 获取当前星期
function getWeekly () {
    const now = new Date()
    const weekday = now.getDay()
    const weekdayMap = {
        1: "星期一",
        2: "星期二",
        3: "星期三",
        4: "星期四",
        5: "星期五",
        6: "星期六",
        0: "星期日",
    }

    return weekdayMap[weekday]
}

// 获取当前第几周
export function GetWeek (startMon) {
    const now = new Date()
    const startOfYear = new Date(now.getFullYear(), 0, 1)
    const days = Math.floor((now - startOfYear) / (24 * 60 * 60 * 1000))
    const weeks = Math.ceil((days + startOfYear.getDay() + 1) / 7)

    return (weeks - startMon).toString()
}


// 获取课表数据
export async function GetCnameData () {
    const week = GetWeek(36)
    const weekly = getWeekly()
    let result = ""
    let fetchUrl = `${ApiURL}/api/v1/get_cname_data?week=${week}`

    try {
        const response = await axios.get(fetchUrl)
        const data = response.data


        // 判断是否请求成功
        if (data.code !== 200) {
            console.error("课表数据获取失败!")
            return result
        }

        // 遍历本周数据
        for (const [key, value] of Object.entries(data.data.课程信息.课程数据[weekly])) {
            if (value !== "没课哟") {
                result += `${key} ${value.课程名 || ""} ${value.老师.split('(')[0] || ""} ${value.教室 || ""}\n`
            }
        }
        if (result === "") {
            return "一天都没有课呢,好好休息休息吧~"
        }
        return result
    } catch (error) {
        console.error("Error fetching data:", error)
        return result
    }
}

