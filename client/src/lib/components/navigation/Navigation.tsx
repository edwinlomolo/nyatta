import Link from 'next/link'

import { useUser } from '@auth0/nextjs-auth0/client'

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

import { UserProfile } from '@auth0/nextjs-auth0/client'

import { Dropdown } from '../dropdown'
import { GlobalLoader } from '../loader'

interface Props {
  user: UserProfile | undefined
  login?: () => void
  logout?: () => void
}

function Navigation() {
  const { user, isLoading } = useUser()

  if (isLoading) {
    return <GlobalLoader />
  }

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
            {text: 'Landlord'}
          ]}
        />
      </Flex>
      <Spacer />
      <Flex justifyContent="end">
      {user && (
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
      )}
      {!user && (
        <Button as={"a"} href="/api/auth/login" colorScheme="green">Sign In</Button>
      )}
      </Flex>
    </Flex>
  )
}

export default Navigation
