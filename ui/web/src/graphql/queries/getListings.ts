import { gql } from '@apollo/client'

export const getListings = gql`
  query GetListings($input: ListingsInput!) {
    getListings(input: $input) {
      id
      town
    }
  }
`
