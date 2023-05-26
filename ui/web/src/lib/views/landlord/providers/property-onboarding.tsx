import { useState, PropsWithChildren } from 'react'

import { useQuery } from '@apollo/client'

import { useForm } from 'react-hook-form'

import { getTowns as GET_TOWNS } from '@gql'
import { OnboardingContext } from '../context/property-onboarding'
import { OnboardingSteps } from '../types'

export const OnboardingProvider = ({ children }: PropsWithChildren) => {
  // For default towns select input
  const { data } = useQuery(GET_TOWNS)
  const locations = data?.getTowns.map((item: any) => ({
    id: item.id,
    value: item.town.toLowerCase(),
    label: item.town,
    postalCode: item.postalCode,
  }))

  const [step, setStep] = useState<OnboardingSteps>('units')
  const { control, getValues, reset, setValue, handleSubmit, formState, register } = useForm()

  return (
    <OnboardingContext.Provider
      value={{
        step,
        handleSubmit,
        register,
        formState,
        setStep,
        towns: locations,
        control,
        setValue,
        reset,
        getValues,
      }}
    >
      {children}
    </OnboardingContext.Provider>
  )
}
