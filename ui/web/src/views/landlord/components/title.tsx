import { Box, Text } from '@chakra-ui/react'

import { FormStepTitle } from '@constants'
import { usePropertyOnboarding } from '@hooks'


const Title = () => {
  const { step, unitsCount } = usePropertyOnboarding()

  return (
    <Box>
      <Text>
        {FormStepTitle[step]} {' '}
        {step === 'units' && unitsCount > 0 && <span>({unitsCount})</span>}
      </Text>
    </Box>
  )
    
}

export default Title
