import { createContext } from 'react'

import { LazyQueryResult, LazyQueryHookOptions, OperationVariables } from '@apollo/client'

import { UseFormRegister, FieldValues, FormState, UseFormHandleSubmit } from 'react-hook-form'

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
  formValues: any
}

export const SearchListingContext = createContext<ISearchListings>({
  register: {} as UseFormRegister<SearchListingForm>,
  handleSubmit: {} as UseFormHandleSubmit<FieldValues>,
  formState: {} as FormState<FieldValues>,
  getListings: {} as GetListingsFunction,
  listingsLoading: false,
  listingsData: {},
  formValues: {},
})
