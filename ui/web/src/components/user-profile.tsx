import { Box, Button, HStack, Icon, Flex, Avatar, Menu, MenuButton, MenuList, MenuItem, Spinner } from '@chakra-ui/react'
import { useSession, signOut, signIn } from 'next-auth/react'
import { FaAngleDown } from 'react-icons/fa'

const UserMenu = ({ ...rest }): JSX.Element => {
  const { data: session, status } = useSession()

  return (
    <HStack spacing={4} {...rest}>
      {status === 'loading' && <Spinner />}
      {status !== 'loading' && status !== 'authenticated' && (
        <Button onClick={() => signIn('google')}>Sign In</Button>
      )}
      {status !== 'loading' && status === 'authenticated' && (
        <Flex>
          <Menu>
            <MenuButton>
              <HStack>
                <Avatar
                  src={`${session?.user?.image}`}
                  loading="eager"
                />
                <Box display={{ base: "none", md: "flex" }}>
                  <Icon as={FaAngleDown} />
                </Box>
              </HStack>
            </MenuButton>
            <MenuList>
              <MenuItem>{session?.user?.email}</MenuItem>
              <MenuItem onClick={() => signOut()}>Sign Out</MenuItem>
            </MenuList>
          </Menu>
        </Flex>
      )}
    </HStack>
  )
}

export default UserMenu
