import axios from 'axios'
import React, { useState, useEffect } from 'react'
import { Flex, Button, Card, CardBody, Image, Text, Modal, ModalOverlay, ModalContent, ModalHeader, ModalBody, ModalCloseButton, useDisclosure } from '@chakra-ui/react'

import { PeriodSelector } from './PeriodSelector'
import { GetCnameData, GetWeek, ApiURL } from './CnameData'

const SendChoice = () => {
    const week = GetWeek(36)
    const [selectedPeriod, setSelectedPeriod] = useState(null)
    const [scheduleImage, setScheduleImage] = useState(null)
    const [cnameData, setCnameData] = useState(null)
    const { isOpen, onOpen, onClose } = useDisclosure() // 用于控制模态框的显示和隐藏

    const handleSend = async () => {
        try {
            const response = await axios.get(`${ApiURL}/api/v1/get_cname_table?week=${selectedPeriod}`, {
                responseType: 'arraybuffer',
            })

            const blob = new Blob([response.data], { type: 'image/png' })
            const image = URL.createObjectURL(blob)
            setScheduleImage(image)
        } catch (error) {
            console.error('Error fetching schedule image:', error)
        }
    }

    useEffect(() => {
        const fetchData = async () => {
            const data = await GetCnameData()
            setCnameData(data)
        }
        fetchData()
        setSelectedPeriod(week) // 你可以修改为其他默认周期
    }, [])

    useEffect(() => {
        // 如果需要在选择周期时立即发送请求，可以在这里调用 handleSend()
        if (selectedPeriod !== null) {
            handleSend(selectedPeriod)
        }
    }, [selectedPeriod])

    const handleStore = () => {
        // 这里你可以提供一个下载链接，让用户点击下载
        const link = document.createElement('a')
        link.href = scheduleImage
        link.download = `第${selectedPeriod}周课表.png`
        link.click()
        onClose() // 存储后关闭模态框
    }

    return (
        <Flex direction="column" align="center" mt="8">
            {/* Period Selector */}
            <PeriodSelector onSelect={setSelectedPeriod} selectedPeriod={selectedPeriod} />

            {scheduleImage && (
                <Card maxW={['100%', '400px']} mt="4" onClick={onOpen} cursor="pointer">
                    <Image src={scheduleImage} alt="课表" borderRadius="lg" />
                    <CardBody>
                        <Text align={"center"}>
                            <br />
                            {cnameData}
                        </Text>
                    </CardBody>
                </Card>
            )}

            {/* Modal */}
            <Modal isOpen={isOpen} onClose={onClose}>
                <ModalOverlay />
                <ModalContent>
                    <ModalHeader>放大预览</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody>
                        <Image src={scheduleImage} alt="课表" borderRadius="lg" width="100%" height="auto" />
                    </ModalBody>
                    <Button colorScheme="blue" mt="4" onClick={handleStore}>
                        保存
                    </Button>
                </ModalContent>
            </Modal>
        </Flex>
    )
}

export default SendChoice
