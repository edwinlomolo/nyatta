import { Button, Flex, Text } from '@chakra-ui/react'

interface Props {
  isAuthenticated: boolean
  login: () => void
  logout: () => void
}

function Navigation({ isAuthenticated, login, logout }: Props) {
  return (
    <Flex gap={4}>
      {isAuthenticated ? (
        <Button onClick={logout}>Sign Out</Button>
      ) : (
        <Button onClick={login}>Sign In</Button>
      )}
    </Flex>
  )
}

export default Navigation
