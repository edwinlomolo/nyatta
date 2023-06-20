import { type PropsWithChildren, useState } from 'react'

import { defaultSignInForm } from '@constants'
import { Status, SignInForm } from '@types'
import { SignInContext } from 'contexts/sign-in'

const SignInProvider = ({ children }: PropsWithChildren) => {
  const [status, setStatus] = useState<Status>('sign-in')
  const [signInForm, setSignInForm] = useState<SignInForm>(defaultSignInForm)

  return (
    <SignInContext.Provider
      value={{
        status,
        setStatus,
        signInForm,
        setSignInForm,
      }}
    >
      {children}
    </SignInContext.Provider>
  )
}

export default SignInProvider
