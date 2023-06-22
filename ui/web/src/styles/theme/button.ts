import { defineStyleConfig } from '@chakra-ui/react'

export const Button = defineStyleConfig({
  variants: {
    solid: {
      bg: 'green.700',
    },
  },
  defaultProps: {
    colorScheme: 'green',
    variant: 'solid',
  },
})
