import Head from 'next/head'

import { GlobalLoader } from '@components'
import Search from './components/Search'
import NoListings from './components/NoListings'

import { useSearchListings } from '@usePropertySearch'

function Listings () {
  const { listingsData, listingsLoading } = useSearchListings()

  return (
    <>
      <Head>
        <title>Search listings by town or postal code</title>
      </Head>
      <Search />
      {listingsLoading && <GlobalLoader />}
      {listingsData?.getListings.length === 0 && !listingsLoading && <NoListings />}
    </>
  )
}

export default Listings
