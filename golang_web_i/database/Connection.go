package database


Golang MySQL Tutorial
Elliot Forbes
Elliot Forbes
⏰ 5 Minutes
📅 Apr 9, 2017
As you continue your Golang learning journey, it becomes almost inevitable that you will have to interact with some form of database.

In this tutorial I’ll be demonstrating how you can connect to a MySQL database and perform basic SQL statements using Go.

Why MySQL?
MySQL is one of the most well-known and well-used database technologies available to developers at the present point in time. It has an absolutely massive community around it and it’s quite possibly powering half the web as the main database technology for Wordpress.

It’s incredibly easy to spin up a MySQL instance locally and thus it’s perfect for building some decent applications on top of.

Note - Technology choice shouldn’t be based on popularity, there may be scenarios where you need to consider alternatives such as CockroachDB or NoSQL databases.

Video Tutorial
If you prefer following a video, then this tutorial is available in video format here:


Text Tutorial
In order to do this we’ll be using https://github.com/go-sql-driver/mysql as our MySQL driver. Go-SQL-Driver is a lightweight and fast MySQL driver that supports connections over TCP/IPv4, TCP/IPv6, Unix domain sockets or custom protocols and features automatic handling of broken connections.

Github Repo: go-sql-driver/mysql

Connection Pooling
If you are building high-performance database applications, connection-pooling is an absolute must.

Thankfully, the opensource package that we’ll be using for the basis of this tutorial features automatic connection-pooling thanks to it’s use of of the database/sql standard package.

This essentially means that, every time you query your database, you are using a connection from a pool of connections that have been set up on application startup. These connections are reused, time and time again, and this subsequently means you aren’t creating and destroying a new connection every time you perform a query.

Implementation
We’ll begin by connecting to a database we’ve set up on our local machine and then go on to perform some basic insert and select statements.

Connecting to a MySQL database
Let’s create a new main.go file. Within this, we’ll import a few packages and set up a simple connection to an already running local database. For the purpose of this tutorial, I’ve started MySQL using phpmyadmin and I’ve created a database called test to connect to and create tables within.

We’ll use sql.Open to connect to our database and set up our automatic connection pool, this will return either db or an err that we can handle.

package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func Connect() {

    
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

}