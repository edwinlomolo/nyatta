import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import { Box, Button, Center, FormControl, FormErrorMessage, FormLabel, Input, FormHelperText, Stack, Select as ChakraSelect } from '@chakra-ui/react'
import { yupResolver } from '@hookform/resolvers/yup'
import { Controller, useForm, type SubmitHandler, useFieldArray } from 'react-hook-form'
import Select from 'react-select'

import data from '../../../data/amenities.json'
import { defaultUnitsForm } from '../constants'
import { type UnitsForm } from '../types'
import { unitsSchema } from '../validations'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

const Units = () => {
  const { register, control, getValues, formState: { errors }, handleSubmit } = useForm<UnitsForm>({
    defaultValues: { ...defaultUnitsForm },
    resolver: yupResolver(unitsSchema)
  })
  const { setStep, setUnitsCount } = usePropertyOnboarding()
  const { fields, append, remove } = useFieldArray({
    control,
    name: 'units'
  })

  const { amenities } = data
  const onSubmit: SubmitHandler<UnitsForm> = data => console.log(data)
  const goBack = () => { setStep('caretaker') }
  const appendUnit = () => {
    append({ name: '', type: '', amenities: [] })
    setUnitsCount(getValues().units.length)
  }
  const removeUnit = (unitIndex: number) => {
    remove(unitIndex)
    setUnitsCount(getValues().units.length)
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack h="100%" mx={2}>
        {fields.map((_, unitIndex) => (
          <Stack direction={{ base: "column", md: "row" }} key={unitIndex}>
            <FormControl isInvalid={Boolean((errors?.units?.[unitIndex] as { name: object })?.name)}>
              <FormLabel>Name</FormLabel>
              <Input
                {...register(`units.${unitIndex}.name`)}
                placeholder="Name/ID"
              />
              {(((errors.units?.[unitIndex] as { name: object })?.name) != null) && <FormErrorMessage>{(errors.units?.[unitIndex] as { name: Partial<{ message: string }> })?.name?.message}</FormErrorMessage>}
              <FormHelperText>How you name your units</FormHelperText>
            </FormControl>
            <FormControl isInvalid={Boolean(errors?.units?.[unitIndex]?.type)}>
              <FormLabel>Type</FormLabel>
              <ChakraSelect {...register(`units.${unitIndex}.type`)} placeholder="Unit type">
                <option value="single room">Single room</option>
                <option value="studio">Studio</option>
                <option value="1">1 bedroom</option>
                <option value="2">2 bedroom</option>
                <option value="3">3 bedroom</option>
              </ChakraSelect>
              {((errors.units?.[unitIndex]?.type) != null) && <FormErrorMessage>{(errors.units?.[unitIndex]?.type as Partial<{ message: string }>)?.message}</FormErrorMessage>}
              <FormHelperText>Unit type</FormHelperText>
            </FormControl>
            <FormControl>
              <FormLabel>Amenities</FormLabel>
              {/* TODO filter out shared amenities */}
              <Controller
                name={`units.${unitIndex}.amenities`}
                control={control}
                render={({ field }) => (
                  <Select
                    {...field}
                    placeholder="Amenities"
                    options={amenities}
                    isMulti
                    menuPortalTarget={document.body}
                    styles={{ menuPortal: base => ({ ...base, zIndex: 1 }) }}
                  />
                )}
              />
              <FormHelperText>Amenities offered by this unit</FormHelperText>
            </FormControl>
            <Center>
              <Button onClick={() => removeUnit(unitIndex)} colorScheme="red" size="sm">Delete</Button>
            </Center>
          </Stack>
        ))}
        <Box mt={5}>
          <Button onClick={appendUnit} colorScheme="green">Add Unit</Button>
        </Box>
      </Stack>
      <Box zIndex={1} py={4} display="flex" bg="#ffff" justifyContent="space-between" w="100%" position="sticky" bottom="0" mt={{ base: 4, md: 6 }}>
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

export default Units
