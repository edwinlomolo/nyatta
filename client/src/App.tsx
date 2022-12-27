import { useAuth0 } from '@auth0/auth0-react'

import { ChakraProvider } from '@chakra-ui/react'
import { CookiesProvider } from 'react-cookie'
import { Switch } from 'react-router-dom'

import { UserHome } from './components'
import { Main } from './layout'
import { PublicRoute } from './routes'

function App() {
  const { isLoading } = useAuth0()

  if (isLoading) {
    return <p>Loading ...</p>
  }

  return (
    <CookiesProvider>
      <ChakraProvider>
        <Switch>
          <PublicRoute
            layout={Main}
            component={UserHome}
            path="/"
          />
        </Switch>
      </ChakraProvider>
    </CookiesProvider>
  );
}

export default App;
