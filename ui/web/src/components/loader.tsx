'use client'

import { Center, Spinner } from '@chakra-ui/react'

const Loader = (): JSX.Element => (
  <Center>
    <Spinner size="xl" thickness="10px" />
  </Center>
)

export default Loader
