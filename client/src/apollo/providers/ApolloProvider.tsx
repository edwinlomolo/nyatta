import { PropsWithChildren, useState, useCallback, useEffect } from 'react'

import { useAuth0 } from '@auth0/auth0-react'
import { ApolloClient, ApolloProvider as ApolloClientProvider, NormalizedCacheObject } from '@apollo/client'
import { useCookies } from 'react-cookie'

import createClient from '../createClient'
import { Http } from '../../utils'
import { apiUrl } from '../../helpers'

import { GlobalLoader } from '../../components/'

let apolloClient: ApolloClient<NormalizedCacheObject> | null

export const getApolloClient = (): ApolloClient<NormalizedCacheObject> | null => apolloClient
export let resetApp = (): void => {}

function ApolloProvider({ children }: PropsWithChildren) {
  const { isAuthenticated, isLoading, user } = useAuth0()
  const [client, setClient] = useState<ApolloClient<NormalizedCacheObject> | null>(null)
  const [cookies, setCookie] = useCookies(['jwt'])
  const shouldCreateClient = useCallback(
    (jwt?: string) => {
      return createClient(jwt)
    },
    []
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

    const nextClient = shouldCreateClient(cookies.jwt)

    setClient(nextClient)
    initializeClient()

  }, [user, setCookie, isAuthenticated, isLoading, shouldCreateClient])

  if (isLoading) return <GlobalLoader highlight="Setting up" />
  if (!client) {
    return null
  }

  return <ApolloClientProvider client={client}>{children}</ApolloClientProvider>
}

export default ApolloProvider
