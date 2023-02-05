import { useForm } from 'react-hook-form'

import { Button, Flex, FormControl, FormErrorMessage, Input, Select } from '@chakra-ui/react'

function Search() {
  const { handleSubmit, register, formState: { errors } } = useForm()
  const onSubmit = (data: any) => {
    console.log(data)
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Flex
        p={5}
        gap={4}
        flexDirection={{ md: "row", base: "column" }}
      >
        <FormControl isInvalid={!!errors.town}>
          <Input
            {...register('town', {
              required: 'Town is required',
              pattern: {
                value: /^[A-Za-z ]+$/i,
                message: 'Should be a string value',
              },
            })}
            placeholder="Town"
          />
          {errors.town && <FormErrorMessage>{`${errors.town.message}`}</FormErrorMessage>}
        </FormControl>
        <FormControl isInvalid={!!errors.propertyType}>
          <Select {...register('propertyType', { required: 'Select property type' })} placeholder="Property type">
            <option value="single">Single room</option>
            <option value="studio">Studio</option>
            <option value="1">1 bedroom</option>
            <option value="2">2 bedrooms</option>
            <option value="3">3 bedrooms</option>
            <option value="4">4 bedrooms</option>
          </Select>
          {errors.propertyType && <FormErrorMessage>{`${errors.propertyType.message}`}</FormErrorMessage>}
        </FormControl>
        <FormControl>
          <Input
            {...register('minPrice')}
            type="number"
            placeholder="Min price"
          />
        </FormControl>
        <FormControl>
          <Input
            {...register('maxPrice')}
            type="number"
            placeholder="Max price"
          />
        </FormControl>
        <Flex>
          <Button w="100%" type="submit" colorScheme="green">
            Search
          </Button>
        </Flex>
      </Flex>
    </form>
  )
}

export default Search
