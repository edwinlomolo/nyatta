import { setContext } from '@apollo/client/link/context'
import { type CookieValueTypes } from 'cookies-next'

const authLink = (jwt?: CookieValueTypes) => setContext((_, previousContext) => {
  const { headers } = previousContext
  return {
    ...previousContext,
    headers: {
      ...headers,
      'keep-alive': 'true',
      ...(jwt && { Authorization: `Bearer ${jwt}` })
    }
  }
})

export default authLink
