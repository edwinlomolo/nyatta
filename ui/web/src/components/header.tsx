import { Flex, IconButton, Text } from '@chakra-ui/react'
import { FaBars } from 'react-icons/fa'

import UserMenu from './user-profile'

const Header = (): JSX.Element => (
  <Flex
    position="sticky"
    px={4}
    top="0"
    h={20}
    alignItems="center"
    zIndex="1"
    borderBottomWidth="1px"
    borderBottomColor="gray.200"
    justifyContent={{base: "space-between", md: "flex-end" }}
  >
    <IconButton
      display={{base: "flex", md: "none"}}
      aria-label="open menu"
      variant="outline"
      icon={<FaBars />}
    />
    <Text
      display={{base: "flex", md: "none"}}
      fontSize="2xl"
      fontWeight="bold"
    >
      Nyatta
    </Text>
    <UserMenu />
  </Flex>
)

export default Header
