
DROP TABLE IF EXISTS `Users`;
DROP TABLE IF EXISTS `Instructors`;
DROP TABLE IF EXISTS `Courses`;
DROP TABLE IF EXISTS `Users`;
DROP TABLE IF EXISTS `Instructor`;
DROP TABLE IF EXISTS `Category`;
DROP TABLE IF EXISTS `Course`;
DROP TABLE IF EXISTS `Tag`;
DROP TABLE IF EXISTS `Sub_Category`;
DROP TABLE IF EXISTS `Section`;
DROP TABLE IF EXISTS `Lectures`;
DROP TABLE IF EXISTS `cart_courses`;
DROP TABLE IF EXISTS `Answer`;
DROP TABLE IF EXISTS `Cart`;
DROP TABLE IF EXISTS `Coupon`;
DROP TABLE IF EXISTS `Coupons_Courses`;
DROP TABLE IF EXISTS `Course_Images`;
DROP TABLE IF EXISTS `Currency`;
DROP TABLE IF EXISTS `Course_Tag_Mappings`;
DROP TABLE IF EXISTS `Exercise`;
DROP TABLE IF EXISTS `Payment_Method`;
DROP TABLE IF EXISTS `Favourite`;
DROP TABLE IF EXISTS `Price`;
DROP TABLE IF EXISTS `enrollments`;
DROP TABLE IF EXISTS `Payment`;
DROP TABLE IF EXISTS `Payment_Course`;
DROP TABLE IF EXISTS `Paypal`;
DROP TABLE IF EXISTS `schema_migrations`;



