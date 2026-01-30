package seeders

import (
	"backend_jamu/models"
	"backend_jamu/utils"
	"context"
	"log"

	"github.com/uptrace/bun"
)

func SeedAdmin(db *bun.DB) {
	ctx := context.Background()

	// Check if roles/users table exists (or just try to count)
	count, _ := db.NewSelect().Model((*models.User)(nil)).Where("role = ?", "admin").Count(ctx)
	if count > 0 {
		log.Println("Admin already exists, skipping seed")
		return
	}

	hashedPassword, _ := utils.HashPassword("admin123")
	admin := &models.User{
		Name:     "Admin Nusantara",
		Email:    "admin@gmail.com",
		Password: hashedPassword,
		Role:     "admin",
	}

	_, err := db.NewInsert().Model(admin).Exec(ctx)
	if err != nil {
		log.Printf("Failed to seed admin: %v\n", err)
	} else {
		log.Println("Admin seeded successfully: admin@example.com / password123")
	}
}
