import React, { useCallback, useEffect } from 'react'

import { useAuth0 } from '@auth0/auth0-react'
import { useCookies } from 'react-cookie'

import AuthContext from './AuthContext'
import { Http } from '../../utils'
import { apiUrl } from '../../helpers'

import { GlobalLoader } from '../../components'

interface Props {
  children: React.ReactNode
}

function AuthProvider({ children }: Props) {
  const { isLoading, loginWithRedirect, logout, user, isAuthenticated } = useAuth0()
  const [cookies, setCookie, removeCookie] = useCookies(['jwt'])
  const login = useCallback(() => loginWithRedirect(), [loginWithRedirect])
  const handleLogout = useCallback(
    () => {
      removeCookie('jwt')
      logout({ returnTo: window.location.origin })
    },
    [logout, removeCookie]
  )

  useEffect(() => {
    async function initializeClient() {
      const http = new Http()
      // Auth user
      if (isAuthenticated) {
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
  }, [isAuthenticated, cookies, setCookie, user])

  if (isLoading) return <GlobalLoader />

  return (
    <AuthContext.Provider
      value={{
        user,
        login,
        isAuthenticated,
        logout: handleLogout,
        cookies,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider
