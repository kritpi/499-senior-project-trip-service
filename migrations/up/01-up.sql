ALTER DATABASE postgres SET TIMEZONE='Asia/Bangkok';

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

-- trip_members (roles)
CREATE TABLE trip_members (
    trip_id INT NOT NULL REFERENCES trip(id),
    member_id TEXT NOT NULL REFERENCES member(id),
    member_role TEXT NOT NULL, -- VIEWER, OWNER, EDITOR
    created_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (trip_id, member_id)
);
