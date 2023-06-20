'use client'

import { VStack } from '@chakra-ui/react'
import Head from 'next/head'

import Footer from './components/Footer'
import HomeHeader from './components/HomeHeader'

const Home = () => (
    <VStack>
      <Head>
        <title>Nyatta - Find homes or apartments for rent.</title>
      </Head>
      <HomeHeader />
      <Footer />
    </VStack>
  )

export default Home
