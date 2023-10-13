import React, { useState } from 'react'
import {
  ChakraProvider,
  Flex,
  Divider,
  Text,
  Box,
  SimpleGrid,
} from '@chakra-ui/react'
import { Navigation } from './components/Navigation'
import SendChoice from './components/SendChoice'
import HeartLoading from './components/HeartLoading'
import TypingText from './components/TypingText'

function App () {
  const [typingDone, setTypingDone] = useState(false)

  const Load = () => {
    setTypingDone(true)
  }

  return (
    <ChakraProvider>
      <Flex direction="column" align="center" justify="center" h="100vh">
        {/* Loading 动画和打字效果 */}
        <Flex direction="column" align="center">
          {!typingDone && (
            <Flex align="center" justify="center" h="20vh">
              <HeartLoading />
            </Flex>
          )}

          {!typingDone && (
            <Flex align="center">
              <TypingText
                text="追风赶月莫停留，平芜尽处是春山。"
                speed={100}
                onFinish={Load}
                timeOut={1500}
                isclear={true}
              />
            </Flex>
          )}
        </Flex>

        {typingDone && (
          <Box
            bg="rgba(255, 255, 255, 0.1)"
            borderRadius="8px"
            p="4"
            w={['100%', '80%', '60%']}
            mx="auto" // 水平居中
            boxShadow="0px 2px 4px rgba(0, 0, 0, 0.1)" // 减小阴影
          >
            <SimpleGrid columns={1} spacing={4} alignItems="center">
              {/* 状态栏组件 */}
              <Box bg="rgba(255, 255, 255, 0.2)" p="2" borderRadius="8px">
                <Navigation />
              </Box>
              {/* 选择和发送组件 */}
              <SendChoice />

              {/* 分割线 */}
              <Divider />

              {/* 备案信息 */}
              <Text align="center">xx备xxx号</Text>
            </SimpleGrid>
          </Box>
        )}
      </Flex>
    </ChakraProvider>
  )
}

export default App
