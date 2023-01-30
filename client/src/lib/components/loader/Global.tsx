import { Center, CircularProgress } from '@chakra-ui/react'

function GlobalLoader() {
  return (
    <Center>
      <CircularProgress isIndeterminate />
    </Center>
  )
}

export default GlobalLoader
