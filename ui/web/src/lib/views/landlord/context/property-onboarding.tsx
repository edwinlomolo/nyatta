import { Dispatch, SetStateAction, createContext } from 'react'

import { FormState, Control, UseFormRegister, UseFormGetValues, UseFormReset, UseFormSetValue, UseFormHandleSubmit, FieldValues } from 'react-hook-form'

import { OnboardingStep } from '../types'

interface OnboardingContext {
  step: OnboardingStep
  setStep: Dispatch<SetStateAction<OnboardingStep>>
  handleSubmit: UseFormHandleSubmit<FieldValues>
  register: UseFormRegister<FieldValues>
  formState: FormState<FieldValues>
  towns: any[]
  control: Control<FieldValues>
  setValue: UseFormSetValue<FieldValues>
  getValues: UseFormGetValues<FieldValues>
  reset: UseFormReset<FieldValues>
}

export const OnboardingContext = createContext<OnboardingContext>({
  step: 'description',
  setStep: () => {},
  handleSubmit: {} as UseFormHandleSubmit<FieldValues>,
  formState: {} as FormState<FieldValues>,
  register: {} as UseFormRegister<FieldValues>,
  towns: [],
  control: {} as Control<FieldValues>,
  setValue: {} as UseFormSetValue<FieldValues>,
  getValues: {} as UseFormGetValues<FieldValues>,
  reset: {} as UseFormReset<FieldValues>,
})
