import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import { Button, FormControl, FormErrorMessage, FormLabel, FormHelperText, Input, HStack, Spacer, VStack } from '@chakra-ui/react'
import { yupResolver } from '@hookform/resolvers/yup'
import { useForm, type SubmitHandler } from 'react-hook-form'


import { type PriceForm } from '../types'
import { priceSchema } from '../validations'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

const Pricing = () => {
  const { priceForm, setPriceForm, setStep } = usePropertyOnboarding()
  const { register, handleSubmit, formState: { errors } } = useForm<PriceForm>({
    defaultValues: priceForm,
    resolver: yupResolver(priceSchema)
  })

  const onSubmit: SubmitHandler<PriceForm> = data => {
    setPriceForm(data)
    setStep('caretaker')
  }
  const goBack = () => { setStep('amenities') }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack spacing={{ base: 4, md: 6 }}>
        <FormControl isInvalid={Boolean(errors?.minPrice)}>
          <FormLabel>Minimum Price</FormLabel>
          <Input
            {...register('minPrice')}
            type="number"
          />
          {((errors?.minPrice) != null) && <FormErrorMessage>{`${errors?.minPrice.message}`}</FormErrorMessage>}
          <FormHelperText>This is the lowest priced unit</FormHelperText>
        </FormControl>
        <FormControl isInvalid={Boolean(errors?.maxPrice)}>
          <FormLabel>Maximum Price</FormLabel>
          <Input
            {...register('maxPrice')}
            type="number"
          />
          {((errors?.maxPrice) != null) && <FormErrorMessage>{`${errors?.maxPrice.message}`}</FormErrorMessage>}
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
