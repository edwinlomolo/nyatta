import { Text } from '@chakra-ui/react'
import { useAuth0 } from '@auth0/auth0-react'

function UserHome() {
  const { user } = useAuth0()

  return (
    <Text>{JSON.stringify(user)}</Text>
  )
}

export default UserHome
