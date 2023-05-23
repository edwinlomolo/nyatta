import { FormControl, FormLabel, Input, Select as ChakraSelect, FormErrorMessage, FormHelperText, SimpleGrid, VStack } from '@chakra-ui/react'
import Select from 'react-select'
import { Controller } from 'react-hook-form'

import { usePropertyOnboarding } from '../hooks/property-onboarding'

const propertyOptions = ['Apartment', 'Bungalow', 'Condominium']

function Description() {
  const { control, towns, setValue, getValues, handleSubmit, register, formState: { errors }, setStep } = usePropertyOnboarding()
  const onSubmit = () => setStep('location')

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <SimpleGrid columns={{ base: 1, md: 2 }} spacing={{ base: 4, md: 8 }}>
        <VStack spacing={{ base: 4, md: 6 }}>
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
          <FormControl>
            <FormLabel>Property Type</FormLabel>
            <ChakraSelect {...register('propertyType', { required: 'Property type is required' })} placeholder="Select property type">
              {propertyOptions.map((item, index) => <option key={index} value={item}>{item}</option>)}
            </ChakraSelect>
            <FormHelperText>This is your property type</FormHelperText>
            {errors.propertyType && <FormErrorMessage>{`${errors.propertyType.message}`}</FormErrorMessage>}
          </FormControl>
          <FormControl>
            <FormLabel>Town</FormLabel>
            <Controller
              name="town"
              control={control}
              rules={{ required: { value: true, message: "This is required" } }}
              render={({ field }) => (
                <Select
                  {...field}
                  isClearable
                  isSearchable
                  options={towns}
                  onChange={(newV, _) => { setValue("town", newV); setValue("postalCode", newV?.postalCode) }}
                  value={getValues()?.town}
                  placeholder="Select town"
                />
              )}
            />
            {errors.town && <FormErrorMessage>{`${errors.town.message}`}</FormErrorMessage>}
          </FormControl>
        </VStack>
        <VStack spacing={{ base: 4, md: 6 }}>
          <FormControl>
            <FormLabel>Postal Code</FormLabel>
            <Input
              disabled
              {...register("postalCode")}
            />
          </FormControl>
          <FormControl>
            <FormLabel>Minimum Price</FormLabel>
            <Input
              type="number"
            />
            <FormHelperText>This is the lowest priced unit</FormHelperText>
          </FormControl>
          <FormControl>
            <FormLabel>Maximum Price</FormLabel>
            <Input
              type="number"
            />
            <FormHelperText>This is the highest priced unit</FormHelperText>
          </FormControl>
        </VStack>
      </SimpleGrid>
   </form>
  )
}

export default Description
