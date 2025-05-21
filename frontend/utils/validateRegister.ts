export const validateRegister = (username: string, email: string, password: string) => {
    const errors: string[] = []
    if (!username || username.length < 3) errors.push('ชื่อผู้ใช้ต้องมากกว่า 3 ตัวอักษร')
    if (!email || !email.includes('@')) errors.push('กรุณากรอกอีเมลให้ถูกต้อง')
    if (!password || password.length < 6) errors.push('รหัสผ่านต้องมากกว่า 6 ตัวอักษร')
    return errors
  }
  