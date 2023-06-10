import { type Dispatch, type SetStateAction, createContext } from 'react'

import { type FormState, type Control, type UseFormRegister, type UseFormGetValues, type UseFormReset, type UseFormSetValue, type UseFormHandleSubmit } from 'react-hook-form'
import { type GroupBase } from 'react-select'

import { type OnboardingStep, type FormValues, type DescriptionForm, type PriceForm, type LocationForm, type CaretakerForm, type UnitsForm, type LocationOption, type AmenitiesForm } from '../types'

interface OnboardingContext {
  step: OnboardingStep
  setStep: Dispatch<SetStateAction<OnboardingStep>>
  handleSubmit: UseFormHandleSubmit<FormValues>
  register: UseFormRegister<FormValues>
  formState: FormState<FormValues>
  towns: GroupBase<LocationOption>[]
  control: Control<FormValues>
  setValue: UseFormSetValue<FormValues>
  getValues: UseFormGetValues<FormValues>
  reset: UseFormReset<FormValues>
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
}

export const OnboardingContext = createContext<OnboardingContext>({
  step: 'description',
  setStep: () => {},
  handleSubmit: {} as UseFormHandleSubmit<FormValues>,
  formState: {} as FormState<FormValues>,
  register: {} as UseFormRegister<FormValues>,
  towns: [],
  control: {} as Control<FormValues>,
  setValue: {} as UseFormSetValue<FormValues>,
  getValues: {} as UseFormGetValues<FormValues>,
  reset: {} as UseFormReset<FormValues>,
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
  setAmenitiesForm: () => {}
})
