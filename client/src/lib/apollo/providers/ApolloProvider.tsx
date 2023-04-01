import { PropsWithChildren, useContext, useState, useCallback } from 'react'

import { ApolloClient, ApolloProvider as ApolloClientProvider, NormalizedCacheObject } from '@apollo/client'
import { CookieValueTypes, getCookie } from 'cookies-next'
import useDeepCompareEffect from 'use-deep-compare-effect'

import { AuthContext } from '../../auth'
import createClient from '../createClient'

let apolloClient: ApolloClient<NormalizedCacheObject> | null

export const getApolloClient = (): ApolloClient<NormalizedCacheObject> | null => apolloClient
export let resetApp = (): void => {}

function ApolloProvider({ children }: PropsWithChildren) {
  const jwt = getCookie('jwt')
  const { isAuthenticating } = useContext(AuthContext)
  const [client, setClient] = useState<ApolloClient<NormalizedCacheObject> | null>(null)
  const shouldCreateClient = useCallback(
    (jwt?: CookieValueTypes) => {
      return createClient(jwt)
    },
    []
  )

  useDeepCompareEffect(() => {
    const nextClient = shouldCreateClient(jwt)
    apolloClient = nextClient
    setClient(nextClient)
  }, [shouldCreateClient])

  if (!client || isAuthenticating) {
    return null
  }

  console.log(jwt)
  return <ApolloClientProvider client={client}>{children}</ApolloClientProvider>
}

export default ApolloProvider
