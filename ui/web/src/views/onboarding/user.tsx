'use client'

import { Box, Container, Text } from '@chakra-ui/react'

import UserOnboarding from 'form/user-onboarding'

const Page = (): JSX.Element => (
    <Container>
      <Text fontSize={{base: "2xl", md: "3xl"}}>Finish onboarding your profile</Text>
      <Box>
        <UserOnboarding />
      </Box>
    </Container>
  )

export default Page
