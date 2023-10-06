# oauth2

https://tools.ietf.org/pdf/rfc6749.pdf
https://spring.io/guides/tutorials/spring-boot-oauth2/

The OAuth 2.0 authorization framework enables a third-party application to obtain limited access to an HTTP service.

Instead of using credentials a token is issued which has a limited time of access, among other things.

I use my username and password to get a token, which is then used until it expires at which point I have to log in again

TLS should be used to encrypt any user/pass combinations before the token is received

* Actors
  The different roles in the process
* Clients
  The app, you are using to access data from

## roles

* resource owner
  * the owner of the resources which are protected
* resource server
  * the server hosting the protected resources
* client application
  * an app making protected resources requests.
* authorization server
  * the server issuing access tokens to the client after successfully authenticating the resource owner and obtaining authorization
  * the client registers to the authorization server before being able to obtain a token

I write a client application and configure it to use an authorization server
When I log into the client application it redirects to the auth server with my credentials to request an access token
If that is successful I am redirected back to my client app and I can access a resource server with the token
  The access token is a cookie, JSESSIONID by default.

## access token

are credentials used to access protected resources.
is a string representing an authorization issued to a client.
is usually opaque to the client.
represent specific scopes and durations of access and enforced by the resource server and authorization server.

## refresh token

a token give to a client which can be used to get a new access token when the current one expires

## protocol endpoints

* authorization endpoint
  * used by client to obtain authorization grant
* token endpoint
  * used by client to exchange a grant for access token
* redirection endpoint
  * used by the auth server to return responses containing auth creds to the client via the resource owner user agent

## redirection

When your client redirects to auth server it gives a redirect uri to go to after authentication is successful

## flow

client request to auth server, with username and password.
auth server responds with code
client requests access token from auth server with code.
client now has access token to access protected resources.

## terminology

Scopes
* additive bundles of permissions asked by the client when requesting a token
* Decouples auth policy decisions from enforcement
* 
