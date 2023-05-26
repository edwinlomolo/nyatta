import { Button, Container, FormControl, FormLabel, Input, Select as ChakraSelect, FormErrorMessage, FormHelperText, VStack } from '@chakra-ui/react'
import { ArrowForwardIcon } from '@chakra-ui/icons'

import { usePropertyOnboarding } from '../hooks/property-onboarding'

const propertyOptions = ['Apartment', 'Bungalow', 'Condominium']

function Description() {
  const { handleSubmit, register, formState: { errors }, setStep } = usePropertyOnboarding()
  const onSubmit = () => setStep('location')

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Container>
        <VStack spacing={{ base: 4, md: 10 }}>
          <FormControl isInvalid={Boolean(errors.name)}>
            <FormLabel>Name</FormLabel>
            <Input
              {...register('name', {
                required: 'Property name is required',
                pattern: {
                  value: /^[A-Za-z ]+$/i,
                  message: 'Should be a string value',
                }
              })}
            />
            <FormHelperText>This is the name of your property</FormHelperText>
            {errors.name && <FormErrorMessage>{`${errors.name.message}`}</FormErrorMessage>}
          </FormControl>
          <FormControl isInvalid={Boolean(errors?.propertyType)}>
            <FormLabel>Property Type</FormLabel>
            <ChakraSelect {...register('propertyType', { required: 'Property type is required' })} placeholder="Select property type">
              {propertyOptions.map((item, index) => <option key={index} value={item}>{item}</option>)}
            </ChakraSelect>
            <FormHelperText>This is your property type</FormHelperText>
            {errors.propertyType && <FormErrorMessage>{`${errors.propertyType.message}`}</FormErrorMessage>}
          </FormControl>
          <Button colorScheme="green" rightIcon={<ArrowForwardIcon />} type="submit">Next</Button>
        </VStack>
      </Container>
   </form>
  )
}

export default Description
