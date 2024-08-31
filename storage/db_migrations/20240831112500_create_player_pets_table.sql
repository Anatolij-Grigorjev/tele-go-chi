-- +goose Up
CREATE TABLE IF NOT EXISTS player_pets (
    id INT PRIMARY KEY AUTO_INCREMENT,
    player_id varchar(64) NOT NULL,
    pet_uuid varchar(40) NOT NULL,
    pet_emoji varchar(10) NOT NULL,
    pet_alive BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX player_pets_index (player_id),
    INDEX player_livestock_index (player_id, pet_alive),
    INDEX player_pet_index (player_id, pet_uuid)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;