import { type PropsWithChildren } from 'react'

import { useLazyQuery } from '@apollo/client'

import { useForm } from 'react-hook-form'

import { getListings as GET_LISTINGS } from '@gql'

import { SearchListingContext } from '../context/search-listings'

import { type SearchListingForm } from '@form'

export const SearchListingProvider = ({ children }: PropsWithChildren) => {
  const [getListings, { loading, data }] = useLazyQuery(GET_LISTINGS)
  const { handleSubmit, control, register, formState, getValues, setValue } = useForm<SearchListingForm>({
    defaultValues: { town: '', minPrice: 0, maxPrice: 0 }
  })

  return (
    <SearchListingContext.Provider
      value={{
        control,
        getListings,
        handleSubmit,
        register,
        setValue,
        formState,
        listingsLoading: loading,
        listingsData: data,
        formValues: getValues()
      }}
    >
      {children}
    </SearchListingContext.Provider>
  )
}
