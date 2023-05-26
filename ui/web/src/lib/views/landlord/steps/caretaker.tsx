import { Button, HStack, Spacer, VStack } from '@chakra-ui/react'
import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

function Caretaker() {
  const { handleSubmit, setStep } = usePropertyOnboarding()
  const onSubmit = () => {}
  const goBack = () => setStep("pricing")

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack spacing={{ base: 4, md: 6 }}>
      </VStack>
      <HStack>
        <Button colorScheme="green" leftIcon={<ArrowBackIcon />} onClick={goBack}>Go back</Button>
        <Spacer />
        <Button type="submit" colorScheme="green" rightIcon={<ArrowForwardIcon />}>Next</Button>
      </HStack>
    </form>
  )
}

export default Caretaker
