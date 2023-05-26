import { Button, FormControl, FormLabel, FormErrorMessage, FormHelperText, Input, VStack, HStack, Spacer } from '@chakra-ui/react'
import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import Select from 'react-select'
import { Controller } from 'react-hook-form'


import { usePropertyOnboarding } from '../hooks/property-onboarding'

function Location() {
  const { control, towns, setStep, setValue, getValues, handleSubmit, register, formState: { errors } } = usePropertyOnboarding()
  const onSubmit = () => setStep("pricing")

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack spacing={{ base: 4, md: 10 }}>
        <FormControl isInvalid={Boolean(errors?.town)}>
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
          <FormHelperText>Which town makes your home?</FormHelperText>
          {errors.town && <FormErrorMessage>{`${errors.town.message}`}</FormErrorMessage>}
        </FormControl>
        <FormControl>
          <FormLabel>Postal Code</FormLabel>
          <Input
            disabled
            {...register("postalCode")}
          />
        </FormControl>
      </VStack>
      <HStack mt={{ base: 4, md: 6 }}>
        <Button onClick={() => setStep("description")} colorScheme="green" leftIcon={<ArrowBackIcon />}>Go back</Button>
        <Spacer />
        <Button type="submit" colorScheme="green" rightIcon={<ArrowForwardIcon />}>Next</Button>
      </HStack>
    </form>
  )
}

export default Location
