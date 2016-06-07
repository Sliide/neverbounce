# neverbounce

This is a client for the Email validation service Neverbounce(https://neverbounce.com/)

The implementation of this library was based on the docs: https://neverbounce.com/help/api/verifying-an-email/ 

Here is an example.

Usage:

```
   
neverbounce.Init(&neverbounce.NeverBounceCli{        
    ApiUsername: "NEVERBOUNCE_USERNAME",  
    ApiPassword: "NEVERBOUNCE_PASSWORD",  
})                                                   
neverbounce.VerifyEmail("someemail@example.com")

```  



Due to the fact that this service might be used on the web in order to validate forms. This minimalist client allows
a user to specify its test mode

``` 
    
neverbounce.Init(&neverbounce.NeverBounceCli{        
    ApiUsername: "NEVERBOUNCE_USERNAME",  
    ApiPassword: "NEVERBOUNCE_PASSWORD",
    TestMode: true,
})

neverbounce.VerifyEmail("anemail@valid.com")
neverbounce.VerifyEmail("anemail@invalid.com")
neverbounce.VerifyEmail("anemail@catchall.com")
neverbounce.VerifyEmail("anemail@disposable.com")
neverbounce.VerifyEmail("anemail@unknown.com")

```


When doing so this module will fake the never bounce service and the client will than make requests to the fake server.

The emails will than be validated follwing the creteria:

- Emails ending in @valid.com  will be considered valid
- Enails ending in @invalid.com will be considered invalid
- Emails ending in @disposable.com will result on a disposable result code. 
- Emails ending in @catchall.com will result on a catch all result code.
- Email ending in @unknow.com will result on an unknown status code.


Please refer the services API documentation for further information.

https://neverbounce.com/help/api/verifying-an-email/

