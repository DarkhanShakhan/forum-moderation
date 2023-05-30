PRAGMA foreign_keys = ON;


CREATE TABLE IF NOT EXISTS user_roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE
);

INSERT OR IGNORE INTO user_roles(name) VALUES
    ("Admin"), ("Moderator"), ("User") ON CONFLICT DO NOTHING; 

UPDATE `sqlite_sequence`
    SET `seq` = (SELECT MAX(`id`) FROM 'user_roles')
    WHERE `name` = 'user_roles';


CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role_id INTEGER NOT NULL,
    FOREIGN KEY(role_id) REFERENCES user_roles(id)
);

INSERT OR IGNORE INTO users (username, email, password, role_id) VALUES("admin", "admin@mail.com", "admin", 1);

UPDATE `sqlite_sequence`
    SET `seq` = (SELECT MAX(`id`) FROM 'users')
    WHERE `name` = 'users';

CREATE TABLE IF NOT EXISTS tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    title TEXT NOT NULL UNIQUE
);

INSERT OR IGNORE INTO tags (title) VALUES
    ('ok'),('irrelevant'), ('obscene'), ('illegal'), ('insulting');

UPDATE `sqlite_sequence`
    SET `seq` = (SELECT MAX(`id`) FROM 'tags')
    WHERE `name` = 'tags';

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author_id INTEGER ,
    title TEXT NOT NULL,
    content TEXT,
    visible INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL,
    deleted_at TEXT,
    deleted_by INTEGER,
    delete_message TEXT,
    delete_category INTEGER,
    FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS posts_reactions (
    post_id INTEGER,
    user_id INTEGER,
    reaction INTEGER,
    created_at TEXT,
    UNIQUE(post_id, user_id),
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts_categories (
    post_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    UNIQUE(post_id, category_id),
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author_id INTEGER ,
    title TEXT NOT NULL,
    content TEXT,
    visible INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL,
    deleted_at TEXT,
    deleted_by INTEGER,
    delete_message TEXT,
    moderator_viewed INTEGER NOT NULL DEFAULT 0,
    moderator_viewed_by INTEGER,
    FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS posts_tags (
    post_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    UNIQUE(post_id, tag_id),
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY(tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS posts_reports (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    reported_by INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    created_at TEXT NOT NULL,
    message TEXT,
    admin_viewed INTEGER NOT NULL DEFAULT 0,
    admin_viewed_by INTEGER
    deleted_at TEXT
);

CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER,
    author_id INTEGER,
    content TEXT,
    visible INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL,
    deleted_at TEXT,
    delete_message TEXT,
    moderator_viewed INTEGER NOT NULL DEFAULT 0,
    moderator_viewed_by INTEGER,
    FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comment_reactions (
    comment_id INTEGER,
    user_id INTEGER,
    reaction INTEGER,
    UNIQUE(comment_id, user_id),
    FOREIGN KEY(comment_id) REFERENCES comments(id) ON DELETE CASCADE,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments_reports (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    reported_by INTEGER NOT NULL,
    comment_id INTEGER NOT NULL,
    created_at TEXT NOT NULL,
    message TEXT,
    admin_viewed INTEGER NOT NULL DEFAULT 0,
    admin_viewed_by INTEGER
    deleted_at TEXT
);

CREATE TABLE IF NOT EXISTS post_reactions (
    post_id INTEGER,
    user_id INTEGER,
    reaction INTEGER,
    UNIQUE(post_id, user_id),
    FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);