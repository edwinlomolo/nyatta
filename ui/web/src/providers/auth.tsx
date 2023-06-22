import { ReactNode } from 'react'

import { SessionProvider, useSession } from 'next-auth/react'

interface Props {
  children: ReactNode
}

export const Auth = ({ children }: Props) => {
  const { status } = useSession()

  console.log(status)
  if (status === 'loading') {
    <>Loading...</>
  }

  return children
}
