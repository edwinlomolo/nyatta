import { ApolloProvider, type ApolloClient } from '@apollo/client'
import { UserProvider } from '@auth0/nextjs-auth0/client'
import { ChakraProvider } from '@chakra-ui/react'
import { getCookie } from 'cookies-next'
import type { AppProps } from 'next/app'
import Head from 'next/head'

import { createClient } from '../lib/apollo/createClient'
import { OnboardingProvider } from '../lib/views/landlord/providers/property-onboarding'
import { SearchListingProvider } from '../lib/views/listings/providers/search-listings'




import { AuthProvider } from '@auth'
import Layout from '@layout'
import { theme } from '@styles'

export default function App ({ Component, pageProps }: AppProps) {
  const jwt = getCookie('jwt')
  const client = createClient(jwt)

  return (
    <UserProvider>
      <AuthProvider>
        <ApolloProvider client={client as ApolloClient<any>}>
          <ChakraProvider theme={theme} cssVarsRoot="body">
            <Head>
              <meta
                name="viewport"
                content="minimum-scale=1, initial-scale=1, width=device-width, shrink-to-fit=no, viewport-fit=cover"
              />
            </Head>
            <SearchListingProvider>
              <OnboardingProvider>
                <Layout>
                  <Component {...pageProps} />
                </Layout>
              </OnboardingProvider>
            </SearchListingProvider>
          </ChakraProvider>
        </ApolloProvider>
      </AuthProvider>
    </UserProvider>
  )
}
