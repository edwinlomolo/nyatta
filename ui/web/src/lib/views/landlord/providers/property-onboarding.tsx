import { useState, type PropsWithChildren } from 'react'

import { useQuery } from '@apollo/client'

import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { array, object, string } from 'yup'

import { getTowns as GET_TOWNS } from '@gql'
import { OnboardingContext } from '../context/property-onboarding'
import { type OnboardingStep, type FormValues, type DescriptionForm, type LocationForm, type CaretakerForm, type PriceForm, type UnitsForm } from '../types'
import { defaultDescriptionForm, defaultLocationForm, defaultPriceForm, defaultUnitsForm, defaultCaretakerForm } from '../constants'

const validationSchema = object().shape({
  name: string().required('Property name required'),
  propertyType: string().required('What is your property type?'),
  minPrice: string().required('Minimum price required'),
  maxPrice: string().required('Maximum price required'),
  town: object().shape({
    label: string().required(),
    value: string().required(),
    postalCode: string().required(),
    id: string().required()
  }).required('Town is required'),
  postalCode: string().required(),
  units: array()
    .of(
      object().shape({
        text: string().required('Unit name required')
      })
    )
    .required('If you got here, then your flat has unit(s) to be registered')
})

export const OnboardingProvider = ({ children }: PropsWithChildren) => {
  const [descriptionForm, setDescriptionForm] = useState<DescriptionForm>(defaultDescriptionForm)
  const [locationForm, setLocationForm] = useState<LocationForm>(defaultLocationForm)
  const [priceForm, setPriceForm] = useState<PriceForm>(defaultPriceForm)
  const [caretakerForm, setCaretakerForm] = useState<CaretakerForm>(defaultCaretakerForm)
  const [unitsForm, setUnitsForm] = useState<UnitsForm>(defaultUnitsForm)
  // For default towns select input
  const { data } = useQuery(GET_TOWNS)
  const locations = data?.getTowns.map((item: any) => ({
    id: item.id,
    value: item.town.toLowerCase(),
    label: item.town,
    postalCode: item.postalCode
  }))

  const [step, setStep] = useState<OnboardingStep>('caretaker')
  const { control, getValues, reset, setValue, handleSubmit, formState, register } = useForm<FormValues>({
    mode: 'onChange',
    resolver: yupResolver(validationSchema)
  })

  return (
    <OnboardingContext.Provider
      value={{
        priceForm,
        setPriceForm,
        descriptionForm,
        setDescriptionForm,
        caretakerForm,
        setCaretakerForm,
        unitsForm,
        setUnitsForm,
        locationForm,
        setLocationForm,
        step,
        handleSubmit,
        register,
        formState,
        setStep,
        towns: locations,
        control,
        setValue,
        reset,
        getValues
      }}
    >
      {children}
    </OnboardingContext.Provider>
  )
}
