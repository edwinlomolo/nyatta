import { FC } from 'react'

import { Route, RouteProps } from 'react-router-dom'

type RouteWithLayout = RouteProps & {
  layout: FC<any>,
  component: FC<any>,
}

function PublicRoute(props: RouteWithLayout) {
  const { layout: Layout, component: Component, ...rest } = props

  return (
    <Route
      {...rest}
      render={matchProps => (
        <Layout>
          <Component {...matchProps} />
        </Layout>
      )}
    />
  )
}

export default PublicRoute
