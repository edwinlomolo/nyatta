import { Text } from '@chakra-ui/react'

import { FormStepTitle } from '../constants'

import { usePropertyOnboarding } from '@usePropertyOnboarding'


const Title = () => {
  const { step, unitsCount } = usePropertyOnboarding()

  return (
    <Text fontSize={{ base: '2xl', md: '3xl' }}>
      {FormStepTitle[step]} {' '}
      {step === 'units' && unitsCount > 0 && <span>({unitsCount})</span>}
    </Text>
  )
}

export default Title
