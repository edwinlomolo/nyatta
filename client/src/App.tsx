import { ChakraProvider } from '@chakra-ui/react'
import { CookiesProvider } from 'react-cookie'

import { AuthProvider } from './auth'
import { ApolloProvider } from './apollo'
import { RootRouter } from './routes'

import { GlobalStyle, theme } from './theme'

function App() {
  return (
    <CookiesProvider>
      <AuthProvider>
        <GlobalStyle />
        <ApolloProvider>
          <ChakraProvider theme={theme}>
            <RootRouter />
          </ChakraProvider>
        </ApolloProvider>
      </AuthProvider>
    </CookiesProvider>
  );
}

export default App;
