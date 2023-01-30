import {
  Avatar,
  Button,
  Box,
  Flex,
  Heading,
  Spacer,
  Text,
} from '@chakra-ui/react'
import { ChevronDownIcon } from '@chakra-ui/icons'
import { Link as ReactLink } from 'react-router-dom'

import { UserProfile } from '@auth0/nextjs-auth0/client'

import { Dropdown } from '../dropdown'

interface Props {
  isAuthenticated: boolean
  user: UserProfile | undefined
  login?: () => void
  logout?: () => void
}

function Navigation({ user, isAuthenticated, login, logout }: Props) {
  return (
    <Flex p={2} align="center">
      <Flex gap={4} justifyContent="start">
        {/*<Heading as={ReactLink} to="/" size="md">Nyatta</Heading>*/}
        <Dropdown
          children={
            <>
              Partners
              <ChevronDownIcon />
            </>
          }
          options={[
            {text: 'Landlord'}
          ]}
        />
      </Flex>
      <Spacer />
      <Flex justifyContent="end">
      {isAuthenticated ? (
        <Dropdown
          children={
            <>
              <Avatar src={user?.picture} name={user?.name} />
            </>
          }
          options={[
            {
              text: (
                <Box>
                  <Text as="b">{user?.name}</Text>
                  <Text>{user?.email}</Text>
                </Box>
              ),
            },
            {
              text: (
                <a href="/api/auth/logout">Log out</a>
              ),
            }
          ]}
        />
      ) : (
        <Button as={"a"} href="/api/auth/login" colorScheme="green" onClick={login}>Sign In</Button>
      )}
      </Flex>
    </Flex>
  )
}

export default Navigation
