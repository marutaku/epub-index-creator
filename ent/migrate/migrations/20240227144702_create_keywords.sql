-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_pages" table
CREATE TABLE `new_pages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `title` text NOT NULL, `book_pages` integer NULL, CONSTRAINT `pages_books_pages` FOREIGN KEY (`book_pages`) REFERENCES `books` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "pages" to new temporary table "new_pages"
INSERT INTO `new_pages` (`id`, `title`) SELECT `id`, `title` FROM `pages`;
-- Drop "pages" table after copying rows
DROP TABLE `pages`;
-- Rename temporary table "new_pages" to "pages"
ALTER TABLE `new_pages` RENAME TO `pages`;
-- Create "keywords" table
CREATE TABLE `keywords` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `keyword` text NOT NULL, `page_keywords` integer NULL, CONSTRAINT `keywords_pages_keywords` FOREIGN KEY (`page_keywords`) REFERENCES `pages` (`id`) ON DELETE SET NULL);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
