import { gql } from '@apollo/client'

export const verifyVerificationCode = gql`
  mutation VerifyVerificationCode($input: VerificationInput!) {
    verifyVerificationCode(input: $input) {
      success
    }
  }
`
