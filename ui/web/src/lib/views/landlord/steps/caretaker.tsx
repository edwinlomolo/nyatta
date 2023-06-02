import { Button, HStack, FormControl, FormErrorMessage, FormLabel, Input, Spacer, VStack } from '@chakra-ui/react'
import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import { yupResolver } from '@hookform/resolvers/yup'

import { useForm, SubmitHandler } from 'react-hook-form'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

import { caretakerSchema } from '../validations'
import { CaretakerForm } from '../types'

function Caretaker() {
  const { setStep, caretakerForm, setCaretakerForm } = usePropertyOnboarding()
  const { register, handleSubmit, formState: { errors } } = useForm<CaretakerForm>({
    defaultValues: caretakerForm,
    resolver: yupResolver(caretakerSchema),
  })

  const onSubmit: SubmitHandler<CaretakerForm> = data => {
    setCaretakerForm(data)
    setStep("units")
  }
  const goBack = () => setStep("pricing")

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack spacing={{ base: 4, md: 6 }}>
        <FormControl isInvalid={Boolean(errors?.firstName)}>
          <FormLabel>First Name</FormLabel>
          <Input
            {...register("firstName")}
          />
          {errors?.firstName && <FormErrorMessage>{errors?.firstName.message}</FormErrorMessage>}
        </FormControl>
      </VStack>
      <HStack>
        <Button colorScheme="green" leftIcon={<ArrowBackIcon />} onClick={goBack}>Go back</Button>
        <Spacer />
        <Button type="submit" colorScheme="green" rightIcon={<ArrowForwardIcon />}>Next</Button>
      </HStack>
    </form>
  )
}

export default Caretaker
