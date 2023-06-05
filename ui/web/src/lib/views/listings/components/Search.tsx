import { Button, Flex, FormControl, FormErrorMessage, Input, Select as ChakraSelect } from '@chakra-ui/react'
import { Controller, type SubmitHandler } from 'react-hook-form'
import Select, { type GroupBase } from 'react-select'

import { usePropertyOnboarding } from '@usePropertyOnboarding'
import { useSearchListings } from '@usePropertySearch'

const Search = () => {
  const { towns } = usePropertyOnboarding()
  const { control, getListings, handleSubmit, register, formState: { errors } } = useSearchListings()
  const onSubmit: SubmitHandler<any> = async data => {
    await getListings({
      variables: {
        input: {
          town: data.town.label,
          minPrice: Number(data.minPrice),
          maxPrice: Number(data.maxPrice)
        }
      }
    })
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Flex
        p={5}
        gap={4}
        flexDirection={{ md: 'row', base: 'column' }}
      >
        <FormControl isInvalid={!(errors.town == null)}>
          <Controller
            name="town"
            rules={{ required: { value: true, message: 'This is required' } }}
            control={control}
            render={({ field }) => (
              <Select
                {...field}
                isClearable
                isSearchable
                options={(towns as unknown) as GroupBase<string>[]}
                placeholder="Town"
              />
            )}
          />
          {(errors.town != null) && <FormErrorMessage>{`${errors.town.message}`}</FormErrorMessage>}
        </FormControl>
        <FormControl isInvalid={!(errors.propertyType == null)}>
          <ChakraSelect {...register('propertyType', { required: 'Select property type' })} placeholder="Property type">
            <option value="single">Single room</option>
            <option value="studio">Studio</option>
            <option value="1">1 bedroom</option>
            <option value="2">2 bedrooms</option>
            <option value="3">3 bedrooms</option>
            <option value="4">4 bedrooms</option>
          </ChakraSelect>
          {(errors.propertyType != null) && <FormErrorMessage>{`${errors.propertyType.message}`}</FormErrorMessage>}
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
