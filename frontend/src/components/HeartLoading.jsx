import React from 'react'
import { Box } from '@chakra-ui/react'


const HeartLoading = () => {
  const heartStyle = {
    position: 'relative',
    width: '100px',
    height: '90px',
    left: '10px',
    top: '10px',
    animation: 'heart infinite 2s linear',
  }

  const heartBeforeAfterStyle = {
    position: 'absolute',
    top: '0',
    left: '30px',
    width: '30px',
    height: '50px',
    content: '""',
    transform: 'rotate(-45deg)',
    transformOrigin: '0 100%',
    borderRadius: '30px 30px 0 0',
    background: 'pink',
  }

  return (
    <Box style={heartStyle} className="loading">
      <Box style={heartBeforeAfterStyle}></Box>
      <Box style={{ ...heartBeforeAfterStyle, left: '0', transform: 'rotate(45deg)', transformOrigin: '100% 100%' }}></Box>
    </Box>
  )
}

export default HeartLoading
