import { useContext } from 'react'

import { SignInContext } from '../contexts/sign-in'

export const useSignIn = () => useContext(SignInContext)
