# Gator

Website aggregator that uses postgres and go to keep up to date on the posts that matter to you!

Installation notes: 
1) You'll need go and Postgres to run the program
2) You can install the files by doing typing ```go install github.com/coolarif/Gator@latest```  
3) Look at how to setup the database at Setting up the database
4) Read the Usage guide on how to enter commands into your command line.

# Setting up the database
1) Install postgres.  

For Linux/WSL:  
```
sudo apt update 
sudo apt install postgresql postgresql-contrib
```  
For MacOS: 
```
brew install postgresql@16
```

2) Check it is installed with ```psql --version```

3) Set up password (for linux only) ```sudo passwd postgres```

4) Set up the postgres server in the background<br>
for Mac: ```brew services start postgresql```
for Linux: ```sudo service postgresql start```

5) Enter the psql shell:<br>
for Mac: ```psql postgres```
for Linux: ```sudo -u postgres psql```

You should now see a prompt like this: 
```
postgres=#
```

6) Create a new database, I called mine Gator:
```
\c gator
```
You should see a new prompt that looks like this
```
gator=#
```

7) Set the user password (linux only)
```
ALTER USER postgres PASSWORD 'postgres';
```  
In this case the password is postgres  
  
# Usage
To be Updated
<<<<<<< HEAD
update check
=======
update
>>>>>>> 914181e (pulled the thing locally hopefully everything is fine)
