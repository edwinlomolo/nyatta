import { ApolloProvider, type ApolloClient, type NormalizedCacheObject } from '@apollo/client'
import { UserProvider } from '@auth0/nextjs-auth0/client'
import { ChakraProvider } from '@chakra-ui/react'
import Layout from '@layout'
import localFont from '@next/font/local'
import { theme } from '@styles'
import { getCookie } from 'cookies-next'
import type { AppProps } from 'next/app'
import Head from 'next/head'

import { createClient } from '../apollo/createClient'
import { OnboardingProvider } from '../views/landlord/providers/property-onboarding'
import { SearchListingProvider } from '../views/listings/providers/search-listings'

const mabryFont = localFont({ src: '../styles/assets/font/MabryPro-Regular.ttf' })

const App = ({ Component, pageProps }: AppProps): JSX.Element => {
  const jwt = getCookie('jwt')
  const client = createClient(jwt)

  return (
    <main>
      <style jsx global>
        {`
          :root {
            --font-mabry: ${mabryFont.style.fontFamily};
          }
        `}
      </style>
      <Head>
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width, shrink-to-fit=no, viewport-fit=cover"
        />
      </Head>
      <UserProvider>
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
      </UserProvider>
    </main>
  )
}

export default App
