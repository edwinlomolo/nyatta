import { HttpLink } from '@apollo/client'

const httpLink = new HttpLink({
  uri: `${process.env.NEXT_PUBLIC_BASE_API}/api`,
})

export default httpLink
