import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import { Button, HStack, Spacer } from '@chakra-ui/react'
import { yupResolver } from '@hookform/resolvers/yup'
import { SubmitHandler, useForm } from 'react-hook-form'

import { PropertyTypeForm } from '../types'

import { usePropertyOnboarding } from '@usePropertyOnboarding'
import FormRadio from 'components/form-radio'
import data from 'data/data.json'
import { PropertyTypeSchema } from 'form/validations'

const PropertyType = (): JSX.Element => {
  const { setStep, propertyType, setPropertyType } = usePropertyOnboarding()
  const { control, handleSubmit } = useForm<PropertyTypeForm>({
    defaultValues: propertyType,
    resolver: yupResolver(PropertyTypeSchema),
  })

  const goBack = () => setStep("description")
  const onSubmit: SubmitHandler<PropertyTypeForm> = data => {
    setPropertyType(data)
    setStep('location')
  }

  const { types: propertyTypes } = data

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormRadio
        control={control}
        options={propertyTypes}
        name="propertyType"
      />
      <HStack spacing={{ base: 4, md: 6 }}>
        <Button onClick={goBack} colorScheme="green" leftIcon={<ArrowBackIcon />}>Go back</Button>
        <Spacer />
        <Button type="submit" colorScheme="green" rightIcon={<ArrowForwardIcon />}>Done</Button>
      </HStack>
    </form>
  )
}

export default PropertyType
