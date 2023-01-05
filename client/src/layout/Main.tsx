import { PropsWithChildren, useContext } from 'react'

import { Box } from '@chakra-ui/react'

import { AuthContext } from '../auth'
import { Navigation } from '../components/navigation'

function Main({ children }: PropsWithChildren) {
  const { user, login, logout, isAuthenticated } = useContext(AuthContext)

  return (
    <Box p={2}>
      <Navigation user={user} logout={logout} login={login} isAuthenticated={isAuthenticated} />
      {children}
    </Box>
  )
}

export default Main
