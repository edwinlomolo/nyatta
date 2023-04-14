export const apiUrl = `${process.env.NEXT_PUBLIC_BASE_API}/handshake`

export const isPrototypeEnv = (): boolean => {
  if (process.env.NEXT_PUBLIC_ENV === 'development') {
    return true
  } else {
    return false
  }
}
