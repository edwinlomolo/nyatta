import { ArrowForwardIcon } from '@chakra-ui/icons'
import { Button, FormControl, Input, FormErrorMessage, FormHelperText, VStack } from '@chakra-ui/react'
import { yupResolver } from '@hookform/resolvers/yup'
import { useForm, type SubmitHandler } from 'react-hook-form'

import { type DescriptionForm } from '../types'
import { DescriptionSchema } from '../../../form/validations'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

const Description = () => {
  const { setStep, descriptionForm, setDescriptionForm } = usePropertyOnboarding()
  const { register, formState: { errors }, handleSubmit } = useForm<DescriptionForm>({
    defaultValues: descriptionForm,
    resolver: yupResolver(DescriptionSchema)
  })

  const onSubmit: SubmitHandler<DescriptionForm> = values => {
    setDescriptionForm(values)
    setStep('type')
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack spacing={{ base: 4, md: 6 }}>
        <FormControl isInvalid={Boolean(errors.name)}>
          <Input {...register('name')} />
          {(errors.name != null) && <FormErrorMessage>{`${errors.name.message}`}</FormErrorMessage>}
          <FormHelperText>This is how your property will be referred on the platform</FormHelperText>
        </FormControl>
        <Button colorScheme="green" rightIcon={<ArrowForwardIcon />} type="submit">{`Let's Go`}</Button>
      </VStack>
   </form>
  )
}

export default Description
