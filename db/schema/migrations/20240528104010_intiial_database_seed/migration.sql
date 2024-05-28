-- CreateTable
CREATE TABLE "User" (
    "telegram_id" TEXT NOT NULL PRIMARY KEY,
    "display_username" TEXT NOT NULL,
    "has_verified" BOOLEAN NOT NULL DEFAULT false,
    "create_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" DATETIME NOT NULL
);

-- CreateTable
CREATE TABLE "Options" (
    "key" TEXT NOT NULL PRIMARY KEY,
    "value" TEXT NOT NULL
);
