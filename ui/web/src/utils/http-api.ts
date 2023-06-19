import axios from 'axios'

interface HttpRequest {
  method: 'get' | 'post'
  url: string
  data: Record<string, any>
}

class Http {
  async post (url: string, data: any) {
    try {
      return await this.request({ method: 'post', url, data })
    } catch (error) {
      console.error(error)
    }
  }

  async get (url: string, data: any) {
    try {
      return await this.request({ method: 'get', url, data })
    } catch (error) {
      console.error(error)
    }
  }

  async request ({ method, url, data }: HttpRequest) {
    try {
      const response = await axios({
        method,
        url,
        data,
        headers: {
          'content-type': 'application/json'
        }
      })
      return response.data
    } catch (error) {
      console.error(error)
    }
  }
}

export default Http
