import {
  Avatar,
  Button,
  Flex,
  Heading,
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
  Portal,
  Spacer,
} from '@chakra-ui/react'

interface Props {
  isAuthenticated: boolean
  user: Record<string, any> | undefined
  login: () => void
  logout: () => void
}

function Navigation({ user, isAuthenticated, login, logout }: Props) {
  return (
    <Flex align="center">
      <Flex justifyContent="start">
        <Heading>Nyatta</Heading>
      </Flex>
      <Spacer />
      <Flex justifyContent="end">
      {isAuthenticated ? (
        <Menu>
          <MenuButton type="button">
            <Avatar src={user?.picture} name={user?.name} />
          </MenuButton>
          <Portal>
            <MenuList>
              <MenuItem textDecoration="underline" onClick={logout}>Sign Out</MenuItem>
            </MenuList>
          </Portal>
        </Menu>
      ) : (
        <Button colorScheme="green" onClick={login}>Sign In</Button>
      )}
      </Flex>
    </Flex>
  )
}

export default Navigation
