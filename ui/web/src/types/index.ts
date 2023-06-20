export type Status = 'pending' | 'approved' | 'sign-in'
export interface SignInForm {
  phone: number | undefined
  countryCode: string | undefined
}
export interface VerifySignInForm {
  code: number
}
