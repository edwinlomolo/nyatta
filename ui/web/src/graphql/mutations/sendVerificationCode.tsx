import { gql } from '@apollo/client'

export const sendVerificationCode = gql`
  mutation SendVerificationCode($input: VerificationInput!) {
    sendVerificationCode(input: $input) {
      success
    }
  }
`
