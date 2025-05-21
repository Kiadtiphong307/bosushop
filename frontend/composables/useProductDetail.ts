import { ref } from 'vue'
import axios from 'axios'

// ดึงข้อมูลสินค้า 
export function useProductDetail() {
  const product = ref<any>(null)

  const fetchProduct = async (id: string) => {
    try {
      const res = await axios.get(`/api/products/${id}`)
      product.value = res.data
    } catch (err) {
      console.error('โหลดข้อมูลสินค้าไม่สำเร็จ:', err)
    }
  }

  return { product, fetchProduct }
}
