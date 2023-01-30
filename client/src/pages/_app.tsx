import type { AppProps } from 'next/app'
import Head from 'next/head'
import { UserProvider } from '@auth0/nextjs-auth0/client'

import { Chakra } from '../lib/components'

import { ApolloProvider } from '@apollo/client'
import { createClient } from '../lib/apollo'
import Layout from '../lib/layout'

const client = createClient()

export default function App({ Component, pageProps }: AppProps) {
  return (
    <UserProvider>
      <ApolloProvider client={client}>
        <Chakra>
          <Head>
            <meta
              name="viewport"
              content="minimum-scale=1, initial-scale=1, width=device-width, shrink-to-fit=no, viewport-fit=cover"
            />
          </Head>
          <Layout>
            <Component {...pageProps} />
          </Layout>
        </Chakra>
      </ApolloProvider>
    </UserProvider>
  )
}
