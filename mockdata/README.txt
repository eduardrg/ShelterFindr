The .csv files in this directory were generated using Mockaroo. To load them
into the database, we connected to the database and manually executed this
statement for each table:

\copy <tablename> FROM '<tablename>.csv' DELIMITER ',' CSV HEADER;

Donation.csv and donation(2).csv are both intended to be loaded into the
donation table. Donation(2).csv is for donations made by individual donors
and donation.csv is for donations made by organizational donors.