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
  styles: {
    global: () => ({})
  }
})
