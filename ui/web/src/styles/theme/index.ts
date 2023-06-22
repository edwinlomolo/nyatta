import { extendTheme, type ThemeConfig } from '@chakra-ui/react'

import { Button } from './button'

const colorConfig: ThemeConfig = {
  initialColorMode: 'light',
  useSystemColorMode: false
}

export const theme = extendTheme({
  config: colorConfig,
  components: {
    Button,
    FormLabel: {
      baseStyle: {
        fontWeight: 'bold'
      }
    }
  },
  fonts: {
    heading: 'var(--font-ibm-plex-mono)',
    body: 'var(--font-ibm-plex-mono)',
  },
  styles: {
    global: () => ({})
  }
})

export { chakraStylesConfig } from './chakra-select'
