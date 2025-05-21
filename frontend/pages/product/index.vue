<template>
  <div class="container mx-auto py-6">
    <h1 class="text-2xl font-bold mb-4">รายการสินค้า</h1>
    <div v-if="loading">กำลังโหลด...</div>
    <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div
        v-for="item in products"
        :key="item.id"
        class="border rounded-xl p-4 shadow hover:shadow-lg transition"
      >
        <img :src="item.image_url" alt="" class="w-full h-48 object-cover mb-2 rounded" />
        <h2 class="font-semibold text-lg">{{ item.name }}</h2>
        <p class="text-green-600 font-bold">฿{{ item.price }}</p>
        <NuxtLink
          :to="`/product/${item.slug}`"
          class="mt-2 inline-block text-blue-500 hover:underline"
        >ดูรายละเอียด</NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useProduct } from '@/composables/useProduct'

const { products, loading, fetchProducts } = useProduct()

onMounted(fetchProducts)
</script>
