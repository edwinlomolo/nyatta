import { useContext } from 'react'

import { OnboardingContext } from '../context/property-onboarding'

export const usePropertyOnboarding = () => useContext(OnboardingContext)
