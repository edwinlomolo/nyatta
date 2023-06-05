import { useContext } from 'react'

import { ChevronDownIcon } from '@chakra-ui/icons'
import {
  Avatar,
  Button,
  Box,
  Flex,
  Heading,
  Spacer,
  Text
} from '@chakra-ui/react'
import Link from 'next/link'


import AuthContext from '../../auth/providers/AuthContext'
import { Dropdown } from '../dropdown'

const Navigation = () => {
  const { user, logout } = useContext(AuthContext)

  return (
    <Flex p={2} align="center">
      <Flex gap={4} justifyContent="start">
        <Heading as={Link} href="/" size="md">Nyatta</Heading>
        <Dropdown
          children={
            <>
              Partners
              <ChevronDownIcon />
            </>
          }
          options={[
            { text: <Link href="/landlord">Home Owner</Link> }
          ]}
        />
      </Flex>
      <Spacer />
      <Flex justifyContent="end">
      {(user != null) && (
        <Dropdown
          children={
            <>
              <Avatar src={user?.picture!} name={user?.name!} />
            </>
          }
          options={[
            {
              text: (
                <Box>
                  <Text as="b">{user?.name}</Text>
                  <Text>{user?.email}</Text>
                </Box>
              )
            },
            {
              text: (
                <a onClick={logout} href="/api/auth/logout">Log out</a>
              )
            }
          ]}
        />
      )}
      {(user == null) && (
        <Button as={'a'} href="/api/auth/login" colorScheme="green">Sign In</Button>
      )}
      </Flex>
    </Flex>
  )
}

export default Navigation
