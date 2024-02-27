-- Create "books" table
CREATE TABLE `books` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `isbn` text NOT NULL, `title` text NOT NULL, `language` text NOT NULL, `author` text NOT NULL, `publisher` text NOT NULL);
