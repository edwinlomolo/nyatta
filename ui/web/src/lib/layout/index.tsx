import type { ReactNode } from 'react'

import { Box } from '@chakra-ui/react'

interface LayoutProps {
  children: ReactNode
}

const Layout = ({ children }: LayoutProps) => (
    <Box
      minH="100vh"
      bg="gray.100"
    >
      {/* TODO Header */}
      <Box ml={{base: 0, md: 60 }}>
        {children}
      </Box>
    </Box>
  )

export default Layout
