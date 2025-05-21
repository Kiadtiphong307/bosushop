<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white p-8 rounded shadow-md w-full max-w-md">
      <h2 class="text-3xl font-bold mb-6 text-center text-green-600">สมัครสมาชิก</h2>
      <form @submit.prevent="handleRegister" class="space-y-4">
        <input v-model="username" placeholder="Username" class="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-green-400" required />
        <input v-model="email" type="email" placeholder="Email" class="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-green-400" required />
        <input v-model="password" type="password" placeholder="Password" class="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-green-400" required />
        <button type="submit" class="w-full bg-green-600 text-white py-2 px-4 rounded hover:bg-green-700">สมัคร</button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'
const { register } = useAuth()
const username = ref('')
const email = ref('')
const password = ref('')

const handleRegister = async () => {
  await register({ username: username.value, email: email.value, password: password.value })
}
</script>