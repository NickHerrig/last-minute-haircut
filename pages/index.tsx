import type { NextPage } from 'next'
import Head from 'next/head'

const Home: NextPage = () => {
  return (
    <div className="">
      <Head>
        <title>Last Minute DSM</title>
        <meta name="A web app for last minute haircuts from paramount barber"/>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="">
        <h1 className="text-3xl font-bold underline">Last Minute DSM</h1>
      </main>

      <footer className="">
      </footer>
    </div>
  )
}

export default Home
