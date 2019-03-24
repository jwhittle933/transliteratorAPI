if [ ! -d ~/go ]; 
    then
        mkdir ~/go
        mkdir ~/go/src
        mkdir ~/go/bin
        mkdir ~/go/pkg
        export GOPATH=$HOME/go
        export GOROOT=/usr/local/opt/go/libexec
        export GOBIN=$GOPATH/bin
        export PATH=$PATH:$GOPATH/bin
        export PATH=$PATH:$GOROOT/bin
    else 
     echo "found."  
fi