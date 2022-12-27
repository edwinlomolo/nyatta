import { PropsWithChildren } from 'react'

import { useAuth0 } from '@auth0/auth0-react'
import { Box } from '@chakra-ui/react'
import { useCookies } from 'react-cookie'

import { Navigation } from '../components/navigation'

function Main({ children }: PropsWithChildren) {
  const { isAuthenticated, loginWithRedirect, logout } = useAuth0()
  // TODO: rfr to somewhere modular
  const [cookies, setCookie, removeCookie] = useCookies(['jwt'])

  const handleLogout = () => {
    removeCookie('jwt')
    logout({ returnTo: window.location.origin })
  }

  return (
    <Box>
      <Navigation logout={handleLogout} login={loginWithRedirect} isAuthenticated={isAuthenticated} />
      {children}
    </Box>
  )
}

export default Main
