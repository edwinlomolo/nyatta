import { useContext } from 'react'

import { OnboardingContext } from '../contexts/property-onboarding'
import { SearchListingContext } from '../contexts/search-listings'
import { SignInContext } from '../contexts/sign-in'

export const useSearchListings = () => useContext(SearchListingContext)
export const usePropertyOnboarding = () => useContext(OnboardingContext)
export const useSignIn = () => useContext(SignInContext)
