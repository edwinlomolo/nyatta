import { createContext, Dispatch, SetStateAction } from 'react'

import { Status, SignInForm } from '@types'

interface SignInContext {
  status: Status | undefined
  setStatus: Dispatch<SetStateAction<Status>>
  signInForm: SignInForm | undefined
  setSignInForm: Dispatch<SetStateAction<SignInForm>>
}

export const SignInContext = createContext<SignInContext>({
  status: undefined,
  setStatus: () => {},
  signInForm: undefined,
  setSignInForm: () => {},
})
