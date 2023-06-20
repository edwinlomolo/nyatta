import { Button, HStack } from '@chakra-ui/react'
import Link from 'next/link'

const UserMenu = (): JSX.Element => (
  <HStack spacing={4}>
    <Button as={Link} href="/login/user">Sign In</Button>
  </HStack>
)

export default UserMenu
