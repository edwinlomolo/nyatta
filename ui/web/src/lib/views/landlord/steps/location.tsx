import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import { Button, FormControl, FormLabel, FormErrorMessage, FormHelperText, Input, VStack, HStack, Spacer } from '@chakra-ui/react'
import { Controller, useForm, type SubmitHandler } from 'react-hook-form'
import Select from 'react-select'

import { usePropertyOnboarding } from '../hooks/property-onboarding'
import { type LocationForm } from '../types'

const Location = () => {
  const { locationForm, setLocationForm, towns, setStep } = usePropertyOnboarding()
  const { control, handleSubmit, register, setValue, getValues, formState: { errors } } = useForm<LocationForm>({
    defaultValues: locationForm,
  })

  const onSubmit: SubmitHandler<LocationForm> = data => {
    setLocationForm(data)
    setStep('amenities')
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack spacing={{ base: 4, md: 6 }}>
        <FormControl isInvalid={Boolean(errors?.town)}>
          <FormLabel>Town</FormLabel>
          <Controller
            name="town"
            control={control}
            rules={{ required: { value: true, message: 'Town is required' } }}
            render={({ field }) => (
              <Select
                {...field}
                isClearable
                isSearchable
                options={towns}
                onChange={(newV, _) => { setValue('town', newV); setValue('postalCode', newV?.postalCode) }}
                value={getValues()?.town}
                placeholder="Select town"
              />
            )}
          />
          {(errors.town != null) && <FormErrorMessage>{`${errors.town.message}`}</FormErrorMessage>}
          <FormHelperText>Which town makes your home?</FormHelperText>
        </FormControl>
        <FormControl>
          <FormLabel>Postal Code</FormLabel>
          <Input
            disabled
            {...register('postalCode')}
          />
        </FormControl>
      </VStack>
      <HStack mt={{ base: 4, md: 6 }}>
        <Button onClick={() => { setStep('description') }} colorScheme="green" leftIcon={<ArrowBackIcon />}>Go back</Button>
        <Spacer />
        <Button type="submit" colorScheme="green" rightIcon={<ArrowForwardIcon />}>Next</Button>
      </HStack>
    </form>
  )
}

export default Location
