CREATE TABLE IF NOT EXISTS dishes (
    id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL DEFAULT 'images/dishes/default.png',
    video_link VARCHAR(255) NOT NULL,
    calories REAL NOT NULL,
    variation INTEGER NOT NULL,
    day_time INTEGER NOT NULL
);
