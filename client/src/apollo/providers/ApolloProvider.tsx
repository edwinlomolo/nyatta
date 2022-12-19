import { PropsWithChildren, useState, useCallback, useEffect } from 'react'

import { useAuth0 } from '@auth0/auth0-react'
import { ApolloClient, ApolloProvider as ApolloClientProvider, NormalizedCacheObject } from '@apollo/client'

import createClient from '../createClient'
import { Http } from '../../utils'
import { apiUrl } from '../../helpers'

function ApolloProvider({ children }: PropsWithChildren) {
  const { isAuthenticated, isLoading, user } = useAuth0()
  const [client, setClient] = useState<ApolloClient<NormalizedCacheObject> | null>(null)
  const shouldCreateClient = useCallback(
    (jwt?: string) => {
      return createClient(jwt)
    },
    []
  )

  useEffect(() => {
    const nextClient = shouldCreateClient()

    async function initializeClient() {
      const http = new Http()
      // Generate user auth token
      if (isAuthenticated) {
        const newUser = {
          first_name: user?.given_name,
          last_name: user?.family_name,
          email: user?.email
        }
        try {
          const res = await http.post(apiUrl, newUser)
          console.log(res)
        } catch (error) {
          console.error(error)
        }
      }
    }

    if (isLoading) return
    if (!isAuthenticated) return setClient(nextClient)

    setClient(nextClient)

    initializeClient()

  }, [user, isAuthenticated, isLoading, shouldCreateClient])

  if (!client) {
    return null
  }

  return <ApolloClientProvider client={client}>{children}</ApolloClientProvider>
}

export default ApolloProvider
