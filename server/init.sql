-- Create Role table
CREATE TABLE IF NOT EXISTS Role (
    rolID SERIAL PRIMARY KEY,
    roltype VARCHAR(255) NOT NULL
);

-- Insert sample roles
INSERT INTO Role (roltype) VALUES ('admin'), ('guest'), ('member');

-- Create User table
CREATE TABLE IF NOT EXISTS "User" (
    userID SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Create Group table
CREATE TABLE IF NOT EXISTS "Group" (
    groupID SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    createdDate DATE NOT NULL
);

-- Create Room table
CREATE TABLE IF NOT EXISTS Room (
    roomID SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    private BOOLEAN NOT NULL,
    groupID INT REFERENCES "Group"(groupID)
);




-- Create UserGroupRol table
CREATE TABLE IF NOT EXISTS UserGroupRol (
    userGroupRolID SERIAL PRIMARY KEY,
    userID INT REFERENCES "User"(userID),
    rolID INT REFERENCES Role(rolID),
    groupID INT REFERENCES "Group"(groupID),
    PRIMARY KEY (userID, rolID, groupID)
);

-- Create RoomUserRol table
CREATE TABLE IF NOT EXISTS RoomUserRol (
    roomUserRolID SERIAL PRIMARY KEY,
    rollID INT REFERENCES Role(rolID),
    userID INT REFERENCES "User"(userID),
    roomID INT REFERENCES Room(roomID),
    deuda DECIMAL,
    pagado BOOLEAN,
    PRIMARY KEY (userID, roomID)
);

