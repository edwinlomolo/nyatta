import { ApolloClient, InMemoryCache, from, NormalizedCacheObject } from '@apollo/client'
import { RetryLink } from '@apollo/client/link/retry'
import { CookieValueTypes } from 'cookies-next'

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

export default createClient
