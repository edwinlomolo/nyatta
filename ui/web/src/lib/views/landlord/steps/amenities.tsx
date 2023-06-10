import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import { Box, Button, FormControl, FormLabel, FormErrorMessage, Stack } from '@chakra-ui/react'
import { Controller, useForm, type SubmitHandler } from 'react-hook-form'
import Select from 'react-select'

import data from '../../../data/amenities.json'
import { type AmenitiesForm } from '../types'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

const Amenities = (): JSX.Element => {
  const { setStep, amenitiesForm, setAmenitiesForm } = usePropertyOnboarding()
  const { control, handleSubmit, formState: { errors } } = useForm<AmenitiesForm>({
    defaultValues: amenitiesForm,
  })

  const goBack = () => setStep('location')
  const onSubmit: SubmitHandler<AmenitiesForm> = data => {
    setAmenitiesForm(data)
    setStep('pricing')
  }

  const { amenities } = data

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack direction={{ base: "column", md: "row"}}>
        <FormControl isInvalid={Boolean(errors?.amenities)}>
          <FormLabel>Select amenities</FormLabel>
          <Controller
            name="amenities"
            control={control}
            rules={{ required: { value: true, message: "Amenities required" } }}
            render={({ field }) => (
              <Select
                {...field}
                options={amenities}
                isMulti
              />
            )}
          />
          {((errors.amenities) != null) && <FormErrorMessage>{errors.amenities.message}</FormErrorMessage>}
        </FormControl>
      </Stack>
      <Box bg="#ffff" display="flex" mt={{ base: 4, md: 6 }} justifyContent="space-between">
        <Box>
          <Button onClick={goBack} colorScheme="green" leftIcon={<ArrowBackIcon />}>Go back</Button>
        </Box>
        <Box>
          <Button type="submit" colorScheme="green" rightIcon={<ArrowForwardIcon />}>Next</Button>
        </Box>
      </Box>
    </form>
  )
}

export default Amenities
