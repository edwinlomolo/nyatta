import { Dispatch, SetStateAction, createContext } from 'react'

import { FormState, Control, UseFormRegister, UseFormGetValues, UseFormReset, UseFormSetValue, UseFormHandleSubmit } from 'react-hook-form'

import { OnboardingStep, FormValues, DescriptionForm, PriceForm, LocationForm, CaretakerForm, UnitsForm } from '../types'

interface OnboardingContext {
  step: OnboardingStep
  setStep: Dispatch<SetStateAction<OnboardingStep>>
  handleSubmit: UseFormHandleSubmit<FormValues>
  register: UseFormRegister<FormValues>
  formState: FormState<FormValues>
  towns: any[]
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
})
