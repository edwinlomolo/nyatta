import Head from 'next/head'

import { useQuery } from '@apollo/client'

import { getListings } from '@gql'

import Search from './components/Search'
import NoListings from './components/NoListings'

function Listings() {
  const { data } = useQuery(getListings, {
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
      {data?.getListings.length === 0 && <NoListings />}
    </>
  )
}

export default Listings
