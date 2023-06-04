import React from 'react'

import { type UserProfile } from '@auth0/nextjs-auth0/client'
interface AuthContext {
  user: UserProfile | undefined
  isAuthenticated: boolean
  isAuthenticating: boolean
  logout: () => void
}

export default React.createContext<AuthContext>({
  user: undefined,
  isAuthenticated: false,
  isAuthenticating: false,
  logout: () => {}
})
