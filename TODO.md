## Todo

### issue

the problem is we need some kind of mechanisim to retry our commit if we lost our database connection. these are some
notes:
we need add a decorator to retry handler inside that and its important that session should be new session because the
old one is expired when we lost connection
(go episode 3 from module 6 of cqrs learning)

1. How can I add a session factory to our database decorator to handle retry session?
2. How can ad a general Database Retry Decorator for our database?
3. Add transaction support
4. 