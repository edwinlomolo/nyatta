import { createContext } from 'react'

import { LazyQueryResult, LazyQueryHookOptions, OperationVariables } from '@apollo/client'

import { UseFormRegister, FieldValues, FormState, Control, UseFormSetValue, UseFormHandleSubmit } from 'react-hook-form'

import { SearchListingForm } from '@form'

type GetListingsFunction = (
  options?: LazyQueryHookOptions<any, OperationVariables>
) => Promise<LazyQueryResult<any, OperationVariables>>

interface ISearchListings {
  register: UseFormRegister<SearchListingForm>
  handleSubmit: UseFormHandleSubmit<FieldValues>
  formState: FormState<FieldValues>
  getListings: GetListingsFunction
  listingsLoading: boolean
  listingsData: any
  control: Control<SearchListingForm>
  formValues: any
  setValue: UseFormSetValue<SearchListingForm>
}

export const SearchListingContext = createContext<ISearchListings>({
  register: {} as UseFormRegister<SearchListingForm>,
  handleSubmit: {} as UseFormHandleSubmit<FieldValues>,
  formState: {} as FormState<FieldValues>,
  getListings: {} as GetListingsFunction,
  setValue: {} as UseFormSetValue<SearchListingForm>,
  listingsLoading: false,
  listingsData: {},
  control: {} as Control<SearchListingForm>,
  formValues: {},
})
