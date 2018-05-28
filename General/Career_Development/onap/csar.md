# csar

cloud service archive format from ONAP

container file using the zip file format
A container file is an archive file which contains multiple files of possibly different file types

includes:

* service template of a cloud app
* all artifacts required to manage the lifecycle of the corresponding cloud app
* all artifacts to execute the cloud app

must contain:

/Meta-Inf
* MANIFEST.MF
    metadata of the other files in the csar
    in the form of k/v pairs