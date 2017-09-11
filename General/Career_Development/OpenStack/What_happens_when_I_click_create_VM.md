Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2016-12-22T09:05:10+00:00

###### What happens when I click create VM ######
Created Thursday 22 December 2016

1. Log into the **Dashboard** UI
2. If authentication token is not cached send a request to **Keystone **to get it
3. **Keystone **authenticates and authorizses user (sends token back to be used in all subsequent requests)
4. Specify vm parameters and hit launch
5. **Nova API **takes parameters, checks token
