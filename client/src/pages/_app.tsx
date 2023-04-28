import type { AppProps } from 'next/app'
import Head from 'next/head'
import { UserProvider } from '@auth0/nextjs-auth0/client'
import { getCookie } from 'cookies-next'
import { ChakraProvider } from '@chakra-ui/react'

import { theme } from '@styles'

import { SearchListingProvider } from '../lib/views/listings/providers/search-listings'

import { ApolloProvider, ApolloClient } from '@apollo/client'
import { AuthProvider } from '@auth'
import { useApolloClient } from '@apollo'
import Layout from '@layout'

export default function App({ Component, pageProps }: AppProps) {
  const jwt = getCookie('jwt')
  const { client } = useApolloClient({ jwt })

  return (
    <UserProvider>
      <AuthProvider>
        <ApolloProvider client={client as ApolloClient<any>}>
          <ChakraProvider theme={theme}>
            <Head>
              <meta
                name="viewport"
                content="minimum-scale=1, initial-scale=1, width=device-width, shrink-to-fit=no, viewport-fit=cover"
              />
            </Head>
            <SearchListingProvider>
              <Layout>
                <Component {...pageProps} />
              </Layout>
            </SearchListingProvider>
          </ChakraProvider>
        </ApolloProvider>
      </AuthProvider>
    </UserProvider>
  )
}
