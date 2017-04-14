# rnspool
car pool back-end web service in Golang with Postgresql database

This is a project I'm taking on because one I think its a cool Idea and would love to see it come alive in the future and two
I really enjoy programing a golang and want put a challange forward to help me better understand the language and all its power.
I'm taking it one step at a time, for the time being I'm staying away form goroutines, but the idea is to bring them in once 
I have a solid version of the back-end webservice working.

general idea of back-end webservice:

This back-end webservice is intended for a car pooling application. Essentially its a REST API that takes in Json request from
a client a provides Json responses to client.

There is two types of users: Driver and Rider
  Driver offers rides, Riders needs ride. The nuts and bolts of the API is that given a route offered by driver, and route needed
  by rider it links these two. Providing rich json responses to enable a front-end client to link up a driver and rider.
  
  
Ideas for the future:
  Add goroutines
  
  develope front-end apps using React Native (android and ios)
