import { useEffect } from 'react'

import { gapi } from 'gapi-script'

import { Login, Logout } from './components'

const { REACT_APP_GOOGLE_CLIENT_ID } = process.env

function App() {

  useEffect(() => {
    const initializeClient = () => {
      gapi.client.init({
        clientId: REACT_APP_GOOGLE_CLIENT_ID,
        scope: '',
      })
    }

    gapi.load('client:auth2', initializeClient)
  })

  return (
    <>
      <h1>Welcome to Nyatta!</h1>
      <Login />
      <Logout />
    </>
  );
}

export default App;
