import { gql } from '@apollo/client'

export const HANDSHAKE_USER = gql`
  mutation Handshake($input: HandshakeInput!) {
    handshake(input: $input) {
      id
      onboarding
    }
  }
`
