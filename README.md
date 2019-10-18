# utilities

Defines and implements some repetitive parts of code, which can be reused inside micro services, which define xdevices system. To not repeat code (keeping DRY principle), we have to have place, where the repetitive code should be available for reuse. Contains templates for:

- registration tickets for go-based micro services, which makes them able to register inside eureka's discovery service
- connectivity between go-based micro services and rabbit messaging broker
- database connectivity (sqlite) for go-based micro services

---

Part of xdevices application - simple micro services-based system for monitoring. <br/>
**Developed for udemy course "Beginners guide to microservices with Go, Spring and RaspPi".**
