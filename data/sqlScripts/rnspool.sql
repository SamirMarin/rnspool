
DROP TABLE IF EXISTS rideOffered, rideNeeded;
DROP TABLE IF EXISTS ride;
DROP TABLE IF EXISTS  vehicle, route;
DROP TABLE IF EXISTS rider, driver, session_table;
DROP TABLE IF EXISTS user_table;
DROP TABLE IF EXISTS address;
DROP TABLE IF EXISTS location;



CREATE TABLE location (
  id SERIAL PRIMARY KEY,
  city VARCHAR(100),
  province VARCHAR(100),
  country VARCHAR(100)
);

CREATE TABLE address (
  id SERIAL PRIMARY KEY ,
  aptNum INTEGER,
  houseNum INTEGER,
  street VARCHAR(100),
  postalCode VARCHAR(10),
  locationId INTEGER NOT NULL REFERENCES location (id)
);

CREATE TABLE user_table (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR(64) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE, -- max 254 may need to update
  password VARCHAR(128) NOT NULL,
  firstName VARCHAR(255), -- arbitrary limit??
  lastName VARCHAR(255),
  aboutMe VARCHAR(300),
  createdAt TIMESTAMP NOT NULL,
  addressId INTEGER NOT NULL REFERENCES address (id)
);

CREATE TABLE session_table (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR(64) NOT NULL UNIQUE,
  email VARCHAR(255),
  firstName VARCHAR(255),
  lastName VARCHAR(255) ,
  aboutMe VARCHAR(300),
  userId INTEGER REFERENCES user_table (id),
  createdAt TIMESTAMP NOT NULL

);

CREATE TABLE rider (
  userId INTEGER PRIMARY KEY REFERENCES user_table (id),
  riderRating REAL CHECK (riderRating >= 0 AND riderRating <= 5)
);

CREATE TABLE driver (
  userId INTEGER PRIMARY KEY REFERENCES user_table (id),
  driverRating REAl CHECK (driverRating >= 0 AND driverRating <= 5)
);

CREATE TABLE vehicle (
  id SERIAL PRIMARY KEY,
  licence VARCHAR(20),
  make VARCHAR(50),
  model VARCHAR(50),
  year INTEGER,
  numPassengers INTEGER,
  type VARCHAR(50),
  driverId INTEGER REFERENCES driver (userId)
);

CREATE TABLE route (
  id SERIAL PRIMARY KEY,
  startPoint POINT,
  endPoint POINT,
  pickUpDesc VARCHAR(75),
  dropOffDesc VARCHAR(75)
);

CREATE TABLE ride (
  id SERIAL PRIMARY KEY,
  date DATE,
  description VARCHAR(150),
  locId INTEGER NOT NULL REFERENCES location (id),
  routeId INTEGER NOT NULL REFERENCES route (id)
);

CREATE TABLE rideOffered(
  rideId INTEGER PRIMARY KEY REFERENCES ride (id),
  availableSeats INTEGER,
  vehicleId INTEGER NOT NULL REFERENCES vehicle (id)
);

CREATE TABLE rideNeeded(
  rideId INTEGER PRIMARY KEY REFERENCES ride (id),
  neededSeats INTEGER,
  riderId INTEGER REFERENCES rider (userId)
);
