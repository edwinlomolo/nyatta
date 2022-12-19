import { useAuth0 } from '@auth0/auth0-react'

import { Login, Logout, Profile } from './components'



function App() {
  const { isAuthenticated, isLoading } = useAuth0()

  if (isLoading) {
    return <p>Loading ...</p>
  }

  return (
    <div>
      <h1>Welcome to Nyatta!</h1>
      {isAuthenticated && <Profile />}
      {!isAuthenticated && <Login />}
      {isAuthenticated && <Logout />}
    </div>
  );
}

export default App;
