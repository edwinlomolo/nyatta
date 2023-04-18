import Head from 'next/head'

import { useQuery } from '@apollo/client'

import { getListings } from '@gql'

import { GlobalLoader } from '@components'
import Search from './components/Search'
import NoListings from './components/NoListings'

function Listings() {
  const { data, loading } = useQuery(getListings, {
    variables: {
      input: {
        town: "Ngong Hills",
        minPrice: 0,
        maxPrice: 0,
      },
    },
  })

  return (
    <>
      <Head>
        <title>Search listings by town or postal code</title>
      </Head>
      <Search />
      {loading && <GlobalLoader />}
      {data?.getListings.length === 0 && !loading && <NoListings />}
    </>
  )
}

export default Listings
