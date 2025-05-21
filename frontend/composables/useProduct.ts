// composables/useProduct.ts
import axios from 'axios'

export interface Product {
  id: number
  name: string
  slug: string
  image_url: string
  price: number
}

export const useProduct = () => {
  const products = ref<Product[]>([])
  const product = ref<Product | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  // ดึงข้อมูลสินค้าทั้งหมด
  const fetchProducts = async () => {
    loading.value = true
    try {
      const res = await axios.get('/api/products')
      products.value = res.data
    } catch (err) {
      error.value = 'ไม่สามารถโหลดสินค้าได้'
    } finally {
      loading.value = false
    }
  }

  // ดึงข้อมูลสินค้าตาม slug
  const fetchProductBySlug = async (slug: string) => {
    loading.value = true
    try {
      const res = await axios.get(`/api/products/${slug}`)
      product.value = res.data
    } catch (err) {
      error.value = 'ไม่พบสินค้านี้'
    } finally {
      loading.value = false
    }
  }

  return {
    products,
    product,
    loading,
    error,
    fetchProducts,
    fetchProductBySlug
  }
}
