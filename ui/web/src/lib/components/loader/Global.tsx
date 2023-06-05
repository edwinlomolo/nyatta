import { Box, Center, CircularProgress, Text } from '@chakra-ui/react'

interface Props {
  text?: string
}

const GlobalLoader = ({ text }: Props) => (
    <Center>
      <Box textAlign="center">
        {text && <Text>{text}</Text>}
        <CircularProgress isIndeterminate />
      </Box>
    </Center>
  )

export default GlobalLoader
