import { extendTheme, StyleFunctionProps, ThemeConfig } from '@chakra-ui/react'

// import { Button } from './components'
import { fonts } from './fonts'

const colorConfig: ThemeConfig = {
  initialColorMode: 'light',
  useSystemColorMode: false,
}

export const theme = extendTheme({
  config: colorConfig,
  fonts,
  /*
  components: {
    Button,
  },
  */
  styles: {
    global: (props: StyleFunctionProps) => ({
      body: {
        bg: "#ffff",
      },
    })
  },
})

