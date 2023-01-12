import { useContext } from 'react'

import { AuthContext } from '../../auth'

import HomePage from './components/Home'
import Footer from './components/Footer'
import ListingsPage from '../listings/ListingsPage'

function LandingPage() {
  const { isAuthenticated } = useContext(AuthContext)

  return isAuthenticated ? (
    <ListingsPage />
  ) : (
    <>
      <HomePage />
      <Footer />
    </>
  )
}

export default LandingPage
