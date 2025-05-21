package services

import (
	"backend/database"
	"backend/models"
	"errors"
	"time"
)

func CreateOrder(userID uint, productID uint, couponCode *string) (*models.Order, error) {
	var product models.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		return nil, errors.New("ไม่พบสินค้า")
	}

	if product.Stock <= 0 {
		return nil, errors.New("สินค้าหมดแล้ว")
	}

	total := product.Price

	var coupon *models.Coupon
	if couponCode != nil && *couponCode != "" {
		var found models.Coupon
		if err := database.DB.Where("code = ?", *couponCode).First(&found).Error; err != nil {
			return nil, errors.New("ไม่พบคูปอง")
		}

		if found.ExpireAt.Before(time.Now()) {
			return nil, errors.New("คูปองหมดอายุแล้ว")
		}

		if found.UsedCount >= found.MaxUsage {
			return nil, errors.New("คูปองถูกใช้เต็มจำนวนแล้ว")
		}

		total = total - (total * float64(found.DiscountPercent) / 100)
		coupon = &found
	}

	// บันทึกคำสั่งซื้อ
	order := &models.Order{
		UserID:     userID,
		ProductID:  productID,
		TotalPrice: total,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	if coupon != nil {
		order.CouponCode = &coupon.Code

		// อัปเดตจำนวนใช้คูปอง
		coupon.UsedCount++
		if err := database.DB.Save(coupon).Error; err != nil {
			return nil, errors.New("ใช้คูปองไม่ได้")
		}
	}

	// ลด stock
	product.Stock--
	if err := database.DB.Save(&product).Error; err != nil {
		return nil, errors.New("ไม่สามารถอัปเดต stock ได้")
	}

	if err := database.DB.Create(order).Error; err != nil {
		return nil, errors.New("สร้างคำสั่งซื้อไม่สำเร็จ")
	}

	return order, nil
}

func GetOrdersByUser(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Preload("Product").Where("user_id = ?", userID).Order("created_at desc").Find(&orders).Error
	return orders, err
}

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Preload("Product").Preload("User").Order("created_at desc").Find(&orders).Error
	return orders, err
}
