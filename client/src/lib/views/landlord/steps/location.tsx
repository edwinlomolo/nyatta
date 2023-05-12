import { Button, FormControl, FormLabel, FormErrorMessage, Input } from '@chakra-ui/react'
import Select from 'react-select'
import { Controller } from 'react-hook-form'


import { HStack, Spacer } from '@chakra-ui/react'

import { usePropertyOnboarding } from '../hooks/property-onboarding'

function Location() {
  const { control, towns, setValue, getValues, handleSubmit, register, formState: { errors }, setStep } = usePropertyOnboarding()
  const onSubmit = (data: any) => { console.log(data) }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl mb={5} isInvalid={Boolean(errors.town)}>
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
      <FormControl mb={5}>
        <FormLabel>Postal Code</FormLabel>
        <Input
          isDisabled
          {...register("postalCode")}
        />
      </FormControl>
      <HStack>
        <Button onClick={() => setStep('description')} colorScheme="green">Back</Button>
        <Spacer />
        <Button colorScheme="green" type="submit">Create</Button>
      </HStack>
    </form>
  )
}

export default Location
