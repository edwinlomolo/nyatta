import type { ReactNode } from 'react'

import { Metadata } from 'next'
import { Space_Grotesk } from 'next/font/google'

import Layout from '@layout'
import AppProvider from 'providers/app-provider'
import Providers from 'providers/root'

interface Props {
  children: ReactNode
}

export const metadata: Metadata = {
  title: 'Nyatta - Find local rental homes',
}

const ibm = Space_Grotesk({
  weight: '600',
  subsets: ['latin'],
  display: 'swap',
})

const AppLayout = ({ children }: Props) => (
  <html lang="en" className={ibm.className}>
    <body>
      <Providers>
        <AppProvider>
          <Layout>
            {children}
          </Layout>
        </AppProvider>
      </Providers>
    </body>
  </html>
)

export default AppLayout
