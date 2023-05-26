export type OnboardingSteps = 'description' | 'location' | 'amenities' | 'pricing' | 'caretaker' | 'units'

export type LocationOption = {
  readonly label: string
  readonly value: string
  readonly postalCode: string
  readonly id: string
}

