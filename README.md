# this project under structure (a lot of random code, not optimize, use it on your own risk)

# our target to host golang modules on ipfs

- we need to host git-repo in ipfs : 
    - the command could be : ipfs.house publish . 
    - called the project publisher 
    - host-git-in-ipfs : https://docs-beta.ipfs.io/how-to/host-git-style-repo/




## to check 
- check host-git-in-ipfs : https://docs-beta.ipfs.io/how-to/host-git-style-repo/
- hostname-and-a-port to domain localhost
- check ipfs tutorials https://proto.school/#/
- how gitlab support go module : https://gitlab.com/gitlab-org/gitlab-foss/issues/65681 


check how to use gitlab from here
https://www.youtube.com/watch?v=EkvxotEYTok
https://docs.gitlab.com/omnibus/docker/



docker run --rm -it -p 443:443 -p 80:80 -p 22:22 \
--volume $GITLAB_HOME/gitlab/config:/etc/gitlab \
--volume $GITLAB_HOME/gitlab/logs:/var/log/gitlab \
--volume $GITLAB_HOME/gitlab/data:/var/opt/gitlab gitlab/gitlab-ce:latest


go get "gateway.ipfs.io/ipfs/QmfHGarfN3gLJMjihjNxnh8emdLZNM5RKuAYdpnLFc9zva/gopkg.git"
go get -insecure local.ipfs.house/root/hello.git
 
- delete a file from ipfs: https://discuss.ipfs.io/t/how-can-i-delete-a-file-from-ipfs/1556
You can remove the pin and then run the garbage collector.
Ideally that would remove it, but just from your node.
```
ipfs pin rm $YOUR_HASH

ipfs repo gc
```
