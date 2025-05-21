<template>
    <div class="flex items-center justify-center min-h-screen bg-gray-100">
      <div class="bg-white p-8 rounded shadow-md w-full max-w-md">
        <h2 class="text-3xl font-bold mb-6 text-center text-blue-600">เข้าสู่ระบบ</h2>
        <form @submit.prevent="handleLogin" class="space-y-4">
          <input v-model="email" type="email" placeholder="Email" class="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-400" required />
          <input v-model="password" type="password" placeholder="Password" class="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-400" required />
          <button type="submit" class="w-full bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700">เข้าสู่ระบบ</button>
        </form>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { useAuth } from '@/composables/useAuth'
  import { validateLogin } from '@/utils/validateLogin'
  const { login } = useAuth()
  const email = ref('')
  const password = ref('')
  
  const handleLogin = async () => {
    const errors = validateLogin(email.value, password.value)
    if (errors.length > 0) {
      alert(errors.join('\n'))
      return
    }
    await login({ email: email.value, password: password.value })
  }
  </script>