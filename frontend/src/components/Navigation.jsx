import React, { useState } from 'react'
import { Flex, Text, Image, Box } from '@chakra-ui/react'

const AppLogo = "https://github.com/Fromsko/JishouSchedule/raw/main/res/logo.png"

export const Navigation = () => {
    const [isHovered, setIsHovered] = useState(false)

    const logoStyle = {
        height: '6em',
        padding: '1.5em',
        willChange: 'filter',
        transition: 'filter 300ms',
        borderRadius: '50%',
        filter: isHovered ? 'drop-shadow(0 0 2em #61dafbaa)' : 'none',
    }

    const containerStyle = {
        backdropFilter: 'blur(20px)', // 调整毛玻璃效果更明显
        margin: 'auto', // 居中
        borderRadius: '80px', // 圆滑边角
        width: '70%', // 调整容器宽度
    }

    return (
        <Box style={containerStyle} borderRadius="20px">
            <Flex align="center" justify="space-between" p="4" bg="gray.100" boxShadow={isHovered ? 'lg' : 'none'} transition="box-shadow 300ms">
                <Flex align="center">
                    <a
                        href="https://github.com/Fromsko/JishouSchedule"
                        target="_blank"
                        onMouseEnter={() => setIsHovered(true)}
                        onMouseLeave={() => setIsHovered(false)}
                    >
                        <Image src={AppLogo} style={logoStyle} alt="App logo" />
                    </a>
                    <Text fontSize="xl" lineHeight="1" verticalAlign="middle">
                        吉首大学课表查询
                    </Text>
                </Flex>
            </Flex>
        </Box>
    )
}
