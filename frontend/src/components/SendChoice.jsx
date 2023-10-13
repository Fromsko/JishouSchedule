import axios from 'axios'
import React, { useState, useEffect } from 'react'
import { Flex, Button, Card, Image, Modal, ModalOverlay, ModalContent, ModalHeader, ModalBody, ModalCloseButton, useDisclosure } from '@chakra-ui/react'

import { PeriodSelector } from './PeriodSelector'
import { GetCnameData, GetWeek, ApiURL, CnameTable, ParseCnameData } from './CnameData'
import TypingText from './TypingText'


// 触发请求
const handleSend = async (WeeklyChoice, ImgStatus) => {
    try {
        const response = await axios.get(`${ApiURL}/api/v1/get_cname_table?week=${WeeklyChoice}`, {
            responseType: 'arraybuffer',
        })

        const blob = new Blob([response.data], { type: 'image/png' })
        ImgStatus(URL.createObjectURL(blob))
    } catch (error) {
        console.error('Error fetching schedule image:', error)
    }
}

const SendChoice = () => {
    const week = GetWeek(36)
    const [Selected, setSelected] = useState(null) // 周期选择
    const [ImgStatus, setImgStatus] = useState(null) // 图片存储
    const [cnameData, setCnameData] = useState(null) // 数据存储
    const [loading, setLoading] = useState(true) // 添加 loading 状态
    const { isOpen, onOpen, onClose } = useDisclosure() // 用于控制模态框的显示和隐藏

    useEffect(() => {
        const fetchData = async () => {
            try {
                const data = await GetCnameData()
                setCnameData(data)
            } finally {
                setLoading(false) // 设置 loading 为 false，表示数据加载完成
            }
        }

        fetchData()
        setSelected(week)
    }, [])

    useEffect(() => {
        // 根据选择改变
        if (Selected !== null) {
            handleSend(Selected, setImgStatus)
        }
    }, [Selected])

    const handleStore = () => {
        // 这里你可以提供一个下载链接，让用户点击下载
        const link = document.createElement('a')
        link.href = ImgStatus
        link.download = `第${Selected}周课表.png`
        link.click()
        onClose() // 存储后关闭模态框
    }

    return (
        <Flex direction="column" align="center">
            {/* Period Selector */}
            <PeriodSelector onSelect={setSelected} selectedPeriod={Selected} />

            {ImgStatus && (
                <Card maxW={['100%', '500px']} mt="3" onClick={onOpen} cursor="pointer">
                    <Image src={ImgStatus} alt="课表" borderRadius="lg" />
                </Card>
            )}


            <Flex direction="column" align="start" mt="5">
                {loading ? (
                    <TypingText
                        text="Loading"
                        speed={100}
                        onFinish={() => { }}
                        timeOut={500}
                    />
                ) : (
                    cnameData !== null && (
                        <CnameTable data={ParseCnameData(cnameData)} />
                    )
                )}
            </Flex>

            {/* Modal */}
            < Modal isOpen={isOpen} onClose={onClose} >
                <ModalOverlay />
                <ModalContent>
                    <ModalHeader>放大预览</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody>
                        <Image src={ImgStatus} alt="课表" borderRadius="lg" width="100%" height="auto" />
                    </ModalBody>
                    <Button colorScheme="blue" mt="4" onClick={handleStore}>
                        保存
                    </Button>
                </ModalContent>
            </Modal >
        </Flex>
    )
}

export default SendChoice
