#AWS RDS MySQL
this is just for learning purpose.

1. create RDS instance
   please see the [screenshots](screenshots)
    ```
    Endpoint: database-1.ch3fcoesrgva.us-east-2.rds.amazonaws.com
    user: admin
    pwd: yzTvwnqe6AJ6CYb   
   ```

2. connect to RDS instance
    ```shell
    mysql -h database-1.ch3fcoesrgva.us-east-2.rds.amazonaws.com -u admin -p
    ```
    then enter the password
    ```shell
    Enter password:
    Welcome to the MySQL monitor.  Commands end with ; or \g.
    Your MySQL connection id is 14
    Server version: 8.0.20 Source distribution
    
    Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.
    
    Oracle is a registered trademark of Oracle Corporation and/or its
    affiliates. Other names may be trademarks of their respective
    owners.
    
    Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
    
    mysql>
    mysql> show databases;
    +--------------------+
    | Database           |
    +--------------------+
    | information_schema |
    | mysql              |
    | performance_schema |
    +--------------------+
    3 rows in set (0.29 sec)
    
    mysql>
    ```


3. create schema
   ```mysql
    create schema test2;
    ```

4. create table
    ```mysql
    create table amigos
        (
        aID   int auto_increment
        primary key,
        aName varchar(250) null
        );
    ```

5. insert data to table
    ```mysql
    INSERT INTO test2.amigos (aName) VALUES ('Jorge');
    INSERT INTO test2.amigos (aName) VALUES ('Felipe');
    INSERT INTO test2.amigos (aName) VALUES ('Padre');
    INSERT INTO test2.amigos (aName) VALUES ('Alberto');
    ```

6. query data from table
    ```mysql
    select * from amigos;
    ```
   ```mysql
    +-----+---------+
    | aID | aName   |
    +-----+---------+
    |   1 | Jorge   |
    |   2 | Felipe  |
    |   3 | Padre   |
    |   4 | Alberto |
    +-----+---------+
    4 rows in set (0.25 sec)
    ```


