import { useCookie, useFetch } from '#app'
import { useRouter } from 'vue-router'

export const useAuth = () => {
  const user = useState<AuthUser | null>('user', () => null)
  const token = useCookie<string | null>('token', { maxAge: 60 * 60 * 24 })
  const router = useRouter()

  type AuthUser = {
    id: number
    username: string
    email: string
    role: string
  }

  type AuthResponse = {
    token: string
    user: AuthUser
  }

  // เข้าสู่ระบบ
  const login = async (data: { email: string; password: string }) => {
    const { data: res, error } = await useFetch<AuthResponse>('/api/auth/login', {
      method: 'POST',
      body: data
    })
    if (error.value) throw error.value

    token.value = res.value?.token || null
    user.value = res.value?.user || null
    await router.push('/')
  }

  // สมัครสมาชิก
  const register = async (data: { username: string; email: string; password: string }) => {
    const { data: res, error } = await useFetch<AuthResponse>('/api/auth/register', {
      method: 'POST',
      body: data
    })
    if (error.value) throw error.value

    token.value = res.value?.token || null
    user.value = res.value?.user || null
    await router.push('/')
  }

  // ออกจากระบบ
  const logout = () => {
    token.value = null
    user.value = null
    router.push('/')
  }

  return { user, token, login, register, logout }
}
