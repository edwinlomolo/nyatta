import { ArrowBackIcon } from '@chakra-ui/icons'
import { Box, Button, FormControl, FormErrorMessage, FormLabel, HStack, Input, FormHelperText, VStack } from '@chakra-ui/react'
import { useForm, type SubmitHandler, useFieldArray } from 'react-hook-form'

import { defaultUnitsForm } from '../constants'
import { UnitsForm } from '../types'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

const Units = () => {
  const { register, control, formState: { errors }, handleSubmit } = useForm<UnitsForm>({
    defaultValues: { ...defaultUnitsForm },
    mode: "onChange"
  })
  const { setStep } = usePropertyOnboarding()
  const { fields, append } = useFieldArray({
    control,
    name: 'units'
  })
  const onSubmit: SubmitHandler<UnitsForm> = data => console.log(data)
  const goBack = () => { setStep('caretaker') }
  const appendUnit = () => {
    append({ name: 'Harambe' })
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack overflowY="auto" h="20vh" spacing={{ base: 4, md: 6 }}>
        {fields.map((_, unitIndex) => (
          <Box w="100%" gap={4} key={unitIndex} >
            <FormControl isInvalid={Boolean(errors?.units)}>
              <FormLabel>Unit name</FormLabel>
              <Input
                size="sm"
                {...register(`units.${unitIndex}.name`)}
                placeholder="Name/ID"
              />
              {((errors?.units) != null) && <FormErrorMessage>{`${errors?.units.message}`}</FormErrorMessage>}
              <FormHelperText>How do you name your units?</FormHelperText>
            </FormControl>
          </Box>
        ))}
      </VStack>
      <HStack mt={{ base: 4, md: 6 }}>
        <Button onClick={appendUnit} colorScheme="green">Add Unit</Button>
        <Button onClick={goBack} colorScheme="green" leftIcon={<ArrowBackIcon />}>Go back</Button>
      </HStack>
    </form>
  )
}

export default Units
