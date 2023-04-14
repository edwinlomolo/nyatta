import Head from 'next/head'

import { useQuery } from '@apollo/client'
import {HelloDocument} from '../../../gql/graphql'

import Search from './components/Search'

function Listings() {
  const { data } = useQuery(HelloDocument)

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
