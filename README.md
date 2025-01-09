# CC Simple Auth Server
This is the server portion of a simple, self-contained auth system. It maintains a database of users, and provides REST (albeit, not RESTful) endpoints for user registration and login.

This module can be used on its own or combined with CC Simple Auth Client to create an end to end auth system that supports user sessions.

The server module itself does not issue session tokens, leaving that particular implementation up to the client for greater flexibility.

This module is part one of my project to break up my LibrePub (librepub.org) web app into discrete, self-contained services that can be hosted separately.
