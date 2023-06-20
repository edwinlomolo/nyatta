import { useEffect } from 'react'

import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'
import { Accordion, AccordionButton, AccordionPanel, AccordionItem, HStack, Box, Button, FormControl, FormErrorMessage, FormLabel, Input, FormHelperText, Select as ChakraSelect, Text, Tag } from '@chakra-ui/react'
import { yupResolver } from '@hookform/resolvers/yup'
import { Select } from 'chakra-react-select'
import { Controller, useForm, type SubmitHandler, useFieldArray } from 'react-hook-form'

import { type UnitsForm } from '../types'

import { chakraStylesConfig } from '@styles'
import { usePropertyOnboarding } from '@usePropertyOnboarding'
import data from 'data/data.json'
import { UnitsSchema } from 'form/validations'

const Units = () => {
  const { setStep, setUnitsCount, unitsForm, setUnitsForm } = usePropertyOnboarding()
  const { register, control, clearErrors, getValues, setError, formState: { errors }, handleSubmit, watch } = useForm<UnitsForm>({
    defaultValues: unitsForm,
    resolver: yupResolver(UnitsSchema)
  })
  const { fields, append, remove } = useFieldArray({
    control,
    name: 'units'
  })

  const { amenities } = data
  const onSubmit: SubmitHandler<UnitsForm> = data => {
    // Get unit name
    const unitNames = data.units.map(unit => unit.name)
    // Duplicates
    const duplicateNames = unitNames.filter((unit, unitIndex) => unitNames.indexOf(unit) !== unitIndex)
    if (duplicateNames.length > 0) {
      duplicateNames.forEach(name => {
        const dupIndex = unitNames.lastIndexOf(name)
        setError(`units.${dupIndex}.name`, {
          type: "manual",
          message: "Unit name already taken"
        })
      })
      } else {
      clearErrors()
      setUnitsForm(data)
      setStep('shoot')
    }
  }
  const goBack = () => { setStep('caretaker') }
  const appendUnit = () => {
    append({ name: '', type: '', baths: 0, amenities: [], price: 0, bedrooms: [] })
    setUnitsCount(getValues().units.length)
  }
  const removeUnit = (unitIndex: number) => {
    remove(unitIndex)
    setUnitsCount(getValues().units.length)
  }

  const RenderBedrooms = ({ unitIndex, type }: any) => {
    const { fields, append, remove } = useFieldArray({ control, name: `units.${unitIndex}.bedrooms` })

    useEffect(() => {
      const totalBedrooms = Number(type)
      if (isNaN(totalBedrooms) && fields.length > 0) {
        for (let i = fields.length-1; i >= 0; i--) {
          remove(i)
        }
      } else if (totalBedrooms > fields.length) {
        for (let i = fields.length; i < totalBedrooms; i++) {
          append({ bedroomNumber: i+1, enSuite: "no", master: "no" })
        }
      } else if (totalBedrooms < fields.length) {
        for (let i = fields.length-1; i >= totalBedrooms; i--) {
          remove(i)
        }
      }
    }, [type, append, remove, fields.length])

    return (
      <Box mt={5}>
        {!!type && type !== 'studio' && type !== 'single room' && fields.length > 0 &&<Text>{`Bedrooms(${type})`}</Text>}
        {fields.length > 0 && fields.map((field, itemIndex) => (
          <HStack align="center" key={field.id}>
           <FormControl>
             <FormLabel>Bedroom Number</FormLabel>
             <Input
               {...register(`units.${unitIndex}.bedrooms.${itemIndex}.bedroomNumber`)}
               size="sm"
               disabled
               type="number"
               defaultValue={field.bedroomNumber}
             />
           </FormControl>
           <FormControl>
             <FormLabel>en-Suite</FormLabel>
             <ChakraSelect size="sm" {...register(`units.${unitIndex}.bedrooms.${itemIndex}.enSuite`)}>
               <option value="yes">Yes</option>
               <option value="no">No</option>
             </ChakraSelect>
           </FormControl>
           <FormControl>
             <FormLabel>Master</FormLabel>
             <ChakraSelect size="sm" {...register(`units.${unitIndex}.bedrooms.${itemIndex}.master`)}>
               <option value="yes">Yes</option>
               <option value="no">No</option>
             </ChakraSelect>
           </FormControl>
          </HStack>
        ))}
      </Box>
    )
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Accordion allowToggle h="100%">
        {fields.length > 0 && fields.map((unit, unitIndex: number) => {
          const type = watch(`units.${unitIndex}.type`)

          return (
          <AccordionItem w="100%" mb={4} gap={2} key={unit.id}>
            <AccordionButton>
              <Box as="span" flex="1" textAlign="left">
                Unit {`${unitIndex+1}`}
                {Boolean(errors?.units?.[unitIndex]) && <Tag colorScheme="red" mx={2}>Error</Tag>}
              </Box>
              <Text onClick={() => removeUnit(unitIndex)} textDecoration="underline" color="red">Delete</Text>
            </AccordionButton>
            <AccordionPanel>
              <Box display="flex" flexDirection={{ base: "column", md: "row" }} gap={2}>
                <FormControl isInvalid={Boolean((errors?.units?.[unitIndex] as { name: object })?.name)}>
                  <FormLabel>Name</FormLabel>
                  <Input
                    size="xs"
                    {...register(`units.${unitIndex}.name`, {
                    })}
                    placeholder="Name/ID"
                  />
                  {(((errors.units?.[unitIndex] as { name: object })?.name) != null) && <FormErrorMessage>{(errors.units?.[unitIndex] as { name: Partial<{ message: string }> })?.name?.message}</FormErrorMessage>}
                  <FormHelperText>How you name your units</FormHelperText>
                </FormControl>
                <FormControl isInvalid={Boolean(errors?.units?.[unitIndex]?.type)}>
                  <FormLabel>Type</FormLabel>
                  <ChakraSelect size="xs" {...register(`units.${unitIndex}.type`)} placeholder="Unit type">
                    <option value="single room">Single room</option>
                    <option value="studio">Studio</option>
                    <option value="1">1 bedroom</option>
                    <option value="2">2 bedroom</option>
                    <option value="3">3 bedroom</option>
                  </ChakraSelect>
                  {((errors.units?.[unitIndex]?.type) != null) && <FormErrorMessage>{(errors.units?.[unitIndex]?.type as Partial<{ message: string }>)?.message}</FormErrorMessage>}
                  <FormHelperText>Unit type</FormHelperText>
                </FormControl>
                <FormControl isInvalid={Boolean((errors?.units?.[unitIndex] as { baths: object })?.baths)}>
                  <FormLabel>Bathrooms</FormLabel>
                  <Input
                    {...register(`units.${unitIndex}.baths`)}
                    type="number"
                    size="xs"
                    placeholder="Bathrooms"
                  />
                  {(((errors.units?.[unitIndex] as { baths: object })?.baths) != null) && <FormErrorMessage>{(errors.units?.[unitIndex] as { baths: Partial<{ message: string }> })?.baths?.message}</FormErrorMessage>}
                  <FormHelperText>Total baths</FormHelperText>
                </FormControl>
              </Box>
              <Box display="flex" flexDirection={{ base: "column", md: "row" }} mt={4} gap={2}>
                <FormControl>
                  <FormLabel>Amenities</FormLabel>
                  <Controller
                    name={`units.${unitIndex}.amenities`}
                    control={control}
                    render={({ field }) => (
                      <Select
                        {...field}
                        size="sm"
                        placeholder="Amenities"
                        options={amenities}
                        isSearchable
                        isMulti
                        closeMenuOnSelect={false}
                        menuPortalTarget={document.body}
                        styles={{ menuPortal: base => ({ ...base, zIndex: 1 }) }}
                        chakraStyles={chakraStylesConfig}
                      />
                    )}
                  />
                  <FormHelperText>Amenities offered by this unit</FormHelperText>
                </FormControl>
                <FormControl isInvalid={Boolean((errors?.units?.[unitIndex] as { price: object })?.price)}>
                  <FormLabel>Price</FormLabel>
                  <Input
                    {...register(`units.${unitIndex}.price`, {
                      setValueAs: v => Number(v),
                    })}
                    placeholder="Monthly charge"
                    type="number"
                    size="xs"
                  />
                  {(((errors.units?.[unitIndex] as { price: object })?.price) != null) && <FormErrorMessage>{(errors.units?.[unitIndex] as { price: Partial<{ message: string }> })?.price?.message}</FormErrorMessage>}
                  <FormHelperText>How much will you charge monthly?</FormHelperText>
                </FormControl>
                
              </Box>
              <RenderBedrooms type={type} unitIndex={unitIndex} />
            </AccordionPanel>
          </AccordionItem>
        )})}
        <Box mt={5}>
          <Button size="sm" onClick={appendUnit} colorScheme="green">Add Unit</Button>
        </Box>
      </Accordion>
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
