import { ColorModeScript } from '@chakra-ui/react'
import type { DocumentContext } from 'next/document'
import Document, { Html, Head, Main, NextScript } from 'next/document'

import { theme } from '@styles'

class MyDocument extends Document {
  static async getInitialProps (ctx: DocumentContext) {
    return await Document.getInitialProps(ctx)
  }

  render () {
    return (
      <Html lang="en">
        <Head>
        </Head>
        <body>
          <ColorModeScript initialColorMode={theme.config.initialColorMode} />
          <Main />
          <NextScript />
        </body>
      </Html>
    )
  }
}

export default MyDocument
