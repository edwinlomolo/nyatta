import { Button, FormControl, FormLabel, Input, Select as ChakraSelect, FormErrorMessage, FormHelperText, VStack } from '@chakra-ui/react'
import { ArrowForwardIcon } from '@chakra-ui/icons'
import { useForm, type SubmitHandler } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

import { descriptionSchema } from '../validations'
import { type DescriptionForm } from '../types'

const propertyOptions = ['Apartment', 'Bungalow', 'Condominium']

function Description () {
  const { setStep, descriptionForm, setDescriptionForm } = usePropertyOnboarding()
  const { register, formState: { errors }, handleSubmit } = useForm<DescriptionForm>({
    defaultValues: descriptionForm,
    resolver: yupResolver(descriptionSchema),
    mode: 'onChange'
  })

  const onSubmit: SubmitHandler<DescriptionForm> = values => {
    setDescriptionForm(values)
    setStep('location')
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack spacing={{ base: 4, md: 6 }}>
        <FormControl isInvalid={Boolean(errors.name)}>
          <FormLabel>Name</FormLabel>
          <Input {...register('name')} />
          {(errors.name != null) && <FormErrorMessage>{`${errors.name.message}`}</FormErrorMessage>}
          <FormHelperText>This is the name of your property</FormHelperText>
        </FormControl>
        <FormControl isInvalid={Boolean(errors?.propertyType)}>
          <FormLabel>Property Type</FormLabel>
          <ChakraSelect {...register('propertyType')} placeholder="Select property type">
            {propertyOptions.map((item, index) => <option key={index} value={item}>{item}</option>)}
         </ChakraSelect>
         {(errors.propertyType != null) && <FormErrorMessage>{`${errors.propertyType.message}`}</FormErrorMessage>}
         <FormHelperText>This is your property type</FormHelperText>
        </FormControl>
        <Button colorScheme="green" rightIcon={<ArrowForwardIcon />} type="submit">Next</Button>
      </VStack>
   </form>
  )
}

export default Description
