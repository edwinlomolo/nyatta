import { PropsWithChildren } from 'react'

import { Box } from '@chakra-ui/react'

function Main({ children }: PropsWithChildren) {
  return (
    <Box>
      {children}
    </Box>
  )
}

export default Main
