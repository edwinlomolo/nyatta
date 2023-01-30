import { useEffect, useState } from 'react'

import { Box } from '@chakra-ui/react'

function ClientSideRender({ children, ...delegated }) {
  const [hasMounted, setHasMounted] = useState(false)

  useEffect(() => {
    setHasMounted(true)
  }, [])

  if (!hasMounted) {
    return null
  }

  return (
    <Box {...delegated}>{children}</Box>
  )
}

export default ClientSideRender
