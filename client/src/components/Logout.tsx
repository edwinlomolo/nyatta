import { GoogleLogout } from 'react-google-login'

const { REACT_APP_GOOGLE_CLIENT_ID } = process.env

function Login() {
  return (
    <GoogleLogout
      clientId={REACT_APP_GOOGLE_CLIENT_ID!}
      onLogoutSuccess={() => alert('Logged out')}
      buttonText="Logout"
    />
  )
}

export default Login
