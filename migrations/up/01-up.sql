-- members
CREATE TABLE member (
    id TEXT PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    name TEXT,    
    image_url TEXT
);


-- trip
CREATE TABLE trip (
    id SERIAL PRIMARY KEY,
    owner_id TEXT NOT NULL REFERENCES member(id),
    trip_name TEXT NOT NULL,
    description TEXT,
    start_date DATE,
    end_date DATE,
    main_location TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

ALTER DATABASE postgres SET TIMEZONE='Asia/Bangkok';