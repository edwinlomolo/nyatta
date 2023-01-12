import { PropsWithChildren, useContext } from 'react'

import { Box, ChakraProvider, Flex } from '@chakra-ui/react'

import { AuthContext } from '../auth'
import { Navigation } from '../components/navigation'
import { theme } from '../theme'

function Main({ children }: PropsWithChildren) {
  const { user, login, logout, isAuthenticated } = useContext(AuthContext)

  return (
    <ChakraProvider theme={theme}>
      <Box>
        <Navigation user={user} logout={logout} login={login} isAuthenticated={isAuthenticated} />
        <Flex w="100%" direction="column">
          {children}
        </Flex>
      </Box>
    </ChakraProvider>
  )
}

export default Main
