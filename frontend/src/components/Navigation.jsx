import React, { useState } from 'react'
import { Flex, Text, Image, Box } from '@chakra-ui/react'

export const Navigation = () => {
    const [isHovered, setIsHovered] = useState(false)

    const logoStyle = {
        height: '5em', // 调整 Logo 大小
        willChange: 'filter',
        transition: 'filter 300ms',
        borderRadius: '60%',
        filter: isHovered ? 'drop-shadow(0 0 2em #61dafbaa)' : 'none',
    }

    const containerStyle = {
        backgroundColor: isHovered ? '#f0f0f0' : 'transparent', // 调整背景颜色
        margin: 'auto', // 居中
        borderRadius: '20px',
        width: '100%', // 调整容器宽度
        transition: 'background-color 300ms', // 添加过渡效果
        textAlign: 'center', // 文字居中
    }

    const textStyle = {
        color: isHovered ? '#333' : '#000', // 调整文字颜色
        transition: 'color 300ms', // 添加过渡效果
    }

    return (
        <Box style={containerStyle} borderRadius="50px" overflow="hidden">
            <Flex
                align="center"
                justify="space-between"
                p="3"
                height="60px"  // 调整高度
                boxShadow={isHovered ? 'lg' : 'none'}
                transition="box-shadow 300ms"
            >
                <Flex
                    align="center"
                    onMouseEnter={() => setIsHovered(true)}
                    onMouseLeave={() => setIsHovered(false)}
                >
                    <a href="https://github.com/Fromsko/JishouSchedule" target="_blank">
                        <Image src="/src/assets/logo.png" style={logoStyle} ml="3" />
                    </a>
                    <Text fontSize="xl" noOfLines={1} lineHeight="2" verticalAlign="middle" ml="4" style={textStyle}>
                        吉首大学课表查询
                    </Text>
                </Flex>
            </Flex>
        </Box>
    )
}
