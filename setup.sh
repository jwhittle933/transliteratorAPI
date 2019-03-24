if [ -d ~/go ]; 
    then
        echo "Found ${HOME}/go\n"
        echo "Setting go variables."
        setVars 
    else 
        touch ${HOME}/go
        echo "Created ${HOME}/go\n"
fi


function setVars(){
    export GOPATH=$HOME/go
    export GOROOT=/usr/local/opt/go/libexec
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:$GOPATH/bin
    export PATH=$PATH:$GOROOT/bin
}