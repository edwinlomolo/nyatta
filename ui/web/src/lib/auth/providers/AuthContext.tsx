import React from 'react'

interface User {
  name: string
  picture: string
  email: string
}

interface AuthContext {
  user: User | undefined
  isAuthenticated: boolean
  isAuthenticating: boolean
  logout: () => void
}

export default React.createContext<AuthContext>({
  user: undefined,
  isAuthenticated: false,
  isAuthenticating: false,
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  logout: () => {},
})
