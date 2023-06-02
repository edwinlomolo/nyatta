export type OnboardingStep = 'description' | 'location' | 'amenities' | 'pricing' | 'caretaker' | 'units'

export type LocationOption = {
  readonly label: string
  readonly value: string
  readonly postalCode: string
  readonly id: string
}

type PropertyType = "Apartment" | "Condominium" | "Bungalow"

export type FormValues = {
  name: string
  propertyType: PropertyType | undefined
  minPrice: string
  maxPrice: string
  town: LocationOption
  postalCode: string
  units: Record<string,any>[]
}

export type DescriptionForm = {
  name: string
  propertyType: PropertyType | undefined
}

export type LocationForm = {
  town: LocationOption | null
  postalCode: string | undefined
}

export type PriceForm = {
  minPrice: number
  maxPrice: number
}

export type CaretakerForm = {
  firstName: string
  lastName: string
  phoneNumber: string
  idVerification: string
}

export type UnitsForm = {
  units: Record<string,any>[]
}
