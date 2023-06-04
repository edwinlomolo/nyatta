export type OnboardingStep = 'description' | 'location' | 'amenities' | 'pricing' | 'caretaker' | 'units'

export interface LocationOption {
  readonly label: string
  readonly value: string
  readonly postalCode: string
  readonly id: string
}

type PropertyType = 'Apartment' | 'Condominium' | 'Bungalow'

export interface FormValues {
  name: string
  propertyType: PropertyType | undefined
  minPrice: string
  maxPrice: string
  town: LocationOption
  postalCode: string
  units: Array<Record<string, any>>
}

export interface DescriptionForm {
  name: string
  propertyType: PropertyType | undefined
}

export interface LocationForm {
  town: LocationOption | null
  postalCode: string | undefined
}

export interface PriceForm {
  minPrice: number
  maxPrice: number
}

export interface CaretakerForm {
  firstName: string
  lastName: string
  phoneNumber: string
  idVerification: string
}

export interface UnitsForm {
  units: Array<Record<string, any>>
}
