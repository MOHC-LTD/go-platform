## go-platform

1) ### Configure Go/git on your machine to pull the library from a private repository

   (you only need to do this once)

     ```
     # Tell golang that this is a private repository
     go env -w GOPRIVATE=github.com/MOHC-LTD/*
   
     # Setup git to use SSH for github (so it can use the ssh auth key to pull down the private repo)
     git config --global --add url."git@github.com:".insteadOf "https://github.com/"```
   
     # OR download a local copy of the repository and setup git to use this instead
     ```replace github.com/MOHC-LTD/go-platform v0.0.x => <your_physical_path>\go-platform```

2) ### Integrate this library with a Go app
   In your go.mod, assuming you require the library as follows:

   ```
   require github.com/MOHC-LTD/go-platform v0.0.x
   # or the part of the lib that you require
   require github.com/MOHC-LTD/go-platform/logger v0.0.x
   ```
3) ### In the directory of the target module, install the library

   ```
   go get github.com/MOHC-LTD/go-platfrom
   #or just get all the project dependancies
   go get ./... 
   # to test with a specific dev branch use on the terminal in the project root:
   go get github.com/MOHC-LTD/go-platform@<branch name>
   # update any changed dependancies 
   go get ./...
   ```