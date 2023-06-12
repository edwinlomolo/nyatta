import { extendTheme, type ThemeConfig } from '@chakra-ui/react'

const colorConfig: ThemeConfig = {
  initialColorMode: 'light',
  useSystemColorMode: false
}

export const theme = extendTheme({
  config: colorConfig,
  components: {
    FormLabel: {
      baseStyle: {
        fontWeight: 'bold'
      }
    }
  },
  fonts: {
    heading: 'var(--font-mabry)',
    body: 'var(--font-mabry)',
  },
  styles: {
    global: () => ({})
  }
})

export { chakraStylesConfig } from './chakra-select'
