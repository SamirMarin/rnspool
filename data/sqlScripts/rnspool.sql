
DROP TABLE IF EXISTS rideOffered, rideNeeded;
DROP TABLE IF EXISTS leg;
DROP TABLE IF EXISTS rideHasRoute;
DROP TABLE IF EXISTS route;
DROP TABLE IF EXISTS ride;
DROP TABLE IF EXISTS  vehicle;
DROP TABLE IF EXISTS rider, driver, session_table;
DROP TABLE IF EXISTS user_table;
DROP TABLE IF EXISTS address;
DROP TABLE IF EXISTS location;



CREATE TABLE location (
  id SERIAL PRIMARY KEY,
  city VARCHAR(100),
  province VARCHAR(100),
  country VARCHAR(100),
  UNIQUE (city, province, country)
);

CREATE TABLE address (
  id SERIAL PRIMARY KEY ,
  aptNum INTEGER,
  houseNum INTEGER,
  street VARCHAR(100),
  postalCode VARCHAR(10),
  locationId INTEGER NOT NULL REFERENCES location (id),
  UNIQUE (aptNum, houseNum, street, postalCode)
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


CREATE TABLE ride (
  id SERIAL PRIMARY KEY,
  startDescrip VARCHAR(100), -- address or latlon all in one as per the google api
  endDescrip VARCHAR(100),
  createdAt TIMESTAMP,
  locId INTEGER NOT NULL REFERENCES location (id)
);

CREATE TABLE rideOffered(
  rideId INTEGER,
  availableSeats INTEGER,
  timeLeaving TIMESTAMP,
  vehicleId INTEGER NOT NULL REFERENCES vehicle (id),
  PRIMARY KEY (rideId),
  FOREIGN KEY (rideId) REFERENCES ride(id)

);

CREATE TABLE rideNeeded(
  rideId INTEGER,
  neededSeats INTEGER,
  timePickUp TIMESTAMP,
  riderId INTEGER REFERENCES rider (userId),
  PRIMARY KEY (rideId),
  FOREIGN KEY (rideId) REFERENCES ride(id)
);

CREATE TABLE route (
  id SERIAL PRIMARY KEY,
  startDescrip VARCHAR(100),
  endDescrip VARCHAR(100),
  description VARCHAR(100),
  UNIQUE (startDescrip, endDescrip, description)
);

CREATE TABLE rideHasRoute (
  rideId INTEGER,
  routeId INTEGER,
  PRIMARY KEY (rideId, routeId),
  FOREIGN KEY (rideId) REFERENCES ride(id),
  FOREIGN KEY (routeId) REFERENCES route(id)
);

CREATE TABLE leg (
  startPointLat DOUBLE PRECISION,
  startPointLon DOUBLE PRECISION,
  endPointLat DOUBLE PRECISION,
  endPointLon DOUBLE PRECISION,
  htmlInstr VARCHAR(150),
  duration BIGINT,
  distance INTEGER,
  routeId INTEGER NOT NULL REFERENCES route(id),
  PRIMARY KEY (startPointLat, startPointLon, endPointLat, endPointLon)
);

