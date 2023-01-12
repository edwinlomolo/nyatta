import { Switch, Route } from 'react-router-dom'

import RouteWithLayout from './RouteWithLayout'
import PrivateRoute from './PrivateRoute'
import { LandingPage, NotFoundPage } from '../pages'
import { Main } from '../layout'

function RootRouter() {
  return (
    <>
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
    </>
  )
}

export default RootRouter
