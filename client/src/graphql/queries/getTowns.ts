import { gql } from '@apollo/client'

export const getTowns = gql`
  query GetTowns {
    getTowns {
      id
      town
      postalCode
    }
  }
`
