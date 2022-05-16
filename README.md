# GO Feature Flag demo

This repository contains a demo app which display a webapp containing a grid of users.  
Each square is a different user.  

With this demo app you can modify your flag and visually see which users are impacted by the change.

https://user-images.githubusercontent.com/17908063/168597717-489d64e6-3e52-4d3c-aa7c-86f89dabd278.mp4


## About the app
The app use `labstack/echo` as an http server and serve an HTML page with one square per user.

Every square has his own UUID to represent a user, it means that you play with you flag and directly see which user will be impacted.

## Build the app

To build the app you have to run these command:

```shell
go mod vendor # to retrieve dependencies
go buid .
```

## Report a problem
If you have any issue with this demo app you can open an issue.

If the problem is related to the SDK please open the issue in the [`thomaspoignant/go-feature-flag` repository](https://github.com/thomaspoignant/go-feature-flag/issues/new/choose).
