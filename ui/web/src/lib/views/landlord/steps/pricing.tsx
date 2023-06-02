import { Button, FormControl, FormErrorMessage, FormLabel, FormHelperText, Input, HStack, Spacer, VStack } from '@chakra-ui/react'
import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import { useForm, SubmitHandler } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

import { PriceForm } from '../types'
import { defaultPriceForm } from '../constants'
import { priceSchema } from '../validations'

function Pricing() {
  const { register, handleSubmit, formState: { errors } } = useForm<PriceForm>({
    defaultValues: defaultPriceForm,
    resolver: yupResolver(priceSchema),
  })
  const { setPriceForm, setStep } = usePropertyOnboarding()

  const onSubmit: SubmitHandler<PriceForm> = data => {
    setPriceForm(data)
    setStep("caretaker")
  }
  const goBack = () => setStep("location")

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack spacing={{ base: 4, md: 6 }}>
        <FormControl isInvalid={Boolean(errors?.minPrice)}>
          <FormLabel>Minimum Price</FormLabel>
          <Input
            {...register("minPrice")}
            type="number"
          />
          {errors?.minPrice && <FormErrorMessage>{`${errors?.minPrice.message}`}</FormErrorMessage>}
          <FormHelperText>This is the lowest priced unit</FormHelperText>
        </FormControl>
        <FormControl isInvalid={Boolean(errors?.maxPrice)}>
          <FormLabel>Maximum Price</FormLabel>
          <Input
            {...register("maxPrice")}
            type="number"
          />
          {errors?.maxPrice && <FormErrorMessage>{`${errors?.maxPrice.message}`}</FormErrorMessage>}
          <FormHelperText>This is the highest priced unit</FormHelperText>
        </FormControl>
      </VStack>
      <HStack mt={{ base: 4, md: 6 }}>
        <Button colorScheme="green" onClick={goBack} leftIcon={<ArrowBackIcon />}>Go back</Button>
        <Spacer />
        <Button colorScheme="green" type="submit" rightIcon={<ArrowForwardIcon />}>Next</Button>
      </HStack>
    </form>
  )
}

export default Pricing
