export const validateLogin = (email: string, password: string) => {
    const errors: string[] = []
    if (!email || !email.includes('@')) errors.push('กรุณากรอกอีเมลให้ถูกต้อง')
    if (!password || password.length < 6) errors.push('รหัสผ่านต้องมากกว่า 6 ตัวอักษร')
    return errors
  }
  