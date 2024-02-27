-- Create "pages" table
CREATE TABLE `pages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `book_id` text NOT NULL, `title` text NOT NULL, CONSTRAINT `pages_books_pages` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE SET NULL);
