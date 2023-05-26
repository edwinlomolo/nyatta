import Head from 'next/head'

import { usePropertyOnboarding } from '@usePropertyOnboarding'
import { Description, Location, Pricing, Units } from './steps'

function Landlord() {
  const { step } = usePropertyOnboarding()

  return (
    <div>
      <Head>
        <title>Manage your properties in one place</title>
      </Head>
      {step === 'description' && <Description />}
      {step === 'location' && <Location />}
      {step === 'pricing' && <Pricing />}
      {step === 'units' && <Units />}
    </div>
  )
}

export default Landlord
