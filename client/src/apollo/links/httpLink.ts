import { createHttpLink } from '@apollo/client'

const httpLink = createHttpLink({
  uri: "http://localhost:4000/graphql",
})

export default httpLink
