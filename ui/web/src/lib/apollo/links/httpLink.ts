import { createUploadLink } from 'apollo-upload-client'

const httpLink = createUploadLink({
  uri: `${process.env.NEXT_PUBLIC_BASE_API}/api`,
})

export default httpLink
