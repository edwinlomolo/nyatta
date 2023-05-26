import { Button, Container, FormControl, FormErrorMessage, FormLabel, HStack, Input, FormHelperText, VStack, Spacer } from '@chakra-ui/react'
import { ArrowForwardIcon, ArrowBackIcon } from '@chakra-ui/icons'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

function Units() {
  const { register, setStep, formState: { errors }, handleSubmit } = usePropertyOnboarding()
  const onSubmit = (data: any) => console.log(data)
  const goBack = () => setStep("pricing")

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Container>
        <VStack spacing={{ base: 4, md: 10 }}>
          <FormControl isInvalid={Boolean(errors?.unitName)}>
            <FormLabel>How do you name this unit?</FormLabel>
            <Input
              {...register("unitName", { required: "Unit name is required" })}
              placeholder="Unit name/ID"
            />
            {errors?.unitName && <FormErrorMessage>{`${errors?.unitName.message}`}</FormErrorMessage>}
            <FormHelperText>This can be C4/Kilimanjaro- how do you name your units?</FormHelperText>
          </FormControl>
          <HStack>
            <Button onClick={goBack} colorScheme="green" leftIcon={<ArrowBackIcon />}>Go back</Button>
            <Spacer />
            <Button type="submit" rightIcon={<ArrowForwardIcon />} colorScheme="green">Create Unit</Button>
          </HStack>
        </VStack>
      </Container>
    </form>
  )
}

export default Units
