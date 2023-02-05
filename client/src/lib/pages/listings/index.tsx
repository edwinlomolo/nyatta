import Head from 'next/head'

import Search from './components/Search'

function Listings() {
  return (
    <>
      <Head>
        <title>Search listings by town or postal code</title>
      </Head>
      <Search />
    </>
  )
}

export default Listings
