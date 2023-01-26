import { useContext } from 'react'

import { useQuery } from '@apollo/client'

import { Box } from '@chakra-ui/react'

import { AuthContext } from '../../auth'
import { HELLO } from '../../apollo'

import { GlobalLoader } from '../../components'

function ListingsPage() {
  const { isAuthenticated } = useContext(AuthContext)
  const { loading } = useQuery(HELLO, { skip: !!isAuthenticated })

  if (loading) return <GlobalLoader />

  return (
    <Box>Property listings</Box>
  )
}

export default ListingsPage
