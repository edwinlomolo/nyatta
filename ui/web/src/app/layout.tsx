import type { ReactNode } from 'react'

import Layout from '@layout'
import Providers from 'providers/root'

interface Props {
  children: ReactNode
}

export const metadata: Metadata = {
  title: "Nyatta",
  description: "Find rental homes in your local town",
}

const AppLayout = ({ children }: Props) => (
  <html lang="en">
    <body>
      <Providers>
        <Layout>
          {children}
        </Layout>
      </Providers>
    </body>
  </html>
)

export default AppLayout
