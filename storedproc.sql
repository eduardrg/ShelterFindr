# Returns the number of shelters which have
# the the desired occupancy.
CREATE OR REPLACE FUNCTION sheltersWithOccupancy(requiredSpaces integer)
RETURNS integer as $total$
declare
	total integer;
BEGIN
	SELECT count(*) INTO total FROM shelter
	WHERE id IN (SELECT DISTINCT(shelterid) FROM room
	GROUP BY room.id HAVING capacity - occupancy >= requiredSpaces);
	RETURN total;
END;
$total$ LANGUAGE plpgsql;

# Adds a clientId to a shelterId
# clientId and shelterId are Fks, so they must already exist in
# client and shelter (respectively)
CREATE OR REPLACE FUNCTION addToWaitlist(shelterId integer, clientId integer)
RETURNS void as $BODY$
BEGIN
 	INSERT INTO waitlist (shelterid, clientid, whenadded)
	VALUES (shelterId, clientId, NOW());
END;
$BODY$
LANGUAGE plpgsql;
