import { type ApolloClient, type NormalizedCacheObject } from '@apollo/client'

import { type CookieValueTypes } from 'cookies-next'

export interface UseApolloClientParams {
  jwt?: CookieValueTypes
}

export interface UseApolloClient {
  client: ApolloClient<NormalizedCacheObject> | null
}
