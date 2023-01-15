import { useQuery } from '@apollo/client'

import { Box } from '@chakra-ui/react'

import { HELLO } from '../../apollo'

import { GlobalLoader } from '../../components'

function ListingsPage() {
  const { data, loading } = useQuery(HELLO)

  if (loading) return <GlobalLoader />

  return (
    <Box>Property listings</Box>
  )
}

export default ListingsPage
