import React, { useContext } from 'react'

import { Route, Redirect } from 'react-router-dom'

import { AuthContext } from '../auth'

interface Props {
  children: React.ReactElement
  path: string
}

function PrivateRoute({ children, ...rest }: Props) {
  const { isAuthenticated } = useContext(AuthContext)

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
              state: { from: location },
            }}
          />
        )
      }
    />
  )
}

export default PrivateRoute
