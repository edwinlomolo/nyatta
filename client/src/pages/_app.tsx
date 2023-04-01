import type { AppProps } from 'next/app'
import Head from 'next/head'
import { UserProvider } from '@auth0/nextjs-auth0/client'
import { getCookie } from 'cookies-next'

import { Chakra } from '../lib/components'

import { ApolloProvider } from '@apollo/client'
import { AuthProvider } from '../lib/auth'
import { createClient } from '../lib/apollo'
import Layout from '../lib/layout'

export default function App({ Component, pageProps }: AppProps) {
  const client = createClient(getCookie('jwt'))

  return (
    <UserProvider>
      <AuthProvider>
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
      </AuthProvider>
    </UserProvider>
  )
}
