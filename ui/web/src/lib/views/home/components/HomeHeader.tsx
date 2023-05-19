import { Box, Button, Flex, Heading, Divider } from '@chakra-ui/react'

function HomeHeader() {
  return (
    <Flex>
      <Box
        w="100%"
        h="50vh"
        mb="80px"
      >
        <Heading size="3xl" mt="80px" textAlign="center">
          Find rental properties or homes
        </Heading>
        <Flex gap={4} mt={10} w="100%" justifyContent="center">
          <Box>
            <Button as="a" href="/listings" size="lg">
              Find A Home
            </Button>
          </Box>
          <Box>
            <Divider orientation="vertical" />
          </Box>
          <Box>
            <Button as="a" href="/landlord" size="lg">
              Landlord
            </Button>
          </Box>
        </Flex>
      </Box>
    </Flex>
  )
}

export default HomeHeader
