'use client'

import { AbsoluteCenter, Container, Text } from '@chakra-ui/react'
import Head from 'next/head'

const Home = () => (
    <>
      <Head>
        <title>Nyatta - Find homes or apartments for rent.</title>
      </Head>
      <Container maxW="full">
        <AbsoluteCenter w="100%">
          <Text fontSize={{base:"5xl", md:"7xl"}} textAlign="center">
            Find <span style={{color: 'white', background: '#276749'}}>local</span> rental homes
          </Text>
        </AbsoluteCenter>
      </Container>
    </>
  )

export default Home
