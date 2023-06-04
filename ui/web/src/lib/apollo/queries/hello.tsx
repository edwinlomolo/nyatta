import { gql } from '@apollo/client'

const HELLO = gql`
  query Hello {
    hello
  }
`

export default HELLO
