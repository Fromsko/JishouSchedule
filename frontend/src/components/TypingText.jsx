import React, { useEffect, useState } from 'react'
import { Text } from '@chakra-ui/react'

const TypingText = ({ text, speed, onFinish, timeOut, isclear }) => {
    const [displayText, setDisplayText] = useState('')

    useEffect(() => {
        let index = 0
        const intervalId = setInterval(() => {
            setDisplayText((prev) => {
                if (text[index] === undefined) {
                    return text
                }
                return prev + text[index]
            })
            index++
            if (index === text.length) {
                clearInterval(intervalId)
                // 打字完成后触发 onFinish 回调
                setTimeout(() => {
                    if (isclear) {
                        setDisplayText('')
                    }
                    onFinish()
                }, timeOut) // 等待 timeOut 毫秒后清空 displayText
            }
        }, speed)

        // 清理定时器
        return () => clearInterval(intervalId)
    }, [text, speed, onFinish, timeOut])

    return (
        <Text
            sx={{
                fontFamily: 'monospace', // 使用等宽字体以保证效果
                background: 'linear-gradient(to right, rgba(238, 174, 202, 1), rgba(148, 187, 233, 1), rgba(179, 229, 252, 1))',
                backgroundSize: '200% 100%',
                animation: `rainbow ${text.length * speed}ms linear infinite`,
                whiteSpace: 'pre',
                borderRadius: '4px', // 可以添加圆角效果
                padding: '8px', // 可以调整内边距
                fontWeight: 'bold', // 可以加粗字体
            }}
        >
            {displayText}
        </Text>
    )
}

export default TypingText
