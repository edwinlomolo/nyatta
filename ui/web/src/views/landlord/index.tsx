'use client'

import { Container, HStack } from '@chakra-ui/react'
import Head from 'next/head'

import { Title } from './components'
import { Description, Location, Units, Caretaker, PropertyType, Shoot } from './steps'

import { usePropertyOnboarding } from '@usePropertyOnboarding'


const Landlord = () => {
  const { step } = usePropertyOnboarding()

  return (
    <Container>
      <Head>
        <title>Manage your properties in one place</title>
      </Head>
      <HStack my={{ base: 4, md: 6 }}>
        <Title />
      </HStack>
      {step === 'description' && <Description />}
      {step === 'location' && <Location />}
      {step === 'units' && <Units />}
      {step === 'caretaker' && <Caretaker />}
      {step === 'type' && <PropertyType />}
      {step === 'shoot' && <Shoot />}
    </Container>
  )
}

export default Landlord
