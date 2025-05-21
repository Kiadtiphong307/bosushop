<template>
  <div class="container mx-auto px-4 py-8" v-if="product">
    <div class="flex flex-col md:flex-row gap-6">
      <img :src="product.image_url" alt="" class="w-full md:w-1/2 h-64 object-cover rounded" />
      <div class="flex-1">
        <h1 class="text-2xl font-bold mb-2">{{ product.name }}</h1>
        <p class="text-gray-600 mb-4">{{ product.description }}</p>
        <p class="text-blue-600 font-bold text-xl">฿{{ product.price }}</p>
        <p class="text-sm text-gray-400 mt-2">หมวดหมู่: {{ product.category?.name }}</p>
        <button class="mt-6 bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700">
          สั่งซื้อสินค้า
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useProductDetail } from '~/composables/useProductDetail'

const route = useRoute()
const { product, fetchProduct } = useProductDetail()

onMounted(() => {
  const id = route.params.id
  if (id) fetchProduct(id.toString())
})
</script>
