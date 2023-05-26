import { Text } from '@chakra-ui/react'

import { usePropertyOnboarding } from '@usePropertyOnboarding'

import { FormSteps, FormStepTitle } from '../constants'

function Title() {
  const { step } = usePropertyOnboarding()

  return(
    <Text fontSize={{ base: "2xl", md: "3xl" }}>{FormStepTitle[step]}</Text>
  )
}

export default Title
