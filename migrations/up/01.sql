-- members
CREATE TABLE members (
    id TEXT PRIMARY KEY, -- we will use Google sub or uuid text; if prefer uuid: change to uuid DEFAULT uuid_generate_v4()
    email TEXT NOT NULL UNIQUE,
    user_name TEXT,
    nickname TEXT,
    image_url TEXT,
);

-- trips
CREATE TABLE trips (
    id SERIAL PRIMARY KEY,
    owner_id TEXT NOT NULL REFERENCES members(id),
    trip_name TEXT NOT NULL,
    description TEXT,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    main_location TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    last_accessed_at TIMESTAMPTZ NOT NULL,
);

CREATE INDEX idx_trips_owner_id ON trips(owner_id);
CREATE INDEX idx_trips_last_accessed ON trips(last_accessed_at DESC);

-- trip_members (roles)
CREATE TABLE trip_members (
    trip_id INT NOT NULL REFERENCES trips(id),
    user_id TEXT NOT NULL REFERENCES members(id),
    role trip_role NOT NULL, -- VIEWER, OWNER, EDITOR
    created_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (trip_id, user_id)
);

CREATE INDEX idx_trip_members_user_id ON trip_members(user_id);
CREATE INDEX idx_trip_members_trip_id ON trip_members(trip_id);

-- activities
CREATE TABLE activities (
    id TEXT PRIMARY KEY DEFAULT uuid_generate_v4(),
    trip_id INT NOT NULL REFERENCES trips(id),
    activity_date DATE NOT NULL,
    start_time TIME, -- local time only
    end_time TIME,
    notes TEXT,
    name TEXT,
    location jsonb,   
    category TEXT,
    rank int,
    created_at TIMESTAMPTZ NOT NULL    
);

CREATE INDEX idx_activities_trip_date ON activities(trip_id, activity_date);
CREATE INDEX idx_activities_trip_date_time ON activities(trip_id, activity_date, start_time);

-- groups
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    trip_id INT NOT NULL REFERENCES trips(id),
    name TEXT NOT NULL,
    description TEXT,
    members TEXT[],
    created_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_groups_trip_id ON groups(trip_id);

-- expenses
CREATE TYPE expense_type AS ENUM ('SHARED','EXTRA');

CREATE TABLE expenses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    trip_id INT NOT NULL REFERENCES trips(id),
    title TEXT NOT NULL,
    amount NUMERIC(12,2) NOT NULL CHECK (amount >= 0),
    expense_type TEXT NOT NULL,
    split_type TEXT NOT NULL,
    paid_by TEXT NOT NULL REFERENCES members(id),
    created_by TEXT NOT NULL REFERENCES members(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_expenses_trip_id ON expenses(trip_id);
CREATE INDEX idx_expenses_paid_by ON expenses(paid_by);

CREATE TABLE expense_participants (
    id SERIAL PRIMARY KEY,
    expense_id UUID NOT NULL REFERENCES expenses(id),
    member_id TEXT NOT NULL REFERENCES members(id),
    amount NUMERIC(12,2) NOT NULL CHECK (share >= 0)
);

CREATE INDEX idx_expense_participants_expense_id ON expense_participants(expense_id);
CREATE INDEX idx_expense_participants_member_id ON expense_participants(member_id);


-- invitations
CREATE TABLE trip_invitations (
    id SERIAL PRIMARY KEY,
    trip_id INT NOT NULL REFERENCES trips(id),
    email TEXT NOT NULL,
    token TEXT NOT NULL UNIQUE, -- secure random token
    role trip_role NOT NULL DEFAULT 'VIEWER',
    status TEXT NOT NULL DEFAULT 'PENDING', -- PENDING / ACCEPTED / EXPIRED / REVOKED
    invited_by TEXT NOT NULL REFERENCES members(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    expires_at TIMESTAMPTZ
);

CREATE INDEX idx_trip_invitations_trip_id ON trip_invitations(trip_id);
CREATE INDEX idx_trip_invitations_email ON trip_invitations(email);

