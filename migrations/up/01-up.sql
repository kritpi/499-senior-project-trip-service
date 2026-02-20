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

ALTER TABLE trip
ADD COLUMN image_url TEXT NOT NULL DEFAULT '';

-- activity
CREATE TABLE activity (
    id TEXT PRIMARY KEY NOT NULL, -- UUID
    trip_id INT NOT NULL REFERENCES trip(id),
    activity_date DATE NOT NULL,
    start_time TIME,
    end_time TIME,
    note TEXT,
    description TEXT,
    activity_location JSONB,
    category TEXT,
    rank INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

-- group
CREATE TABLE grouping (
    id TEXT PRIMARY KEY NOT NULL, -- UUID
    trip_id INT NOT NULL REFERENCES trip(id),
    name TEXT NOT NULL,
    description TEXT,    
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

-- group member
CREATE TABLE grouping_members (
    group_id TEXT NOT NULL REFERENCES grouping(id),
    trip_id INT NOT NULL REFERENCES trip(id),    
    member_id TEXT NOT NULL REFERENCES member(id),
    PRIMARY KEY(group_id, trip_id, member_id)
);

-- expense
CREATE TABLE expense (
    id TEXT PRIMARY KEY NOT NULL, -- uuid
    trip_id INT NOT NULL REFERENCES trip(id), 
    title TEXT NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    created_by TEXT NOT NULL REFERENCES member(id),
    split_type TEXT NOT NULL, -- CUSTOM, ALL_EQUAL, SELECTED_EQUAL
    image_url TEXT NOT NULL DEFAULT ''
);

-- expense member
CREATE TABLE expense_members (
    expense_id TEXT NOT NULL REFERENCES expense(id),
    member_id TEXT NOT NULL REFERENCES member(id),
    amount NUMERIC(12,2) NOT NULL,
    PRIMARY KEY (expense_id, member_id)
);