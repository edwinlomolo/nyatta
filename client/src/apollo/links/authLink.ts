import { setContext } from '@apollo/client/link/context'

const authLink = (jwt?: string) => setContext((_, { headers }) => {
  return {
    ...headers,
    Authorization: jwt ? `Bearer ${jwt}` : '',
  }
})

export default authLink
