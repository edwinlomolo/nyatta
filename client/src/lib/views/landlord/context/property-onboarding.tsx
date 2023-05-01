import { Dispatch, SetStateAction, createContext } from 'react'

import { FormState, UseFormRegister, UseFormHandleSubmit, FieldValues } from 'react-hook-form'

import { OnboardingSteps } from '../types'

interface OnboardingContext {
  step: OnboardingSteps
  setStep: Dispatch<SetStateAction<OnboardingSteps>>
  handleSubmit: UseFormHandleSubmit<FieldValues>
  register: UseFormRegister<FieldValues>
  formState: FormState<FieldValues>
}

export const OnboardingContext = createContext<OnboardingContext>({
  step: 'description',
  setStep: () => {},
  handleSubmit: {} as UseFormHandleSubmit<FieldValues>,
  formState: {} as FormState<FieldValues>,
  register: {} as UseFormRegister<FieldValues>,
})
