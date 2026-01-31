* Current
  * Handlers that receive OK just return err instead?
  * Use RESP parsing in server start method
  * User permissions per database?
  * Individual database info
  * Add / remove databases 
    * maybe loop through all users to see if they are connected to the current and remove if they are
  * Notification events on key change?

* Way backlog
  * Update ListSaves to return a [] instead of string
  * Command to get all commands, maybe a string about what they do
  * Pipelining
  * Blocking pop calls
  * Include length calculations in RESP data
