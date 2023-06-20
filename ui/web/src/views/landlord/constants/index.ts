import { type OnboardingStep, type DescriptionForm, type LocationForm, type CaretakerForm, type UnitsForm, type PriceForm, type AmenitiesForm, type PropertyTypeForm, type ContactPersonForm } from '../types'

export const FormSteps: OnboardingStep[] = ['description', 'type', 'location', 'caretaker', 'units', 'shoot']

export const FormStepTitle: Record<OnboardingStep, string> = {
  description: 'How can you name this property?',
  location: 'Describe your property by locality.',
  caretaker: 'We will schedule a professional shoot for your property so there has to be someone who will give us access to your property and guide us through the property. Additionally, this will be the immediate and authenticated goto person for any queries from users when your listing goes live.',
  units: "You've come this far! How best can you describe your units?",
  type: 'How best can you define your property?',
  shoot: 'Schedule a professional shoot.',
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
  units: [{ name: '', type: '', baths: 0, amenities: [], price: 0, bedrooms: [] }]
}

export const defaultAmenitiesForm: AmenitiesForm = {
  amenities: []
}

export const defaultPropertyType: PropertyTypeForm = {
  propertyType: undefined
}

export const defaultContactPerson: ContactPersonForm = {
  contactPerson: undefined,
  shootDate: undefined
}
