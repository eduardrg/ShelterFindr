CREATE TABLE state (
	abbrev char(2) PRIMARY KEY,
	"full" varchar(255) UNIQUE
);

CREATE TABLE address (
	id SERIAL PRIMARY KEY,
	stateAbbrev char(2) REFERENCES state (abbrev) NOT NULL,
	streetAddress1 varchar(255) NOT NULL,
	streetAddress2 varchar(255),
	city varchar(255) NOT NULL,
	zip char(5) NOT NULL,
	zipExtended char(5)
);

CREATE TABLE shelter (
	id SERIAL PRIMARY KEY,
	addressId integer REFERENCES address (id) NOT NULL,
	name varchar(255) NOT NULL,
	"desc" varchar(255),
	phone varchar(16),
	email varchar(255),
	URL varchar(255)
);

CREATE TABLE focusInfo (
	id SERIAL PRIMARY KEY,
	type varchar(128) NOT NULL,
	"desc" varchar(255)
);

CREATE TABLE focus (
	id SERIAL PRIMARY KEY,
	shelterId integer REFERENCES shelter (id) NOT NULL,
	focusInfoId integer REFERENCES focusInfo (id) NOT NULL,
	isStrict boolean
);

CREATE TABLE "group" (
	id SERIAL PRIMARY KEY,
	name varchar(255)
);

CREATE TABLE room (
	id SERIAL PRIMARY KEY,
	shelterId integer REFERENCES shelter (id),
	no varchar(255) NOT NULL,
	capacity smallint NOT NULL,
	occupancy smallint NOT NULL
);

CREATE TABLE client (
	id SERIAL PRIMARY KEY,
	currentRoomId integer REFERENCES room (id),
	desiredShelterId integer REFERENCES shelter (id),
	groupId integer REFERENCES "group" (id),
	firstName varchar(255) NOT NULL,
	lastName varchar(255) NOT NULL,
	yearsOld smallint,
	gender varchar(255),
	bio text,
	waitlistPosition smallint
);

CREATE TABLE waitlist (
	id SERIAL PRIMARY KEY,
	shelterId integer REFERENCES shelter (id) NOT NULL,
	clientId integer REFERENCES client (id) NOT NULL,
	whenAdded timestamp NOT NULL
);

CREATE TABLE item (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	"desc" varchar(255),
	qty numeric NOT NULL,
	units varchar(255) NOT NULL
);

CREATE TABLE itemStock (
	shelterId integer REFERENCES shelter (id) NOT NULL,
	itemId integer REFERENCES item (id) NOT NULL,
	count smallint NOT NULL,
	PRIMARY KEY (shelterId, itemId)
);

CREATE TABLE service (
	id SERIAL PRIMARY KEY,
	type varchar(128) NOT NULL,
	"desc" varchar(255)
);

CREATE TABLE shelterServiceJunction (
	shelterId integer REFERENCES shelter (id) NOT NULL,
	serviceId integer REFERENCES service (id) NOT NULL,
	PRIMARY KEY (shelterId, serviceId)
);

CREATE TABLE "user" (
	id SERIAL PRIMARY KEY,
	clientId integer REFERENCES client (id),
	handle varchar(255) UNIQUE NOT NULL,
	pass varchar(255) NOT NULL,
	phone varchar(16),
	email varchar(255)
);

CREATE TABLE department (
	id SERIAL PRIMARY KEY,
	shelterId integer REFERENCES shelter (id) NOT NULL,
	name varchar(255) NOT NULL
);

CREATE TABLE shift (
	id SERIAL PRIMARY KEY,
	departmentId integer REFERENCES department (id),
	start timestamp NOT NULL,
	"end" timestamp NOT NULL
);

CREATE TABLE donor (
	stripeCustomerId varchar(255)
);

CREATE TABLE individual (
	id SERIAL PRIMARY KEY,
	firstName varchar(255) NOT NULL,
	lastName varchar(255) NOT NULL
) INHERITS (donor);

CREATE TABLE donation (
	id SERIAL PRIMARY KEY,
	shelterId integer REFERENCES shelter (id) NOT NULL,
	individualId integer REFERENCES individual (id),
	orgId integer REFERENCES organization (id),
	amount numeric NOT NULL,
	whenDonated timestamp NOT NULL,
	recurrence interval
);

CREATE TABLE organization (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL
) INHERITS (donor);

CREATE TABLE volunteer (
	id SERIAL PRIMARY KEY,
	position varchar(255),
	firstName varchar(255) NOT NULL,
	lastName varchar(255) NOT NULL
);

CREATE TABLE volunteerShiftJunction (
	volunteerId integer REFERENCES volunteer (id) NOT NULL,
	shiftId integer REFERENCES shift (id) NOT NULL,
	PRIMARY KEY (volunteerId, shiftId)
);