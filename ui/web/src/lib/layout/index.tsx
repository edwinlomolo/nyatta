import type { ReactNode } from 'react'

import { Box, Flex } from '@chakra-ui/react'

import Navigation from '../components/navigation'

interface LayoutProps {
  children: ReactNode
}

const Layout = ({ children }: LayoutProps) => (
    <Box>
      <Navigation />
      <Flex w="100%" direction="column">
        {children}
      </Flex>
    </Box>
  )

export default Layout
