import { Box, Button, Flex, Heading, Divider } from '@chakra-ui/react'

function HomePage() {
  return (
    <Flex
      p={4}
    >
      <Box w="100%" bg="tomato">
        <Heading textAlign="center">
          Home page
        </Heading>
      </Box>
    </Flex>
  )
}

export default HomePage