SET SQL_MODE='ALLOW_INVALID_DATES';
CREATE TABLE `Users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `firstName` varchar(255) NOT NULL,
                         `lastName` varchar(255) NOT NULL,
                         `email` varchar(255) NOT NULL,
                         `password` varchar(255) NOT NULL,
                         `phoneNumber` varchar(255) NOT NULL,
                         `address` varchar(255) NOT NULL,
                         `role` enum('guest','admin') NOT NULL DEFAULT 'guest',
                         `is_instructor` tinyint(1) DEFAULT '0',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` timestamp NULL DEFAULT NULL,
                         `picture` json DEFAULT NULL,
                         `lastLogin` timestamp NULL DEFAULT '0000-00-00 00:00:00',
                         `verified` tinyint(1) DEFAULT '0',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `Instructor` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `user_id` int NOT NULL,
                              `website` varchar(255) DEFAULT NULL,
                              `linkedin` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                              `youtube` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                              `bio` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
                              `num_students` varchar(10) DEFAULT '0',
                              `num_reviews` varchar(10) DEFAULT '0',
                              `rating` float DEFAULT '0',
                              `total_courses` varchar(4) DEFAULT '0',
                              `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `deleted_at` timestamp NULL DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `user_id` (`user_id`) USING BTREE,
                              CONSTRAINT `Instructor_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `Tag` (
                       `id` int NOT NULL AUTO_INCREMENT,
                       `tag_name` varchar(255) NOT NULL,
                       `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                       `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       `deleted_at` timestamp NULL DEFAULT NULL,
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `Category` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Sub_Category` (
                                `id` int NOT NULL AUTO_INCREMENT,
                                `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                `category_id` int NOT NULL,
                                `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                `deleted_at` timestamp NULL DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                KEY `category_id` (`category_id`),
                                CONSTRAINT `Sub_Category_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `Category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `Course` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                          `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
                          `level` enum('All Level','Beginner','Intermediate','Expert') NOT NULL DEFAULT 'All Level',
                          `language` enum('vi','en') NOT NULL DEFAULT 'en',
                          `price_id` int NOT NULL,
                          `is_publish` tinyint(1) DEFAULT '0',
                          `published_time` timestamp NULL DEFAULT NULL,
                          `instructor_id` int NOT NULL,
                          `thumbnail` json DEFAULT NULL,
                          `SubCategory_id` int DEFAULT NULL,
                          `is_paid` tinyint(1) DEFAULT '0',
                          `num_reviews` int DEFAULT '0',
                          `num_subscribers` int DEFAULT '0',
                          `rating` float DEFAULT '0',
                          `goals` text,
                          `requirement` varchar(255) DEFAULT NULL,
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `deleted_at` timestamp NULL DEFAULT NULL,
                          `total_Lectures` int DEFAULT '0',
                          `total_Sections` int DEFAULT '0',
                          PRIMARY KEY (`id`),
                          KEY `instructor_id` (`instructor_id`),
                          KEY `SubCategory_id` (`SubCategory_id`),
                          KEY `level_id` (`level`),
                          CONSTRAINT `Course_ibfk_2` FOREIGN KEY (`instructor_id`) REFERENCES `Instructor` (`id`),
                          CONSTRAINT `Course_ibfk_3` FOREIGN KEY (`SubCategory_id`) REFERENCES `Sub_Category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE `Section` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `course_id` int NOT NULL,
                           `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           `deleted_at` timestamp NULL DEFAULT NULL,
                           PRIMARY KEY (`id`),
                           KEY `course_id` (`course_id`),
                           CONSTRAINT `Section_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `Course` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `Lectures` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `section_id` int NOT NULL,
                            `title` varchar(70) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `is_free` tinyint(1) NOT NULL DEFAULT '0',
                            `sort_order` int DEFAULT NULL,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `section_id` (`section_id`),
                            CONSTRAINT `Lectures_ibfk_1` FOREIGN KEY (`section_id`) REFERENCES `Section` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `Answer` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `answer_text` varchar(255) NOT NULL,
                          `question_id` int NOT NULL,
                          `is_correct` bit(1) NOT NULL DEFAULT b'0',
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `deleted_at` timestamp NULL DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Cart` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `user_id` int NOT NULL,
                        `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `deleted_at` timestamp NULL DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        KEY `user_id` (`user_id`) USING BTREE,
                        CONSTRAINT `Cart_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `cart_courses` (
                                `cart_id` int NOT NULL,
                                `course_id` int NOT NULL,
                                `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                `deleted_at` timestamp NULL DEFAULT NULL,
                                PRIMARY KEY (`cart_id`,`course_id`),
                                KEY `course_id` (`course_id`),
                                CONSTRAINT `cart_courses_ibfk_2` FOREIGN KEY (`course_id`) REFERENCES `Course` (`id`),
                                CONSTRAINT `cart_courses_ibfk_3` FOREIGN KEY (`cart_id`) REFERENCES `Cart` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE `Coupon` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                          `percentage` float NOT NULL,
                          `start_date` timestamp NOT NULL,
                          `expiry_date` timestamp NOT NULL,
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `deleted_at` timestamp NULL DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Coupons_Courses` (
                                   `coupons_id` int NOT NULL,
                                   `courses_id` int NOT NULL,
                                   PRIMARY KEY (`coupons_id`,`courses_id`),
                                   KEY `courses_id` (`courses_id`),
                                   CONSTRAINT `Coupons_Courses_ibfk_1` FOREIGN KEY (`coupons_id`) REFERENCES `Coupon` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                                   CONSTRAINT `Coupons_Courses_ibfk_2` FOREIGN KEY (`courses_id`) REFERENCES `Course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE `Course_Images` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `course_id` int NOT NULL,
                                 `image_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                 `image_order` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                                 `image_width` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                                 `image_height` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 KEY `course_id` (`course_id`),
                                 CONSTRAINT `Course_Images_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `Course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Course_Tag_Mappings` (
                                       `course_id` int NOT NULL,
                                       `tag_id` int NOT NULL,
                                       PRIMARY KEY (`course_id`,`tag_id`),
                                       KEY `tag_id` (`tag_id`),
                                       CONSTRAINT `Course_Tag_Mappings_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `Course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                                       CONSTRAINT `Course_Tag_Mappings_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `Tag` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Currency` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `enrollments` (
                               `user_id` int NOT NULL,
                               `course_id` int NOT NULL,
                               `enrollment_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                               PRIMARY KEY (`user_id`,`course_id`),
                               KEY `course_id` (`course_id`),
                               CONSTRAINT `enrollments_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                               CONSTRAINT `enrollments_ibfk_2` FOREIGN KEY (`course_id`) REFERENCES `Course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Exercise` (
                            `id` int NOT NULL,
                            `exercise_title` varchar(255) NOT NULL,
                            `description` varchar(255) NOT NULL,
                            `sections_id` int NOT NULL,
                            `order_index` int NOT NULL DEFAULT '0',
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `sections_id` (`sections_id`),
                            CONSTRAINT `Exercise_ibfk_1` FOREIGN KEY (`sections_id`) REFERENCES `Section` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Favourite` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `user_id` int NOT NULL,
                             `course_id` int NOT NULL,
                             `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                             `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `deleted_at` timestamp NULL DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             KEY `user_id` (`user_id`),
                             KEY `course_id` (`course_id`),
                             CONSTRAINT `Favourite_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                             CONSTRAINT `Favourite_ibfk_2` FOREIGN KEY (`course_id`) REFERENCES `Course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `Payment_Method` (
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `name` varchar(50) DEFAULT NULL,
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Payment` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `user_id` int NOT NULL,
                           `amount` decimal(12,2) NOT NULL,
                           `note` text,
                           `payment_method` int DEFAULT NULL,
                           `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           `deleted_at` timestamp NULL DEFAULT NULL,
                           PRIMARY KEY (`id`,`user_id`),
                           KEY `user_id` (`user_id`),
                           KEY `payment_method` (`payment_method`),
                           CONSTRAINT `Payment_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`),
                           CONSTRAINT `Payment_ibfk_4` FOREIGN KEY (`payment_method`) REFERENCES `Payment_Method` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Payment_Course` (
                                  `payment_id` int NOT NULL,
                                  `course_id` int NOT NULL,
                                  `price` decimal(12,2) NOT NULL,
                                  `discount` float NOT NULL DEFAULT '0',
                                  `amount` decimal(12,2) NOT NULL,
                                  PRIMARY KEY (`payment_id`,`course_id`),
                                  KEY `course_id` (`course_id`),
                                  CONSTRAINT `Payment_Course_ibfk_2` FOREIGN KEY (`course_id`) REFERENCES `Course` (`id`),
                                  CONSTRAINT `Payment_Course_ibfk_3` FOREIGN KEY (`payment_id`) REFERENCES `Payment` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE `Paypal` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `paypal_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                          `email` varchar(100) NOT NULL,
                          `user_id` int NOT NULL,
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `deleted_at` timestamp NULL DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          KEY `user_id` (`user_id`) USING BTREE,
                          CONSTRAINT `Paypal_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Price` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `value` varchar(15) NOT NULL,
                         `currency` enum('USD','VND') NOT NULL DEFAULT 'USD',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Question`;
CREATE TABLE `Question` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `question_text` varchar(255) NOT NULL,
                            `exercise_id` int NOT NULL,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `exercise_id` (`exercise_id`),
                            CONSTRAINT `Question_ibfk_1` FOREIGN KEY (`exercise_id`) REFERENCES `Exercise` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Resource`;
CREATE TABLE `Resource` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `url` varchar(255) NOT NULL,
                            `duration` time NOT NULL,
                            `lecture_id` int NOT NULL,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `lecture_id` (`lecture_id`),
                            CONSTRAINT `Resource_ibfk_1` FOREIGN KEY (`lecture_id`) REFERENCES `Lectures` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Review`;
CREATE TABLE `Review` (
                          `user_id` int NOT NULL,
                          `course_id` int NOT NULL,
                          `rating` tinyint NOT NULL,
                          `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          PRIMARY KEY (`user_id`,`course_id`),
                          KEY `course_id` (`course_id`),
                          CONSTRAINT `Review_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                          CONSTRAINT `Review_ibfk_2` FOREIGN KEY (`course_id`) REFERENCES `Course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `schema_migrations` (
                                     `version` bigint NOT NULL,
                                     `dirty` tinyint(1) NOT NULL,
                                     PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Transaction`;
CREATE TABLE `Transaction` (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `total_price` decimal(12,2) NOT NULL,
                               `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                               `payment_id` int NOT NULL,
                               `cart_id` int DEFAULT NULL,
                               `data` json DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               KEY `payment_id` (`payment_id`),
                               KEY `cart_id` (`cart_id`),
                               CONSTRAINT `Transaction_ibfk_2` FOREIGN KEY (`payment_id`) REFERENCES `Payment` (`id`),
                               CONSTRAINT `Transaction_ibfk_3` FOREIGN KEY (`cart_id`) REFERENCES `Cart` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `UserAnswer`;
CREATE TABLE `UserAnswer` (
                              `user_id` int NOT NULL,
                              `question_id` int NOT NULL,
                              `answer_id` int NOT NULL,
                              `answer_text` varchar(255) NOT NULL,
                              `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `deleted_at` timestamp NULL DEFAULT NULL,
                              PRIMARY KEY (`user_id`,`question_id`,`answer_id`),
                              KEY `question_id` (`question_id`),
                              KEY `answer_id` (`answer_id`),
                              CONSTRAINT `UserAnswer_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                              CONSTRAINT `UserAnswer_ibfk_2` FOREIGN KEY (`question_id`) REFERENCES `Question` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                              CONSTRAINT `UserAnswer_ibfk_3` FOREIGN KEY (`question_id`) REFERENCES `Question` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                              CONSTRAINT `UserAnswer_ibfk_4` FOREIGN KEY (`answer_id`) REFERENCES `Answer` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;





INSERT INTO `Courses` (`id`, `title`, `description`, `level`, `language`, `price`, `discount`, `duration`, `status`, `rating`, `instructor_id`, `created_at`, `updated_at`, `deleted_at`, `currency`, `thumbnail`)
VALUES
    (1, 'Python For Beginners', 'This python for beginners course is geared to students who want to know how python works and also to those totally new to programming.\n\nThe python language has very simple syntax(way to write it) to learn and it is one of the most powerful languages to learn since it can used for a variety of things.\n\nHere are some:\n\nData analysis\nGame development\nVisualization\nWeb development\nRobotics\nand more....\n\n\nJobs in this field are really lucrative and knowing this language will give you an edge when finding a job and making a lot more money than other developers; python developers are not as many as in other languages since people think is hard. Python is super easy to learn but very powerful since it contains many possibilities.\n\nPython is growing faster and faster everyday and it has surpassed many other languages over the years for a lot of reasons, which you will find out soon enough.\n\n------------------------------------------------------------------------------\n\nWhy take this course?\n\nThis course starts with explaining what programming really is? Have you ever wondered how things actually work in a program?\n\n1. Teaches the foundation of GENERAL programming\n\nEvery student should have some foundations on what programming really is before learning any language, why? Because once you understand the core components of programming it will be a lot easier to learn any language and create better programs.\n\n2. New lectures added all the time\n\nWhen you like what you do, it reflects. This is not a job for me, so I wake up wanting to code and help my students. Basically you pay once, get the course forever and get extra values added all the time, you have to love that :)\n\n3. Fun place to be\n\nI love making my lectures fun and engaging, so no boring lectures here!\n\n4. The support you get in this course in unmatched\n\nhave you ever joined a class and received very little support or none at all? Well that is not going to happen in this class, because I love helping my students.\n\n\n\nWho this course is for:\nStudent totally new to programming\nStudent totally new to python\n',
     'Beginner', 'English', 13.52, NULL, '12:15:00', NULL, NULL, 1, '2023-03-09 19:34:27', '2023-04-29 16:10:27', NULL, 'usd', '{"id": 1, "url": "https://d338hmodjsosv2.cloudfront.net/img/avatar-zhoangp", "width": "250px", "height": "250px"}'),
    (2, 'Photograph', 'This python for beginners course is geared to students who want to know how python works and also to those totally new to programming.\n\nThe python language has very simple syntax(way to write it) to learn and it is one of the most powerful languages to learn since it can used for a variety of things.\n\nHere are some:\n\nData analysis\nGame development\nVisualization\nWeb development\nRobotics\nand more....\n\n\nJobs in this field are really lucrative and knowing this language will give you an edge when finding a job and making a lot more money than other developers; python developers are not as many as in other languages since people think is hard. Python is super easy to learn but very powerful since it contains many possibilities.\n\nPython is growing faster and faster everyday and it has surpassed many other languages over the years for a lot of reasons, which you will find out soon enough.\n\n------------------------------------------------------------------------------\n\nWhy take this course?\n\nThis course starts with explaining what programming really is? Have you ever wondered how things actually work in a program?\n\n1. Teaches the foundation of GENERAL programming\n\nEvery student should have some foundations on what programming really is before learning any language, why? Because once you understand the core components of programming it will be a lot easier to learn any language and create better programs.\n\n2. New lectures added all the time\n\nWhen you like what you do, it reflects. This is not a job for me, so I wake up wanting to code and help my students. Basically you pay once, get the course forever and get extra values added all the time, you have to love that :)\n\n3. Fun place to be\n\nI love making my lectures fun and engaging, so no boring lectures here!\n\n4. The support you get in this course in unmatched\n\nhave you ever joined a class and received very little support or none at all? Well that is not going to happen in this class, because I love helping my students.\n\n\n\nWho this course is for:\nStudent totally new to programming\nStudent totally new to python\n',
     'Beginner', 'English', 13.52, NULL, '12:15:00', NULL, NULL, 1, '2023-03-09 19:34:27', '2023-04-29 16:10:27', NULL, 'usd', '{"id": 1, "url": "https://d338hmodjsosv2.cloudfront.net/img/avatar-zhoangp", "width": "250px", "height": "250px"}'),
    (3, 'Adobe Photoshop', 'This python for beginners course is geared to students who want to know how python works and also to those totally new to programming.\n\nThe python language has very simple syntax(way to write it) to learn and it is one of the most powerful languages to learn since it can used for a variety of things.\n\nHere are some:\n\nData analysis\nGame development\nVisualization\nWeb development\nRobotics\nand more....\n\n\nJobs in this field are really lucrative and knowing this language will give you an edge when finding a job and making a lot more money than other developers; python developers are not as many as in other languages since people think is hard. Python is super easy to learn but very powerful since it contains many possibilities.\n\nPython is growing faster and faster everyday and it has surpassed many other languages over the years for a lot of reasons, which you will find out soon enough.\n\n------------------------------------------------------------------------------\n\nWhy take this course?\n\nThis course starts with explaining what programming really is? Have you ever wondered how things actually work in a program?\n\n1. Teaches the foundation of GENERAL programming\n\nEvery student should have some foundations on what programming really is before learning any language, why? Because once you understand the core components of programming it will be a lot easier to learn any language and create better programs.\n\n2. New lectures added all the time\n\nWhen you like what you do, it reflects. This is not a job for me, so I wake up wanting to code and help my students. Basically you pay once, get the course forever and get extra values added all the time, you have to love that :)\n\n3. Fun place to be\n\nI love making my lectures fun and engaging, so no boring lectures here!\n\n4. The support you get in this course in unmatched\n\nhave you ever joined a class and received very little support or none at all? Well that is not going to happen in this class, because I love helping my students.\n\n\n\nWho this course is for:\nStudent totally new to programming\nStudent totally new to python\n',
     'Beginner', 'English', 13.52, NULL, '12:15:00', NULL, NULL, 1, '2023-03-09 19:34:27', '2023-04-29 16:10:27', NULL, 'usd', '{"id": 1, "url": "https://d338hmodjsosv2.cloudfront.net/img/avatar-zhoangp", "width": "250px", "height": "250px"}'),
    (4, 'Adobe XD', 'This python for beginners course is geared to students who want to know how python works and also to those totally new to programming.\n\nThe python language has very simple syntax(way to write it) to learn and it is one of the most powerful languages to learn since it can used for a variety of things.\n\nHere are some:\n\nData analysis\nGame development\nVisualization\nWeb development\nRobotics\nand more....\n\n\nJobs in this field are really lucrative and knowing this language will give you an edge when finding a job and making a lot more money than other developers; python developers are not as many as in other languages since people think is hard. Python is super easy to learn but very powerful since it contains many possibilities.\n\nPython is growing faster and faster everyday and it has surpassed many other languages over the years for a lot of reasons, which you will find out soon enough.\n\n------------------------------------------------------------------------------\n\nWhy take this course?\n\nThis course starts with explaining what programming really is? Have you ever wondered how things actually work in a program?\n\n1. Teaches the foundation of GENERAL programming\n\nEvery student should have some foundations on what programming really is before learning any language, why? Because once you understand the core components of programming it will be a lot easier to learn any language and create better programs.\n\n2. New lectures added all the time\n\nWhen you like what you do, it reflects. This is not a job for me, so I wake up wanting to code and help my students. Basically you pay once, get the course forever and get extra values added all the time, you have to love that :)\n\n3. Fun place to be\n\nI love making my lectures fun and engaging, so no boring lectures here!\n\n4. The support you get in this course in unmatched\n\nhave you ever joined a class and received very little support or none at all? Well that is not going to happen in this class, because I love helping my students.\n\n\n\nWho this course is for:\nStudent totally new to programming\nStudent totally new to python\n',
     'Beginner', 'English', 13.52, NULL, '12:15:00', NULL, NULL, 1, '2023-03-09 19:34:27', '2023-04-29 16:10:27', NULL, 'usd', '{"id": 1, "url": "https://d338hmodjsosv2.cloudfront.net/img/avatar-zhoangp", "width": "250px", "height": "250px"}'),
    (5, 'Basic JS', 'This python for beginners course is geared to students who want to know how python works and also to those totally new to programming.\n\nThe python language has very simple syntax(way to write it) to learn and it is one of the most powerful languages to learn since it can used for a variety of things.\n\nHere are some:\n\nData analysis\nGame development\nVisualization\nWeb development\nRobotics\nand more....\n\n\nJobs in this field are really lucrative and knowing this language will give you an edge when finding a job and making a lot more money than other developers; python developers are not as many as in other languages since people think is hard. Python is super easy to learn but very powerful since it contains many possibilities.\n\nPython is growing faster and faster everyday and it has surpassed many other languages over the years for a lot of reasons, which you will find out soon enough.\n\n------------------------------------------------------------------------------\n\nWhy take this course?\n\nThis course starts with explaining what programming really is? Have you ever wondered how things actually work in a program?\n\n1. Teaches the foundation of GENERAL programming\n\nEvery student should have some foundations on what programming really is before learning any language, why? Because once you understand the core components of programming it will be a lot easier to learn any language and create better programs.\n\n2. New lectures added all the time\n\nWhen you like what you do, it reflects. This is not a job for me, so I wake up wanting to code and help my students. Basically you pay once, get the course forever and get extra values added all the time, you have to love that :)\n\n3. Fun place to be\n\nI love making my lectures fun and engaging, so no boring lectures here!\n\n4. The support you get in this course in unmatched\n\nhave you ever joined a class and received very little support or none at all? Well that is not going to happen in this class, because I love helping my students.\n\n\n\nWho this course is for:\nStudent totally new to programming\nStudent totally new to python\n',
     'Beginner', 'English', 13.52, NULL, '12:15:00', NULL, NULL, 1, '2023-03-09 19:34:27', '2023-04-29 16:10:27', NULL, 'usd', '{"id": 1, "url": "https://d338hmodjsosv2.cloudfront.net/img/avatar-zhoangp", "width": "250px", "height": "250px"}');

DROP TABLE IF EXISTS `Sections`;
CREATE TABLE `Sections` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `course_id` int NOT NULL,
                            `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `numberOfLectures` int NOT NULL,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `course_id` (`course_id`),
                            CONSTRAINT `Sections_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `Courses` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `Sections` (`id`, `course_id`, `title`, `numberOfLectures`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1, 3, 'Introduce', 3, '2023-04-25 02:44:44', '2023-04-25 02:48:09', NULL),
    (2, 3, 'Main', 3, '2023-04-25 02:44:44', '2023-04-25 02:48:09', NULL);


DROP TABLE IF EXISTS `Lectures`;
CREATE TABLE `Lectures` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `section_id` int NOT NULL,
                            `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `section_id` (`section_id`),
                            CONSTRAINT `Lectures_ibfk_1` FOREIGN KEY (`section_id`) REFERENCES `Sections` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `Lectures` (`id`, `section_id`, `title`, `content`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
                                                                                                                        (1, 1, 'Hello', 'abc', 'ok', '2023-04-25 02:47:16', '2023-04-25 02:47:16', NULL),
                                                                                                                        (2, 2, 'Hello 2', 'xyz', 'ok', '2023-04-25 02:47:16', '2023-04-25 17:41:00', NULL);

DROP TABLE IF EXISTS `LectureResources`;
CREATE TABLE `LectureResources` (
                                    `id` int NOT NULL AUTO_INCREMENT,
                                    `lecture_id` int NOT NULL,
                                    `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                    `duration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    `deleted_at` timestamp NULL DEFAULT NULL,
                                    PRIMARY KEY (`id`),
                                    KEY `lecture_id` (`lecture_id`),
                                    CONSTRAINT `LectureResources_ibfk_1` FOREIGN KEY (`lecture_id`) REFERENCES `Lectures` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `LectureResources` (`id`, `lecture_id`, `url`, `duration`, `created_at`, `updated_at`, `deleted_at`) VALUES
    (1, 1, 'https://courses-storages.s3.ap-northeast-1.amazonaws.com/video/Golang+installation+and+hello+world.mp4', '5:00', '2023-04-25 03:15:36', '2023-04-25 11:02:41', NULL);

INSERT INTO `Paypal` (`id`, `paypal_id`, `email`, `user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
                                                                                                           (2, 'agag', 'ababa', 1, '2023-05-20 02:00:22', '2023-05-19 19:20:40', '2023-05-20 02:20:41'),
                                                                                                           (7, 'agag', '12121', 1, '2023-05-20 02:20:41', '2023-05-20 10:37:42', '2023-05-20 17:37:43'),
                                                                                                           (9, 'agag', '12121', 1, '2023-05-20 02:21:22', '2023-05-20 19:35:04', '2023-05-21 02:35:05'),
                                                                                                           (10, 'NQA4CPD4XL4Q2', 'userb1@gmail.com', 5, '2023-05-20 17:37:43', '2023-05-20 19:34:06', NULL),
                                                                                                           (11, 'B8RCSDPEX4EKW', 'abcxyz123456@gmail.com', 20, '2023-05-21 02:35:05', '2023-05-20 19:35:14', NULL);

INSERT INTO `Price` (`id`, `value`, `currency`, `created_at`, `updated_at`, `deleted_at`) VALUES
                                                                                              (1, '429000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (2, '449000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (3, '479000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (4, '499000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (5, '549000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (6, '579000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (7, '599000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (8, '629000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (9, '649000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (10, '679000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (11, '699000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (12, '729000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (13, '749000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (14, '779000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (15, '799000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (16, '829000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (17, '849000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (18, '879000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (19, '899000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (20, '929000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (21, '949000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (22, '979000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (23, '999000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (24, '1049000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (25, '1099000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (26, '1399000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (27, '1499000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (28, '1699000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (29, '1799000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (30, '1899000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (31, '1999000', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (32, '19.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (33, '24.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (34, '29.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (35, '34.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (36, '39.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (37, '44.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (38, '49.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (39, '54.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (40, '59.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (41, '64.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (42, '69.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (43, '74.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (44, '79.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (45, '84.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (46, '89.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (47, '94.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (48, '99.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (49, '109.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (50, '119.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (51, '124.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (52, '129.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (53, '139.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (54, '149.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (55, '159.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (56, '169.99', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (57, 'free', 'USD', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL),
                                                                                              (58, 'free', 'VND', '2023-05-26 17:54:48', '2023-05-26 17:54:48', NULL);
