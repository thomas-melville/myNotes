# csar

TOSCA YAML Cloud Service Archive format from ETSI

http://www.etsi.org/deliver/etsi_gs/NFV-SOL/001_099/004/02.03.01_60/gs_nfv-sol004v020301p.pdf

container file using the zip file format
A container file is an archive file which contains multiple files of possibly different file types

Structure complies with TOSCA Simple Profile YAML v1.1

includes:

* service template of a cloud app
* all artifacts required to manage the lifecycle of the corresponding cloud app
* all artifacts to execute the cloud app

must contain:

/Meta-Inf
* MANIFEST.MF
    metadata of the other files in the csar
    in the form of k/v pairs

ETSI Video: https://www.brighttalk.com/webcast/12761/265769

VNF Package contains:

* VNFD (descriptor)
* * defines VNF properties and requirements, such as
* * * Resources needed
* * * Connectivity
* * * LCM Behaviour
* * * References to sw images and other artifacts
* * * Affinity and anti-affinity policy rules
* * * Deployment flavours
* software images
* Manifest file
* additional files

VNF Package is a CSAR.
The structure and format of a VNF package shall conform to the TOSCA Simple Profile YAML v1.1 spec for CSAR format
VNFD is the main TOSCA definitions YAML file inside the archive

There are 2 Directory structure options

Option 1 (Which is the option we will be using)

/TOSCA-metadata
| - TOSCA.meta (This has entries pointing to all the files below)
/Definitions
| - MRF.yaml
| - otherTemplates
    | - ...
/Files
| - ChangeLog.txt
| - MRF.cert
| - /images
    | - ...
| - other artifacts
| - /Tests
    | - ...
| - /Licenses
    | - ...
/Scripts
| - ...
/MRF.mf

Option 2

/MRF.yaml
/MRF.mf
/MRF.cert
/ChangeLog.txt
/Tests
| - ...
/Licenses
| - ...
/Artifacts
| - scripts
| - images
| - templates
| - ...

**A lot about ensuring integrity and authenticity of packages.**

It can be handled in one of two ways.

* Individual files can have digests provided which you can compare against.
* the complete csar file is digitally signed by the VNF provider