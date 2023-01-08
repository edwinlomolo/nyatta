import { useContext } from 'react'

import { AuthContext } from '../../auth'
import { ListingsPage } from '../../pages'

function HomePage() {
  const { isAuthenticated } = useContext(AuthContext)

  return isAuthenticated ? (
    <ListingsPage />
  ) : (
    <div>Home page</div>
  )
}

export default HomePage
