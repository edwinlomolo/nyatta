'use client'

import { Box, VStack, Text } from '@chakra-ui/react'

import { useSignIn } from '@hooks'
import SignInForm from 'form/sign-in'
import VerifySignInForm from 'form/verify-signin'

const Login = (): JSX.Element => {
  const { status } = useSignIn()

  return (
    <VStack mt={10}>
      <Box>
        <Text mb={5} align="left" fontWeight="bold" fontSize="2xl">Sign In with Phone</Text>
        {status === 'sign-in' && <SignInForm />}
			  {status === 'pending' && <VerifySignInForm />}
      </Box>
    </VStack>
  )
}

export default Login
