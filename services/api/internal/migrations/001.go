package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/mmtaee/ocserv-dashboard/common/pkg/logger"
	"gorm.io/gorm"
)

var Migration001 = &gormigrate.Migration{
	ID: "001_create_initial_tables",
	Migrate: func(tx *gorm.DB) error {

		// =========================
		// USERS
		// =========================
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS users (
				id BIGSERIAL PRIMARY KEY,
				uid CHAR(26) NOT NULL UNIQUE,
				owner VARCHAR(16) DEFAULT '',
				"group" VARCHAR(16) DEFAULT 'defaults',
				username VARCHAR(16) NOT NULL UNIQUE,
				password VARCHAR(16) NOT NULL,
				is_locked BOOLEAN DEFAULT FALSE,
				created_at TIMESTAMP,
				updated_at TIMESTAMP,
				expire_at DATE,
				deactivated_at DATE,
				traffic_type VARCHAR(32) NOT NULL DEFAULT '1',
				traffic_size BIGINT NOT NULL,
				rx BIGINT NOT NULL DEFAULT 0,
				tx BIGINT NOT NULL DEFAULT 0,
				description TEXT,
				config TEXT
			);
		`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_users_uid ON users(uid);`).Error; err != nil {
			return err
		}

		// =========================
		// USER TOKENS
		// =========================
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS user_tokens (
				id BIGSERIAL PRIMARY KEY,
				user_id BIGINT NOT NULL,
				uid TEXT NOT NULL UNIQUE,
				token TEXT,
				created_at TIMESTAMP,
				expire_at TIMESTAMP,
				CONSTRAINT fk_user_tokens_user
					FOREIGN KEY(user_id)
					REFERENCES users(id)
					ON DELETE CASCADE
			);
		`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_user_tokens_user_id ON user_tokens(user_id);`).Error; err != nil {
			return err
		}

		// =========================
		// SYSTEM
		// =========================
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS system (
				id BIGSERIAL PRIMARY KEY,
				google_captcha_secret TEXT,
				google_captcha_site_key TEXT
			);
		`).Error; err != nil {
			return err
		}

		// =========================
		// OC SERV GROUPS
		// =========================
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS ocserv_groups (
				id BIGSERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL UNIQUE,
				owner VARCHAR(32) DEFAULT '',
				config TEXT
			);
		`).Error; err != nil {
			return err
		}

		// =========================
		// TRAFFIC STATS
		// =========================
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS ocserv_user_traffic_statistics (
				id BIGSERIAL PRIMARY KEY,
				oc_user_id BIGINT NOT NULL,
				created_at TIMESTAMP,
				rx BIGINT DEFAULT 0,
				tx BIGINT DEFAULT 0,
				CONSTRAINT fk_traffic_user
					FOREIGN KEY(oc_user_id)
					REFERENCES users(id)
					ON DELETE CASCADE
			);
		`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
			CREATE INDEX IF NOT EXISTS idx_traffic_statistics_oc_user_id
			ON ocserv_user_traffic_statistics(oc_user_id);
		`).Error; err != nil {
			return err
		}

		logger.Info("migration 001 (Postgres) complete successfully")
		return nil
	},
}
