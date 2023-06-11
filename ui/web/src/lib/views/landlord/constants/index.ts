import { type OnboardingStep, type DescriptionForm, type LocationForm, type CaretakerForm, type UnitsForm, type PriceForm, type AmenitiesForm } from '../types'

export const FormSteps: OnboardingStep[] = ['description', 'location', 'amenities', 'pricing', 'caretaker', 'units']

export const FormStepTitle: Record<OnboardingStep, string> = {
  description: 'Describe your property?',
  location: 'Property location',
  amenities: 'Shared amenities',
  pricing: 'How do you price your units?',
  caretaker: 'Who is the caretaker?',
  units: 'Add property units'
}

export const defaultDescriptionForm: DescriptionForm = {
  name: '',
  propertyType: undefined
}

export const defaultLocationForm: LocationForm = {
  town: null,
  postalCode: ''
}

export const defaultPriceForm: PriceForm = {
  minPrice: 0,
  maxPrice: 0
}

export const defaultCaretakerForm: CaretakerForm = {
  firstName: '',
  lastName: '',
  phoneNumber: '',
  idVerification: '',
  countryCode: ''
}

export const defaultUnitsForm: UnitsForm = {
  units: [{ name: '', type: '', baths: 0, amenities: [] }]
}

export const defaultAmenitiesForm: AmenitiesForm = {
  amenities: []
}
