CREATE TABLE conscious_app.dev_schema.users (
    first_name STRING,
    last_name STRING,
    email STRING
);

CREATE TABLE conscious_app.dev_schema.messages(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    message STRING NOT NULL,
    is_seen BOOL NOT NULL,
    deleted_from_sender BOOL NOT NULL,
    deleted_from_receiver BOOL NOT NULL,
    user_id UUID NOT NULL,
    conversation_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE `conversations`(
    `id` INT NOT NULL,
    `user_one` INT NOT NULL,
    `user_two` INT NOT NULL,
    `status` TINYINT(1) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
);

CREATE TABLE `relationships`(
    `id` BIGINT NOT NULL,
    `follower_id` INT NULL,
    `followed_id` INT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
);

CREATE TABLE `tweets`(
    `id` BIGINT NOT NULL,
    `content` VARCHAR(255) NULL,
    `user_id` INT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
);

CREATE TABLE `users`(
    `id` BIGINT NOT NULL,
    `name` VARCHAR(255) NULL,
    `email` VARCHAR(255) NULL,
    `password_digest` VARCHAR(255) NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `remember_token` VARCHAR(255) NULL,
    `slug` VARCHAR(255) NULL
);