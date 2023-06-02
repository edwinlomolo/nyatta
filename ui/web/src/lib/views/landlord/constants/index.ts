import { OnboardingStep, DescriptionForm, LocationForm, CaretakerForm, UnitsForm, PriceForm } from '../types'

export const FormSteps: OnboardingStep[] = ['description', 'location', 'pricing', 'caretaker', 'units']

export const FormStepTitle: Record<OnboardingStep, string> = {
  description: 'Describe your property?',
  location: 'Property location',
  amenities: 'What does your property offer?',
  pricing: 'How do you price your units?',
  caretaker: 'Who is the caretaker?',
  units: 'Add property units',
}

export const defaultDescriptionForm: DescriptionForm = {
  name: "",
  propertyType: undefined,
}

export const defaultLocationForm: LocationForm = {
  town: null,
  postalCode: "",
}

export const defaultPriceForm: PriceForm = {
  minPrice: 0,
  maxPrice: 0,
}

export const defaultCaretakerForm: CaretakerForm = {
  firstName: "",
  lastName: "",
  phoneNumber: "",
  idVerification: "",
}

export const defaultUnitsForm: UnitsForm = {
  units: [],
}
