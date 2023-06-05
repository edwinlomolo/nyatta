import { Text } from '@chakra-ui/react'

import { FormStepTitle } from '../constants'

import { usePropertyOnboarding } from '@usePropertyOnboarding'


const Title = () => {
  const { step, getValues } = usePropertyOnboarding()
  const { units } = getValues()

  return (
    <Text fontSize={{ base: '2xl', md: '3xl' }}>
      {FormStepTitle[step]} {' '}
      {units?.length > 0 && <span>({units?.length})</span>}
    </Text>
  )
}

export default Title
