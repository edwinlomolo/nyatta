import React from 'react'

import { User } from '@auth0/auth0-react'

interface AuthContext {
  user: User | undefined
  isAuthenticated: boolean
  isAuthenticating: boolean
  login: () => void
  logout: () => void
  cookies: Record<string, any> | undefined,
}

export default React.createContext<AuthContext>({
  user: undefined,
  isAuthenticated: false,
  isAuthenticating: false,
  login: () => {},
  logout: () => {},
  cookies: undefined,
})
