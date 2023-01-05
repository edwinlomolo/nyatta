import { ChakraProvider } from '@chakra-ui/react'
import { CookiesProvider } from 'react-cookie'
import { Switch } from 'react-router-dom'

import { AuthProvider } from './auth'
import { ApolloProvider } from './apollo'
import { UserHome } from './components'
import { Main } from './layout'
import { PrivateRoute, RouteWithLayout } from './routes'

function App() {
  return (
    <CookiesProvider>
      <AuthProvider>
        <ApolloProvider>
          <ChakraProvider>
            <Switch>
              <RouteWithLayout
                layout={Main}
                component={UserHome}
                path="/"
              />
              <PrivateRoute path="/">
                <Switch>
                  <RouteWithLayout
                    layout={Main}
                    component={UserHome}
                    path="/"
                  />
                </Switch>
              </PrivateRoute>
            </Switch>
          </ChakraProvider>
        </ApolloProvider>
      </AuthProvider>
    </CookiesProvider>
  );
}

export default App;
