import { useCookie } from '#app'
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

  // ✅ เข้าสู่ระบบ
  const login = async (data: { email: string; password: string }) => {
    try {
      const res = await $fetch<AuthResponse>('/api/auth/login', {
        method: 'POST',
        body: data
      })

      token.value = res.token
      user.value = res.user

      await router.push('/')
    } catch (error) {
      throw error
    }
  }

  // ✅ สมัครสมาชิก
  const register = async (data: { username: string; email: string; password: string }) => {
    try {
      const res = await $fetch<AuthResponse>('/api/auth/register', {
        method: 'POST',
        body: data
      })

      token.value = res.token
      user.value = res.user

      await router.push('/')
    } catch (error) {
      throw error
    }
  }

  // ✅ ออกจากระบบ
  const logout = () => {
    token.value = null
    user.value = null
    router.push('/')
  }

  return { user, token, login, register, logout }
}
