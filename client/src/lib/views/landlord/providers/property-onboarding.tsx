import { useState, PropsWithChildren } from 'react'

import { useForm } from 'react-hook-form'

import { OnboardingContext } from '../context/property-onboarding'
import { OnboardingSteps } from '../types'

export const OnboardingProvider = ({ children }: PropsWithChildren) => {
  const [step, setStep] = useState<OnboardingSteps>('description')
  const { handleSubmit, formState, register } = useForm()

  return (
    <OnboardingContext.Provider
      value={{
        step,
        handleSubmit,
        register,
        formState,
        setStep,
      }}
    >
      {children}
    </OnboardingContext.Provider>
  )
}
