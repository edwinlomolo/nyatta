'use client'

import type { ReactNode } from 'react'

import { Box } from '@chakra-ui/react'

import Brand from '../components/brand'
import Header from '../components/header'

interface LayoutProps {
  children: ReactNode
}

const Layout = ({ children }: LayoutProps) => (
  <Box
    minH="100vh"
    bg="white"
  >
    <Brand display={{ base: "none", md: "block" }} />
    <Header />
    <Box p={4}>
      {children}
    </Box>
  </Box>
)

export default Layout
