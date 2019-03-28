send_file()
{
    echo curl -F `file=@${1}` localhost:3000/upload
}

export -f send_file
