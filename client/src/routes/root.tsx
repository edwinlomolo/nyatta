import { Switch, Route } from 'react-router-dom'

import { Box } from '@chakra-ui/react'

import RouteWithLayout from './RouteWithLayout'
import PrivateRoute from './PrivateRoute'
import { LandingPage, NotFoundPage } from '../pages'
import { Main } from '../layout'

function RootRouter() {
  return (
    <Box>
      <Switch>
        <RouteWithLayout
          layout={Main}
          component={LandingPage}
          path="/"
          exact
        />
        <PrivateRoute path="/">
          <Switch>
            <RouteWithLayout
              layout={Main}
              component={() => <>Onboard property</>}
              path="/onboard"
              exact
            />
          </Switch>
        </PrivateRoute>
        
        <Route path="*" component={NotFoundPage} />
      </Switch>
    </Box>
  )
}

export default RootRouter
