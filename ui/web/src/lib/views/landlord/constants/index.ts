import { OnboardingStep } from '../types'

export const FormSteps: OnboardingStep[] = ['description', 'location', 'pricing', 'caretaker', 'units']

export const FormStepTitle: Record<OnboardingStep, string> = {
  description: 'Describe your property?',
  location: 'Property location',
  amenities: 'What does your property offer?',
  pricing: 'How do you price your units?',
  caretaker: 'Who is the caretaker?',
  units: 'Add property units',
}
