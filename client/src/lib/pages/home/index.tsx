import Head from 'next/head'

import HomeHeader from './components/HomeHeader'
import Footer from './components/Footer'

function Home() {

  return (
    <>
      <Head>
        <title>Nyatta - Find homes or apartments for rent.</title>
      </Head>
      <HomeHeader />
      <Footer />
    </>
  )
}

export default Home
