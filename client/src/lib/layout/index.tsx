import type { ReactNode } from 'react'

import { Box, Flex } from '@chakra-ui/react'

import { Navigation } from '../components/navigation'

interface LayoutProps {
  children: ReactNode
}

function Layout({ children }: LayoutProps) {
  return (
    <Box>
      <Navigation />
      <Flex w="100%" direction="column">
        {children}
      </Flex>
    </Box>
  )
}

export default Layout
