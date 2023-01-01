import { Avatar, Button, Flex, Menu, MenuButton, MenuList, MenuItem, Portal, Spacer } from '@chakra-ui/react'

interface Props {
  isAuthenticated: boolean
  user: Record<string, any> | undefined
  login: () => void
  logout: () => void
}

function Navigation({ user, isAuthenticated, login, logout }: Props) {
  return (
    <Flex align="center">
      {isAuthenticated && <Flex justifyContent="start" textDecoration="underline">Become a Landlord</Flex>}
      <Spacer />
      <Flex justifyContent="end">
      {isAuthenticated ? (
        <Menu>
          <MenuButton type="button">
            <Avatar src={user?.picture} name={user?.name} />
          </MenuButton>
          <Portal>
            <MenuList>
              <MenuItem onClick={logout}>Sign Out</MenuItem>
            </MenuList>
          </Portal>
        </Menu>
      ) : (
        <Button onClick={login}>Sign In</Button>
      )}
      </Flex>
    </Flex>
  )
}

export default Navigation
