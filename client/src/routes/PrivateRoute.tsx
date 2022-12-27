import { Route, Redirect } from 'react-router-dom'
import { useAuth0 } from '@auth0/auth0-react'

function PrivateRoute({ children, ...rest }: any) {
  const { isAuthenticated } = useAuth0()

  return (
    <Route
      {...rest}
      render={({ location }) =>
        isAuthenticated ? (
          children
        ) : (
          <Redirect
            to={{
              pathname: "/",
              state: { from: location }
            }}
          />
        )
      }
    />
  )
}

export default PrivateRoute
