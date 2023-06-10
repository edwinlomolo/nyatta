import { type Dispatch, type SetStateAction, createContext } from 'react'

import { type GroupBase } from 'react-select'

import { type OnboardingStep, type DescriptionForm, type PriceForm, type LocationForm, type CaretakerForm, type UnitsForm, type LocationOption, type AmenitiesForm } from '../types'

interface OnboardingContext {
  step: OnboardingStep
  setStep: Dispatch<SetStateAction<OnboardingStep>>
  towns: GroupBase<LocationOption>[]
  descriptionForm: DescriptionForm
  setDescriptionForm: Dispatch<SetStateAction<DescriptionForm>>
  locationForm: LocationForm
  setLocationForm: Dispatch<SetStateAction<LocationForm>>
  priceForm: PriceForm
  setPriceForm: Dispatch<SetStateAction<PriceForm>>
  unitsForm: UnitsForm
  setUnitsForm: Dispatch<SetStateAction<UnitsForm>>
  caretakerForm: CaretakerForm
  setCaretakerForm: Dispatch<SetStateAction<CaretakerForm>>
  unitsCount: number
  setUnitsCount: Dispatch<SetStateAction<number>>
  amenitiesForm: AmenitiesForm
  setAmenitiesForm: Dispatch<SetStateAction<AmenitiesForm>>
  caretakerVerified: boolean
  setCaretakerVerified: Dispatch<SetStateAction<boolean>>
}

export const OnboardingContext = createContext<OnboardingContext>({
  step: 'description',
  setStep: () => {},
  towns: [],
  descriptionForm: {} as DescriptionForm,
  setDescriptionForm: () => {},
  locationForm: {} as LocationForm,
  setLocationForm: () => {},
  priceForm: {} as PriceForm,
  setPriceForm: () => {},
  unitsForm: {} as UnitsForm,
  setUnitsForm: () => {},
  caretakerForm: {} as CaretakerForm,
  setCaretakerForm: () => {},
  unitsCount: 0,
  setUnitsCount: () => {},
  amenitiesForm: {} as AmenitiesForm,
  setAmenitiesForm: () => {},
  caretakerVerified: false,
  setCaretakerVerified: () => false
})
