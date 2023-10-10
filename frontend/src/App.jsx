import { ChakraProvider } from '@chakra-ui/react'
import { Navigation } from './components/Navigation'
import SendChoice from './components/SendChoice'

function App () {

  return (
    <ChakraProvider>
      <>
        {/* 状态栏组件 */}
        <Navigation />
        {/* 选择和发送组件 */}
        <SendChoice />
      </>
    </ChakraProvider>
  )
}

export default App
