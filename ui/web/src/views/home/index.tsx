'use client'

import { AbsoluteCenter, Container, Box, Text } from '@chakra-ui/react'
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
        <Box bottom="0" left="0" w="100%" textAlign="center" position="fixed">
          <Text
            as="a"
            href="mailto:edwinmoses535@gmail.com"
            _hover={{
              cursor: 'pointer'
            }}
            textDecoration="underline"
          >
            Contact Us
          </Text>
        </Box>
      </Container>
    </>
  )

export default Home
