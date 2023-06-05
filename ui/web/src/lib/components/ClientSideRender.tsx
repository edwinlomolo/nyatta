import { useEffect, useRef } from 'react'

import { Box } from '@chakra-ui/react'

const ClientSideRender = ({ children, ...delegated }: { children: React.ReactNode }) => {
  const hasMounted = useRef(false)

  useEffect(() => {
    hasMounted.current = true
  }, [])

  if (!hasMounted.current) {
    return null
  }

  return (
    <Box {...delegated}>{children}</Box>
  )
}

export default ClientSideRender
