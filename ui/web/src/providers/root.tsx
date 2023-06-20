'use client'

import { ReactNode } from 'react'

import { ApolloProvider, type ApolloClient, type NormalizedCacheObject } from '@apollo/client'
import { CacheProvider } from '@chakra-ui/next-js'
import { ChakraProvider } from '@chakra-ui/react'
import { getCookie } from 'cookies-next'
import localFont from 'next/font/local'
import Head from 'next/head'

import { createClient } from '../apollo/createClient'
import { OnboardingProvider } from '../views/landlord/providers/property-onboarding'
import { SearchListingProvider } from '../views/listings/providers/search-listings'

import { theme } from '@styles'
import SignInProvider from 'providers/sign-in'

const mabryFont = localFont({ src: '../styles/assets/font/MabryPro-Regular.ttf' })

interface Props {
  children: ReactNode
}

const Providers = ({ children }: Props) => {
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
      <ApolloProvider client={client as ApolloClient<NormalizedCacheObject>}>
        <CacheProvider>
          <ChakraProvider theme={theme} cssVarsRoot="body">
            <SearchListingProvider>
              <OnboardingProvider>
                <SignInProvider>
                  {children}
                </SignInProvider>
              </OnboardingProvider>
            </SearchListingProvider>
          </ChakraProvider>
        </CacheProvider>
        </ApolloProvider>
    </main>
  )
}

export default Providers
