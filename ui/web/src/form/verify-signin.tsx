'use client'

import { useMutation } from '@apollo/client'
import { Button, HStack, FormControl, FormLabel, FormErrorMessage, FormHelperText, Input } from '@chakra-ui/react'
import { yupResolver } from '@hookform/resolvers/yup'
import { useSignIn } from '@hooks'
import { VerifySignInForm } from '@types'
import { useForm, SubmitHandler } from 'react-hook-form'

import { verifyVerificationCode as VERIFY_CODE } from '@gql'
import { VerifySignInSchema } from 'form/validations'


const VerifySignInForm = (): JSX.Element => {
  const [verifyCode, { loading: verifyingCode }] = useMutation(VERIFY_CODE)
  const { setStatus, signInForm } = useSignIn()
  const { handleSubmit, register, formState: { errors } } = useForm<VerifySignInForm>({
    resolver: yupResolver(VerifySignInSchema),
  })

  const onSubmit: SubmitHandler<VerifySignInForm> = async data => {
    if (!verifyingCode) { // Synchronous
      await verifyCode({
        variables: {
          input: {
            phone: `${signInForm?.countryCode}${signInForm?.phone}`,
            countryCode: "KE",
            verifyCode: data.code,
          },
        },
        onCompleted: data => {}
      })
    }
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={Boolean(errors?.code)}>
        <FormLabel>Verify Code</FormLabel>
        <HStack>
          <Input
            {...register("code")}
            type="number"
          />
          <Button isLoading={verifyingCode} type="submit">Sign In</Button>
        </HStack>
        {((errors?.code) != null) && <FormErrorMessage>{`${errors?.code.message}`}</FormErrorMessage>}
        <FormHelperText>Enter 6-digit code sent to your phone</FormHelperText>
      </FormControl>
    </form>
  )
}

export default VerifySignInForm
