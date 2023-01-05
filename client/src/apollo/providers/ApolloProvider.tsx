import { PropsWithChildren, useContext, useState, useCallback, useEffect } from 'react'

import { ApolloClient, ApolloProvider as ApolloClientProvider, NormalizedCacheObject } from '@apollo/client'

import { AuthContext } from '../../auth'
import createClient from '../createClient'

let apolloClient: ApolloClient<NormalizedCacheObject> | null

export const getApolloClient = (): ApolloClient<NormalizedCacheObject> | null => apolloClient
export let resetApp = (): void => {}

function ApolloProvider({ children }: PropsWithChildren) {
  const { cookies } = useContext(AuthContext)
  const [client, setClient] = useState<ApolloClient<NormalizedCacheObject> | null>(null)
  const shouldCreateClient = useCallback(
    (jwt?: string) => {
      return createClient(jwt)
    },
    []
  )

  useEffect(() => {
    const nextClient = shouldCreateClient(cookies?.jwt)
    setClient(nextClient)
  }, [cookies, shouldCreateClient])

  if (!client) {
    return null
  }

  return <ApolloClientProvider client={client}>{children}</ApolloClientProvider>
}

export default ApolloProvider
