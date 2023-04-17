import { useState, useEffect } from 'react'
import { ApolloClient, InMemoryCache, from, NormalizedCacheObject } from '@apollo/client'
import { RetryLink } from '@apollo/client/link/retry'
import { CookieValueTypes } from 'cookies-next'
import { UseApolloClientParams, UseApolloClient } from './types'

import { authLink, errorLink, httpLink } from './links'



const createClient = (jwt?: CookieValueTypes): ApolloClient<NormalizedCacheObject> => {
  // Caching
  const cache = new InMemoryCache({})
  // Error retry link
  const retryLink = new RetryLink({
    delay: {
      initial: 300,
      jitter: true,
    },
    attempts: {
      max: 2,
      retryIf: error => !!error
    },
  })

  return new ApolloClient({
    link: from([
      // apollo authentication
      authLink(jwt),
      // apollo error handler link
      errorLink,
      retryLink,
      // api link
      httpLink,
    ]),
    cache,
    defaultOptions: {
      watchQuery: {
        fetchPolicy: 'cache-and-network' as const,
      },
    },
  })
}

let apolloClient: ApolloClient<NormalizedCacheObject> | null
export let getApolloClient = (): ApolloClient<NormalizedCacheObject> | null => apolloClient

export function useApolloClient({ jwt }: UseApolloClientParams): UseApolloClient {
  const [client, setClient] = useState<ApolloClient<NormalizedCacheObject> | null>(null)
  useEffect(() => {
    const nextClient = createClient(jwt)
    apolloClient = nextClient
    setClient(nextClient)
  }, [jwt])

  return { client }
}

