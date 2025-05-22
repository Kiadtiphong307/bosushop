type AuthUser = {
  id: number
  username: string
  email: string
  role: string
}

type ProfileResponse = {
  user: AuthUser
}

export default defineNuxtRouteMiddleware(async () => {
  const token = useCookie('token')
  const user = useState<AuthUser | null>('user')

  if (!token.value) {
    return navigateTo('/login')
  }

  // ดึงข้อมูลผู้ใช้ถ้ายังไม่มี
  if (!user.value) {
    try {
      const { data, error } = await useFetch<ProfileResponse>('/api/auth/profile', {
        headers: { Authorization: `Bearer ${token.value}` }
      })

      if (error.value || !data.value?.user) {
        throw new Error('Failed to load profile')
      }

      user.value = data.value.user
    } catch {
      token.value = null
      return navigateTo('/login')
    }
  }
})
