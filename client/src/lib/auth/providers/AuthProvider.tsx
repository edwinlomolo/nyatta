import React, { useCallback, useEffect } from 'react'

import { useUser } from '@auth0/nextjs-auth0/client'
import { setCookie, deleteCookie } from 'cookies-next'

import AuthContext from './AuthContext'
import { Http } from '../../utils'
import { apiUrl } from '../../helpers'
import { getApolloClient } from '@apollo'

interface Props {
  children: React.ReactNode
}

function AuthProvider({ children }: Props) {
  const client = getApolloClient()
  const { user, isLoading } = useUser()
  const handleLogout = useCallback(
    () => {
      client?.resetStore()
      deleteCookie('jwt', { path: "/" })
    },
    [client]
  )

  useEffect((): void => {
    async function initializeClient() {
      const http = new Http()
      // Auth user
      if (user && !isLoading) {
        const newUser = {
          first_name: user?.given_name,
          last_name: user?.family_name,
          email: user?.email
        }
        try {
          const res = await http.post(apiUrl, newUser)
          const accessToken: string = res.access_token
          setCookie('jwt', accessToken, { path: '/' })
        } catch (error) {
          console.error(error)
        }
      }
    }
    initializeClient()
  }, [user])

  if (isLoading) return null

  return (
    <AuthContext.Provider
      value={{
        user,
        isAuthenticated: !isLoading && !!user,
        isAuthenticating: isLoading,
        logout: handleLogout,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider
