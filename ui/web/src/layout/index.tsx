'use client'

import type { ReactNode } from 'react'

import { Box, Text } from '@chakra-ui/react'

import Header from '../components/header'

interface LayoutProps {
  children: ReactNode
}

const Layout = ({ children }: LayoutProps) => (
  <Box
    minH="100vh"
    bg="white"
  >
    <Header />
    <Box p={4}>
      {children}
    </Box>
    <Box bottom="0" left="0" w="100%" textAlign="center" position="fixed">
      <Text
        as="a"
        href="mailto:edwinmoses535@gmail.com"
        _hover={{
          cursor: 'pointer'
        }}
        textDecoration="underline"
      >
        Contact Us
      </Text>
    </Box>
  </Box>
)

export default Layout
