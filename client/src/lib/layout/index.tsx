import type { ReactNode } from 'react'

import { useUser } from '@auth0/nextjs-auth0/client'

import { Box, Flex } from '@chakra-ui/react'

import { Navigation } from '../components/navigation'

interface LayoutProps {
  children: ReactNode
}

function Layout({ children }: LayoutProps) {
  const { user } = useUser()

  return (
    <Box>
      <Navigation user={user} isAuthenticated={!!user} />
      <Flex w="100%" direction="column">
        {children}
      </Flex>
    </Box>
  )
}

export default Layout
