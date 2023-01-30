import Head from 'next/head'

import { Heading } from '@chakra-ui/react'

function Listings() {
  return (
    <>
      <Head>
        <title>Search listings by town or postal code</title>
      </Head>
      <Heading>Property listings</Heading>
    </>
  )
}

export default Listings
