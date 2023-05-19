import Head from 'next/head'

import { Flex } from '@chakra-ui/react'

import { usePropertyOnboarding } from '@usePropertyOnboarding'
import { Description, Location } from './steps'

function Landlord() {
  const { step } = usePropertyOnboarding()

  return (
    <Flex justifyContent="center">
      <Head>
        <title>Manage your properties in one place</title>
      </Head>
      {step === 'description' && <Description />}
      {step === 'location' && <Location />}
    </Flex>
  )
}

export default Landlord
