export type Status = 'pending' | 'approved' | 'sign-in'
export interface SignInForm {
  phone: number | undefined
  countryCode: string | undefined
}
export interface VerifySignInForm {
  code: number
}

export type OnboardingStep = 'description' | 'location' | 'caretaker' | 'units' | 'shoot' | 'type'

export interface LocationOption {
  readonly label: string
  readonly value: string
  readonly postalCode: string
  readonly id: string
}

export type PropertyType = 'apartment' | 'condo' | 'bungalow'

export interface FormValues {
  name: string
  propertyType: PropertyType | undefined
  minPrice: string
  maxPrice: string
  town: LocationOption
  postalCode: string
  units: Array<Unit>
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
  countryCode: string
}

interface Amenity {
  id: number
  value: string
  label: string
  category: string
}

interface Bedroom {
  bedroomNumber: number
  enSuite: string
  master: string
}

interface Unit {
  amenities: Amenity[]
  name: string
  type: string
  baths: number
  price: number
  bedrooms: Bedroom[]
}

export interface UnitsForm {
  units: Unit[]
}

export interface AmenitiesForm { amenities: Amenity[] }

export interface PropertyTypeForm {
  propertyType: PropertyType | undefined
}

export interface ContactPersonForm {
  contactPerson: string | undefined
  shootDate: string | undefined
}
