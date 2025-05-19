# Blog Aggregator

- Simple RSS feed aggregator in GO
- CLI tools that allows users to:
  - Add RSS feeds from across the internet to be collected
  - Store the collected posts in a PostgreSQL database
  - Follow and unfollow RSS feeds that other users have added
  - View summaries of the aggregated posts in the terminal,
    with a link to the full post

## Learning Goals

- Learn how to integrate a `Go` application with a `PostgreSQL` database
- Practice using SQL skills to query and migrate database
- Learn how to write a long-running service that continuously fetches new posts
  from RSS feeds and stores them in the database

## Package INFO

- internal/config
  - Responsible for reading and writing the JSON file.
