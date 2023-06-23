'use client'

import { Box, Flex, Text } from '@chakra-ui/react'

import SearchForm from 'form/search-listings'

const Listings = () => (
  <Flex
    flexDirection="column"
  >
    <Box>
      <SearchForm />
    </Box>
    <Flex alignSelf="center">
      <Text fontSize={{base: "3xl", md: "4xl"}}>No Listings Found</Text>
    </Flex>
  </Flex>
)

export default Listings
