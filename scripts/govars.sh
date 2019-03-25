if [ -d ~/go ]; 
    then
        echo ">>>  Found ${HOME}/go"
        if [ -n "$GOPATH" ]; 
            then 
                echo ">>>  Found GOPATH ${GOPATH}."
                echo ">>>  Please ensure that your go variables are set correctly."
            else 
                echo "Setting go variables."
                setVars 
        fi
    else
        echo "Didn't find ${HOME}/go. Creating directory." 
        mkdir ~/go
        mkdir ~/go/src
        mkdir ~/go/bin
        mkdir ~/go/pkg
        echo "Created ${HOME}/go\n"
        echo "Setting go variables."
        setVars
fi


function setVars(){
    export GOPATH=$HOME/go
    export GOROOT=/usr/local/opt/go/libexec
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:$GOPATH/bin
    export PATH=$PATH:$GOROOT/bin
}