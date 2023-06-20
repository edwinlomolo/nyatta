'use client'

import { useMutation } from '@apollo/client'
import { ArrowForwardIcon } from '@chakra-ui/icons'
import { Button, FormControl, FormLabel, FormErrorMessage, Select, Input, HStack } from '@chakra-ui/react'
import { yupResolver } from '@hookform/resolvers/yup'
import { useForm, SubmitHandler } from 'react-hook-form'

import { SignInSchema } from './validations'

import { sendVerificationCode as SEND_VERIFICATION } from '@gql'
import { useSignIn } from '@hooks'
import { SignInForm } from '@types'

const SignInForm = (): JSX.Element => {
	const { setStatus, signInForm, setSignInForm } = useSignIn()
	const [sendVerificationCode, { loading: sendingVerification }] = useMutation(SEND_VERIFICATION)
	const { handleSubmit, register, formState: { errors } } = useForm<SignInForm>({
		defaultValues: signInForm,
		resolver: yupResolver(SignInSchema),
	})

	const onSubmit: SubmitHandler<SignInForm> = async data => {
		setSignInForm(data)
		if (!sendingVerification) { // Synchronous
			await sendVerificationCode({
				variables: {
					input: {
						countryCode: "KE",
						phone: `${data.countryCode}${data.phone}`,
					},
				},
				onCompleted: data => {
					setStatus(data?.sendVerificationCode.success)
				}
			})
		}
	}

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
			<FormControl isInvalid={Boolean(errors?.phone)}>
				<FormLabel>Phone</FormLabel>
				<HStack>
					<Select {...register("countryCode")}>
						<option value="+254">+254</option>
					</Select>
					<Input
						type="tel"
						{...register("phone")}
					/>
					<Button type='submit' isLoading={sendingVerification} rightIcon={<ArrowForwardIcon />} />
				</HStack>
				{((errors?.phone) != null) && <FormErrorMessage>{`${errors?.phone.message}`}</FormErrorMessage>}
			</FormControl>
		</form>
  )
}

export default SignInForm
