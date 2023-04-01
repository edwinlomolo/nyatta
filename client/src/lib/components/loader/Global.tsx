import { Box, Center, CircularProgress, Text } from '@chakra-ui/react'

interface Props {
  text?: string
}

function GlobalLoader({ text }: Props) {
  return (
    <Center>
      <Box textAlign="center">
        {text && <Text>{text}</Text>}
        <CircularProgress isIndeterminate />
      </Box>
    </Center>
  )
}

export default GlobalLoader
