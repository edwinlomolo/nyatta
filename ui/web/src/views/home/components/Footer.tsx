import { Box, Flex, Text } from '@chakra-ui/react'

const Footer = () => (
    <Box
      textAlign="center"
    >
      <Flex gap={4}>
        <Text>
          {`Copyright@${new Date().getFullYear()}`}
        </Text>
        <Text
          as="a"
          href="mailto:edwinmoses535@gmail.com"
          _hover={{
            cursor: 'pointer'
          }}
          textDecoration="underline"
        >
          Contact Us
        </Text>
      </Flex>
    </Box>
  )

export default Footer
