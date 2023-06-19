import { ChakraProvider } from '@chakra-ui/react'

import { theme } from '../styles/theme'

interface ChakraProps {
  children: React.ReactNode
}

export const Chakra = ({ children }: ChakraProps) => <ChakraProvider theme={theme}>{children}</ChakraProvider>
