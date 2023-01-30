import { HttpLink } from '@apollo/client'

const httpLink = new HttpLink({
  uri: "http://localhost:4000/api",
})

export default httpLink
