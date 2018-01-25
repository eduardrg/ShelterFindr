# Welcome to ShelterFindr
Here, you can help find a shelter that best fits your needs.

Created for INFO 340: Introduction To Relational Database Management Systems.

http://heroku-postgres-97f2b1ad.herokuapp.com/

# Main Page
!Screenshot: s[Main landing page](https://github.com/eduardrg/ShelterFindr/blob/master/screen1.PNG)

# Admin View
![Screenshot: sEditing an entry as an admin](https://github.com/eduardrg/ShelterFindr/blob/master/screen2.PNG)

# User View
![Screenshot: searching for shelters as a user](https://github.com/eduardrg/ShelterFindr/blob/master/screen3.PNG)


# Configuration and Deployment Steps
## Setting Up the Database
0. Clone or download this repository and add it to your `GOPATH` if necessary.
1. Install Heroku command line tools.
2. Create a new Postgres database from Heroku's Dashboard > Databases.
3. Copy the Psql command from "Connection Settings." It should look like `heroku pg:psql --app [DATABASE]`. This command will open a psql shell to your database.
4. Set up tables and indices. The ddl.sql file contains SQL statements to create 22 tables with their columns and 8 indices. Pipe the file to the psql command. It should resemble: `cat ddl.sql | heroku pg:psql --app [DATABASE]`
5. Press enter and the command will execute. If successful, it should output 22 lines of `CREATE TABLE` followed by 8 lines of `CREATE INDEX`. Make note of any errors and adjust your ddl.sql accordingly. If you need to start over, `DROP SCHEMA public CASCADE;
CREATE SCHEMA public;` will give you a fresh slate.
6. Find your `DATABASE_URL` in your Connection Settings for your database on Heroku's dashboard. Copy this URL. You will need it to connect your GO app to your database.

## Deploying the Application Locally
7. Create a file in your cloned repo named .env and containing the line `DATABASE_URL=[YOUR-URL-HERE]`.
8. From a terminal, navigate to your repo's cmd/lab7, and run `go install`
9. From your repo's root directory, run `heroku local`
10. Open an HTML page of your app in a browser and it should work.

## Deploying the Application on Heroku
1. From Heroku's dashboard, create a new app. Set your app name (hereafter `app-name`) and runtime.
2. While the app is being created, you can run `heroku login` in a terminal and authenticate.
3. You copied your `DATABASE_URL` earlier--hopefully it's still on your clipboard. Go to Heroku's dashboard, navigate to `app-name`'s settings, and in the config vars, create a variable `DATABASE_URL` and set it to the URL you copied.
4. Navigate to your repo's root directory and run `heroku git:remote -a app-name`
5. You can now push your code up to the Heroku app with `git push heroku master`.
6. After you've pushed your code, you can open your app with `heroku open` or by visiting `http://app-name.herokuapp.com/`

## Contributors:
Eduard Grigoryan, Megh Vakharia, and Brittney Hoy

## Sources:
https://github.com/uw-info340b-sp2016/lab6-7
