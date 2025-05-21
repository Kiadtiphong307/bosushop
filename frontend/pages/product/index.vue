<template>
  <div class="container mx-auto px-4 py-8">
    <h1 class="text-2xl font-bold mb-6">สินค้า</h1>

    <!-- Search bar -->
    <div class="mb-4 flex gap-4">
      <input
        v-model="search"
        @keyup.enter="fetchProducts"
        type="text"
        placeholder="ค้นหาสินค้า..."
        class="border px-4 py-2 rounded w-full"
      />
      <button @click="fetchProducts" class="bg-blue-600 text-white px-4 py-2 rounded">
        ค้นหา
      </button>
    </div>

    <!-- Grid of products -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="product in products" :key="product.ID" class="border rounded shadow p-4">
        <img :src="product.ImageURL" alt="" class="w-full h-48 object-cover rounded mb-2" />
        <h2 class="font-semibold text-lg">{{ product.Name  }}</h2>
        <p class="text-gray-600">{{ product.Description }}</p>
        <p class="text-blue-600 font-bold mt-2">฿{{ product.Price }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const products = ref([])
const search = ref('')

const fetchProducts = async () => {
  try {
    const res = await axios.get('/api/products', {
      params: search.value ? { search: search.value } : {}
    })
    console.log('✅ ได้ข้อมูลสินค้า:', res.data) 
    products.value = res.data
  } catch (err) {
    console.error('โหลดสินค้าไม่สำเร็จ:', err)
  }
}


onMounted(fetchProducts)
</script>
