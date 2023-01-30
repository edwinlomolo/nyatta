import { Box, Flex, Text } from '@chakra-ui/react'

function Footer() {
  return (
    <Box
      textAlign="center"
      mx="auto"
      mt="100px"
    >
      <Flex gap={4}>
        <Text>
          {`Copyright@${new Date().getFullYear()}`}
        </Text>
        <Text
          as="a"
          href="mailto:info@nyatta.app"
          _hover={{
            cursor: "pointer",
          }}
          textDecoration="underline"
        >
          Contact Us
        </Text>
      </Flex>
    </Box>
  )
}

export default Footer
