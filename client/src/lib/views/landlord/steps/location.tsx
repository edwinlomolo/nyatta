import { Button, FormControl, FormLabel, FormErrorMessage, Input } from '@chakra-ui/react'

import { HStack, Spacer } from '@chakra-ui/react'

import { usePropertyOnboarding } from '../hooks/property-onboarding'

function Location() {
  const { handleSubmit, register, formState: { errors }, setStep } = usePropertyOnboarding()
  const onSubmit = (data: any) => console.log(data)

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl mb={5} isInvalid={!!errors.town}>
        <FormLabel>Town</FormLabel>
        <Input
          {...register("town", {
            pattern: {
                value: /^[A-Za-z ]+$/i,
                message: 'Should be a string value',
              },
            })
          }
        />
        {errors.town && <FormErrorMessage>{`${errors.town.message}`}</FormErrorMessage>}
      </FormControl>
      <FormControl mb={5} isInvalid={!!errors.postalCode}>
        <FormLabel>Postal Code</FormLabel>
        <Input
          {...register("postalCode", {
            pattern: {
                value: /^[0-9]{5}$/,
                message: 'Should be a 5-digit value',
              },
            })
          }
        />
        {errors.postalCode && <FormErrorMessage>{`${errors.postalCode.message}`}</FormErrorMessage>}
      </FormControl>
      <HStack>
        <Button onClick={() => setStep('description')} colorScheme="green">Back</Button>
        <Spacer />
        <Button colorScheme="green" type="submit">Create</Button>
      </HStack>
    </form>
  )
}

export default Location
