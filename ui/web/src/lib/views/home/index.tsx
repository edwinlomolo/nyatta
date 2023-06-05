import Head from 'next/head'

import Footer from './components/Footer'
import HomeHeader from './components/HomeHeader'

const Home = () => (
    <>
      <Head>
        <title>Nyatta - Find homes or apartments for rent.</title>
      </Head>
      <HomeHeader />
      <Footer />
    </>
  )

export default Home
