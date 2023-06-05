import Head from 'next/head'

import NoListings from './components/NoListings'
import Search from './components/Search'

import { GlobalLoader } from '@components'
import { useSearchListings } from '@usePropertySearch'

const Listings = () => {
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
