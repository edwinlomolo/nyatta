import { gql } from '@apollo/client'

export const uploadImage = gql`
  mutation UploadImage($file: Upload!) {
    uploadImage(file: $file)
  }
`
