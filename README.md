# jwt-go-rbac
* This repo contains a simple room booking API to demonstrate secure implementation of Role-based Access Control (RBAC) in Go with the jwt-go package.
* Basically two role is defined
1. Admin
2. User
* Makes use of JWT for AUTHENTICATION and AUTHIRIZATIO.
* All the functionalities are grouped under four different route:
1. authRoutes :- under this we have "login and register".
2. protectedRoutes :- requires user JWT authorization and functionalities are:
    * GetUserBookings
	* CreateBooking
3. adminRoutes :- requires admin JWT authorization and functionalities are:
    * GetUsers
	* GetUser(by id)
	* UpdateUser
	* CreateRole
	* GetRoles
	* UpdateRole
	* CreateRoom
	* UpdateRoom
	* GetBookings
4. publicRoutes :- under this we have "get room and get room by id and these two don't need Jwt authorization as these are read only."
