'use client'

import { ReactNode } from 'react'

import { Center, Spinner } from '@chakra-ui/react'
import { useSession } from 'next-auth/react'

interface Props {
  children: ReactNode
}

const AppProvider = ({ children }: Props) => {
  const { status } = useSession()

  // wait for auth
  if (status === 'loading') {
    return (
      <Center>
        <Spinner thickness="8px" color="green.700" size="xl" />
      </Center>
    )
  }
  return <>{children}</>
}

export default AppProvider
