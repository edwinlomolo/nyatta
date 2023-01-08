import React from 'react'

import { Route, RouteProps } from 'react-router-dom'

type Props = RouteProps & {
  layout: React.FC<any>
  component: React.FC<any>
}

function RouteWithLayout(props: Props) {
  const { layout: Layout, component: Component, location, ...rest } = props

  return (
    <Route
      {...rest}
      location={location}
      render={matchProps => (
        <Layout>
          <Component {...matchProps} />
        </Layout>
      )}
    />
  )
}

export default RouteWithLayout
