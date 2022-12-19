import { onError } from '@apollo/client/link/error'

const errorLink = onError(({ networkError }) => {
  console.error(networkError)
})

export default errorLink
