import { Box, Container, SimpleGrid } from '@chakra-ui/react'
import Link from 'next/link'

const PrivacyPolicy = (): JSX.Element => (
  <Container mt={10} textAlign="center">
    <SimpleGrid columns={{sm: 1, md: 2}} spacingY={10}>
      <Box>
        <Box
          mb={10}
          border="1px green solid"
          p={4}
          borderRadius="md"
          as={Link}
          cursor="pointer"
          target="_blank"
          _hover={{
            shadow: "md"
          }}
          href="https://hackmd.io/@qPtRW4tNR2KxOuqBVUBp5A/S1FipYnP3"
        >
          Privacy Policy
        </Box>
      </Box>
      <Box>
        <Box
          border="1px green solid"
          p={4}
          borderRadius="md"
          as={Link}
          cursor="pointer"
          target="_blank"
          _hover={{
            shadow: "md"
          }}
          href="https://hackmd.io/@qPtRW4tNR2KxOuqBVUBp5A/S13cD5hwn"
        >
          Terms of Service
        </Box>
      </Box>
    </SimpleGrid>
  </Container>
)

export default PrivacyPolicy
