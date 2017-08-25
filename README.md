# cctldcentral

Server to retrieve and store all statistics from ccTLDs. It performs a search in
public statistics once a day, and optionally communicates with the server
[cctldstats](http://github.com/rafaeljusto/cctldstats) for retrieving non-public
data.

## Install

You will need a [PostgreSQL](https://www.postgresql.org/) database to store all
statistics retrieved from the ccTLDs.

```
sudo apt-get install postgresql
sudo -u postgres psql < cctldcentral.sql
sudo -u postgres psql cctldcentral < cctldcentral.dump.sql
```