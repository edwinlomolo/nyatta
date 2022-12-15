import { useState } from 'react'

import { GoogleLogin, GoogleLoginResponse, GoogleLoginResponseOffline } from 'react-google-login'

const { REACT_APP_GOOGLE_CLIENT_ID } = process.env

function Login() {
  const [userInfo, setUserInfo] = useState<GoogleLoginResponse | GoogleLoginResponseOffline>()
  return (
    <>
      <div>
        {JSON.stringify(userInfo)}
      </div>
      <GoogleLogin
        clientId={REACT_APP_GOOGLE_CLIENT_ID!}
        buttonText="Login"
        onSuccess={res => setUserInfo(res)}
        onFailure={res => console.log(res)}
        isSignedIn={true}
      />
    </>
  )
}

export default Login
