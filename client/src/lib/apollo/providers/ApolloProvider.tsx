import { PropsWithChildren, useContext, useState, useCallback } from 'react'

import { ApolloClient, ApolloProvider as ApolloClientProvider, NormalizedCacheObject } from '@apollo/client'
import useDeepCompareEffect from 'use-deep-compare-effect'

import { AuthContext } from '../../auth'
import createClient from '../createClient'

let apolloClient: ApolloClient<NormalizedCacheObject> | null

export const getApolloClient = (): ApolloClient<NormalizedCacheObject> | null => apolloClient
export let resetApp = (): void => {}

function ApolloProvider({ children }: PropsWithChildren) {
  const { cookies, isAuthenticating } = useContext(AuthContext)
  const [client, setClient] = useState<ApolloClient<NormalizedCacheObject> | null>(null)
  const shouldCreateClient = useCallback(
    (jwt?: string) => {
      return createClient(jwt)
    },
    []
  )

  useDeepCompareEffect(() => {
    const nextClient = shouldCreateClient(cookies?.jwt)
    apolloClient = nextClient
    setClient(nextClient)
  }, [cookies, shouldCreateClient])

  if (!client || isAuthenticating) {
    return null
  }

  return <ApolloClientProvider client={client}>{children}</ApolloClientProvider>
}

export default ApolloProvider
