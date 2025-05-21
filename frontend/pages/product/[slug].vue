<template>
  <div class="container mx-auto py-6">
    <div v-if="loading">กำลังโหลด...</div>
    <div v-else-if="product" class="max-w-xl mx-auto">
      <img :src="product.image_url" class="rounded mb-4" />
      <h1 class="text-2xl font-bold">{{ product.name }}</h1>
      <p class="text-gray-700 my-2">{{ product.description }}</p>
      <p class="text-green-600 font-bold">ราคา: ฿{{ product.price }}</p>
    </div>
    <div v-else>ไม่พบสินค้า</div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { onMounted } from 'vue'
import { useProduct } from '@/composables/useProduct'

const { product, loading, fetchProductBySlug } = useProduct()
const route = useRoute()

onMounted(() => {
  fetchProductBySlug(route.params.slug as string)
})
</script>
