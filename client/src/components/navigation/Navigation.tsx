import {
  Avatar,
  Button,
  Flex,
  Heading,
  Spacer,
} from '@chakra-ui/react'
import { ChevronDownIcon } from '@chakra-ui/icons'
import { Link as ReactLink } from 'react-router-dom'

import { Dropdown } from '../dropdown'

interface Props {
  isAuthenticated: boolean
  user: Record<string, any> | undefined
  login: () => void
  logout: () => void
}

function Navigation({ user, isAuthenticated, login, logout }: Props) {
  return (
    <Flex p={2} align="center">
      <Flex gap={4} justifyContent="start">
        <Heading as={ReactLink} to="/" size="md">Nyatta</Heading>
        <Dropdown
          children={
            <>
              Partners
              <ChevronDownIcon />
            </>
          }
          options={[
            {text: 'Become a partner'}
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
              text: 'Sign Out',
              onClick: logout,
            }
          ]}
        />
      ) : (
        <Button colorScheme="green" onClick={login}>Sign In</Button>
      )}
      </Flex>
    </Flex>
  )
}

export default Navigation
