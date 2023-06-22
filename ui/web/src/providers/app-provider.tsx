'use client'

import { ReactNode } from 'react'

import { Center, Spinner } from '@chakra-ui/react'
import { useSession } from 'next-auth/react'

interface Props {
  children: ReactNode
}

const AppProvider = ({ children }: Props) => {
  const { status } = useSession()

  return status === 'loading' ? <Center><Spinner thickness="10px" color="green.700" size="xl" /></Center> : <>{children}</>
}

export default AppProvider
