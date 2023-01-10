import { useContext } from 'react'

import { AuthContext } from '../../auth'

import HomePage from './Home'
import ListingsPage from '../listings/ListingsPage'

function LandingPage() {
  const { isAuthenticated } = useContext(AuthContext)

  return isAuthenticated ? (
    <ListingsPage />
  ) : (
    <HomePage />
  )
}

export default LandingPage
