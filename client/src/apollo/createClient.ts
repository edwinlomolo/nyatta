import { ApolloClient, InMemoryCache, from, NormalizedCacheObject } from '@apollo/client'
import { RetryLink } from '@apollo/client/link/retry'

import { authLink, errorLink, httpLink } from './links'

const createClient = (jwt?: string): ApolloClient<NormalizedCacheObject> => {
  const cache = new InMemoryCache({})
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
      authLink(jwt),
      errorLink,
      retryLink,
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

export default createClient
