
import { Container, HStack, Spacer, Show, Text } from '@chakra-ui/react'
import Head from 'next/head'

import { Title } from './components'
import { FormSteps } from './constants'
import { Description, Location, Pricing, Units, Caretaker, Amenities } from './steps'

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
        <Show above="md">
          <Spacer />
          <Text fontSize="4xl">{`${FormSteps.indexOf(step) + 1}/${FormSteps.length}`}</Text>
        </Show>
      </HStack>
      {step === 'description' && <Description />}
      {step === 'location' && <Location />}
      {step === 'amenities' && <Amenities />}
      {step === 'pricing' && <Pricing />}
      {step === 'units' && <Units />}
      {step === 'caretaker' && <Caretaker />}
    </Container>
  )
}

export default Landlord
