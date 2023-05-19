import { Button, FormControl, FormLabel, Input, Select, FormErrorMessage } from '@chakra-ui/react'

import { usePropertyOnboarding } from '../hooks/property-onboarding'

const propertyOptions = ['Apartment', 'Bungalow', 'Condominium', 'Keja']

function Description() {
  const { handleSubmit, register, formState: { errors }, setStep } = usePropertyOnboarding()
  const onSubmit = () => setStep('location')

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl mb={5} isInvalid={Boolean(errors.name)}>
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
        {errors.name && <FormErrorMessage>{`${errors.name.message}`}</FormErrorMessage>}
      </FormControl>
      <FormControl mb={5} isInvalid={Boolean(errors.propertyType)}>
        <FormLabel>Property Type</FormLabel>
        <Select {...register('propertyType', { required: 'Property type is required' })} placeholder="Select property type">
          {propertyOptions.map((item, index) => <option key={index} value={item}>{item}</option>)}
        </Select>
        {errors.propertyType && <FormErrorMessage>{`${errors.propertyType.message}`}</FormErrorMessage>}
      </FormControl>
      <Button colorScheme="green" type="submit">Next</Button>
   </form>
  )
}

export default Description
