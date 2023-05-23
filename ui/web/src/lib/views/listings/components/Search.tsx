import { Controller } from 'react-hook-form'
import { Button, Flex, FormControl, FormErrorMessage, Input, Select as ChakraSelect } from '@chakra-ui/react'
import Select from 'react-select'

import { usePropertyOnboarding } from '@usePropertyOnboarding'
import { useSearchListings } from '@usePropertySearch'

function Search() {
  const { towns } = usePropertyOnboarding()
  const { control, getListings, handleSubmit, register, formState: { errors } } = useSearchListings()
  const onSubmit = async (data: any) => {
    await getListings({
      variables: {
        input: {
          town: data.town.label,
          minPrice: Number(data.minPrice),
          maxPrice: Number(data.maxPrice),
        },
      },
    })
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Flex
        p={5}
        gap={4}
        flexDirection={{ md: "row", base: "column" }}
      >
        <FormControl isInvalid={!!errors.town}>
          <Controller
            name="town"
            rules={{ required: { value: true, message: "This is required" } }}
            control={control}
            render={({ field }) => (
              <Select
                {...field}
                isClearable
                isSearchable
                options={towns}
                placeholder="Town"
              />
            )}
          />
          {errors.town && <FormErrorMessage>{`${errors.town.message}`}</FormErrorMessage>}
        </FormControl>
        <FormControl isInvalid={!!errors.propertyType}>
          <ChakraSelect {...register('propertyType', { required: 'Select property type' })} placeholder="Property type">
            <option value="single">Single room</option>
            <option value="studio">Studio</option>
            <option value="1">1 bedroom</option>
            <option value="2">2 bedrooms</option>
            <option value="3">3 bedrooms</option>
            <option value="4">4 bedrooms</option>
          </ChakraSelect>
          {errors.propertyType && <FormErrorMessage>{`${errors.propertyType.message}`}</FormErrorMessage>}
        </FormControl>
        <FormControl>
          <Input
            {...register('minPrice')}
            type="number"
            placeholder="Min price"
            defaultValue="0"
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