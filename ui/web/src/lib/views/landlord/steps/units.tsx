import { ArrowBackIcon } from '@chakra-ui/icons'
import { Box, Button, FormControl, FormErrorMessage, FormLabel, HStack, Input, FormHelperText, VStack } from '@chakra-ui/react'
import { useFieldArray } from 'react-hook-form'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

const Units = () => {
  const { control, register, setStep, formState: { errors }, handleSubmit } = usePropertyOnboarding()
  const { fields, append } = useFieldArray({
    control,
    name: 'units'
  })
  const onSubmit = (data: any) => { console.log(data) }
  const goBack = () => { setStep('pricing') }
  const appendUnit = () => {
    append({ name: 'Harambe' })
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack overflowY="auto" h="20vh" spacing={{ base: 4, md: 6 }}>
        {fields.map((unit, unitIndex) => (
          <Box w="100%" gap={4} key={unitIndex} >
          <FormControl isInvalid={Boolean(errors?.units)}>
            <FormLabel>Unit name</FormLabel>
            <Input
              size="sm"
              {...register('units', { required: 'Unit name is required' })}
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
