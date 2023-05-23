import { extendTheme, StyleFunctionProps, ThemeConfig } from '@chakra-ui/react'

const colorConfig: ThemeConfig = {
  initialColorMode: 'light',
  useSystemColorMode: false,
}

export const theme = extendTheme({
  config: colorConfig,
  //fonts,
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

