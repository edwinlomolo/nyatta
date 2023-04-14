import { HttpLink } from '@apollo/client'
import { isPrototypeEnv } from '@helpers'

const httpLink = new HttpLink({
  uri: `${isPrototypeEnv() ? process.env.NEXT_PUBLIC_LOCAL_API : process.env.NEXT_PUBLIC_BASE_API}/api`,
})

export default httpLink
