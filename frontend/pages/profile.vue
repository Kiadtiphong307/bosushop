<template>
    <div class="flex items-center justify-center min-h-screen bg-gray-100">
      <div class="bg-white p-8 rounded shadow-md w-full max-w-md">
        <h2 class="text-3xl font-bold mb-6 text-center text-purple-600">ข้อมูลโปรไฟล์</h2>
        <div class="space-y-2">
          <p><strong>ชื่อผู้ใช้:</strong> {{ user?.username }}</p>
          <p><strong>Email:</strong> {{ user?.email }}</p>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { useAuth } from '@/composables/useAuth'
  const { user } = useAuth()
  </script>