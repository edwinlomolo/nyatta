import { ApolloProvider, type ApolloClient, type NormalizedCacheObject } from '@apollo/client'
import { UserProvider } from '@auth0/nextjs-auth0/client'
import { ChakraProvider } from '@chakra-ui/react'
import localFont from '@next/font/local'
import { getCookie } from 'cookies-next'
import type { AppProps } from 'next/app'
import Head from 'next/head'

import { createClient } from '../lib/apollo/createClient'
import { OnboardingProvider } from '../lib/views/landlord/providers/property-onboarding'
import { SearchListingProvider } from '../lib/views/listings/providers/search-listings'

const mabroFont = localFont({ src: '../lib/styles/assets/font/MabryPro-Regular.ttf' })

import { AuthProvider } from '@auth'
import Layout from '@layout'
import { theme } from '@styles'

const App = ({ Component, pageProps }: AppProps): JSX.Element => {
  const jwt = getCookie('jwt')
  const client = createClient(jwt)

  return (
    <main className={mabroFont.className}>
      <Head>
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width, shrink-to-fit=no, viewport-fit=cover"
        />
      </Head>
      <UserProvider>
        <AuthProvider>
          <ApolloProvider client={client as ApolloClient<NormalizedCacheObject>}>
            <ChakraProvider theme={theme} cssVarsRoot="body">
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
    </main>
  )
}

export default App
