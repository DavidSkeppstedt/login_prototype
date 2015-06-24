# A Login Prototype with Go and Martini.
A toy app to learn Go web programming.
This handles a REST-api for which should be used for a client app that can login and logout etc.

## Dependencies
So far we have these dependencies that you can easily fetch with:
```
	go get github.com/_name_of_dep
```

* [Martini](https://github.com/go-martini/martini)
* [PostgreSQL adapter](https://github.com/lib/pq)

You would have to have Go set up correctly ofcourse.

### Setting up Go for development
Follow these links for some pointers.
[Installation](https://golang.org/doc/install)

[Config](https://golang.org/doc/code.html)

Notice: 
The installation path and the workspace is NOT the same!

### First time setup
Run the setup shell script inside the setup folder.

On OSX/Linux you would do something like:
```
sh setup/setup.sh
```

This will generate necessary folders and files for the project.

###Config
####Database
Inside the file dbconfig.config put all your config
for your postgres db. It should follow this convention:
* host_ip
* port
* user
* password
* dbname
* sslmode
