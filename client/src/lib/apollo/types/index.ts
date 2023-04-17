import { ApolloClient, NormalizedCacheObject } from '@apollo/client'

import { CookieValueTypes } from 'cookies-next'

export interface UseApolloClientParams {
  jwt?: CookieValueTypes
}

export interface UseApolloClient {
  client: ApolloClient<NormalizedCacheObject> | null
}
