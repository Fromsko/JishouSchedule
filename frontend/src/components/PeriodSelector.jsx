// PeriodSelector.js
import { Menu, MenuButton, MenuList, MenuItem, Button } from '@chakra-ui/react'
import { ChevronDownIcon } from '@chakra-ui/icons'
import React, { useState } from 'react'

export const PeriodSelector = ({ onSelect, selectedPeriod }) => {
    const [isOpen, setIsOpen] = useState(false)

    const handleSelect = (period) => {
        onSelect(period)
        setIsOpen(false)
    }

    return (
        <Menu isOpen={isOpen} onClose={() => setIsOpen(false)}>
            <MenuButton as={Button} rightIcon={<ChevronDownIcon />} onClick={() => setIsOpen(!isOpen)}>
                {selectedPeriod ? `第${selectedPeriod}周` : '选择周期'}
            </MenuButton>
            <MenuList>
                {[...Array(20)].map((_, index) => (
                    <MenuItem key={index + 1} onClick={() => handleSelect(index + 1)}>
                        第{index + 1}周
                    </MenuItem>
                ))}
            </MenuList>
        </Menu>
    )
}
