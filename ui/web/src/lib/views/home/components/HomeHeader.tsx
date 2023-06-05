import { Box, Button, Flex, Heading, Divider, VStack } from '@chakra-ui/react'

const HomeHeader = () => (
    <Flex>
      <VStack
        w="100%"
        h="50vh"
        spacing={{ base: 6, md: 8 }}
      >
        <Heading size="3xl" mt="80px" textAlign="center">
          Find rental properties or homes
        </Heading>
        <Flex gap={4} justifyContent="center">
            <Button as="a" href="/listings" size="lg">
              Find A Home
            </Button>
          <Box>
            <Divider orientation="vertical" />
          </Box>
            <Button as="a" href="/landlord" size="lg">
              Home Owner
            </Button>
        </Flex>
      </VStack>
    </Flex>
  )

export default HomeHeader
