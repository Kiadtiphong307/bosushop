<template>
  <nav class="bg-gray-100 p-4 flex justify-between items-center shadow">
    <NuxtLink to="/" class="text-xl font-bold text-blue-600">BoShop</NuxtLink>
    <div class="space-x-4">
      <template v-if="!user">
        <NuxtLink to="/login" class="text-blue-600 hover:underline">เข้าสู่ระบบ</NuxtLink>
        <NuxtLink to="/register" class="text-green-600 hover:underline">สมัครสมาชิก</NuxtLink>
      </template>
      <template v-else>
        <NuxtLink to="/profile" class="text-purple-600 hover:underline">
          {{ user.username }}
        </NuxtLink>
        <button @click="logout" class="text-red-600 hover:underline">ออกจากระบบ</button>
      </template>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'
const { user, logout } = useAuth()
</script>
