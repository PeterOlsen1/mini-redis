* Current
  * Use RESP parsing in server start method
    * can't do that becuase of the way we read connections
    * before server updates: 
      BenchmarkGetAndSet-16    	    4035	    265212 ns/op	    1056 B/op	      27 allocs/op
    * after
      BenchmarkGetAndSet-16    	    4260	    268109 ns/op	    1056 B/op	      27 allocs/op
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
