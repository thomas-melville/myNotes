# awk

useful command for extracting specific text from formatted strings

echo "..." | awk '{print $3}'

print the third piece of text per line after each line has been split

## flags

-F <text>   separator to use to split test, default is space
            multiple separators can be specified using |
            this can also be a regular expression
-I <text>   the identifier to use for the input coming from each line

## examples


kubectl get nodes -o wide | awk '{print $6}' | xargs -I {} scp retagged-images.tar {}:/home/eccd
  get all the information about the worker nodes
  awk out the ips
  scp a file on each ip

awk '{print $6}' <(cat nodes.txt) | xargs -I {} ssh {} sudo docker load -i retagged-images.tar;
  had to output the node information to file so I could delete the entries which didn't have valid ips
  
